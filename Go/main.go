package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Item struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Description string `json:"description"`
	Img         string `json:"img"`
	AvailableS  int    `json:"available_S"`
	AvailableM  int    `json:"available_M"`
	AvailableL  int    `json:"available_L"`
}

type Database struct {
	Items []Item `json:"items"`
}

func main() {
	// Charger le contenu du fichier JSON
	data, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Fatal(err)
	}

	// Analyser le contenu JSON dans la structure de données
	var db Database
	err = json.Unmarshal(data, &db)
	if err != nil {
		log.Fatal(err)
	}

	// Définir les routes pour les requêtes HTTP
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Récupérer les paramètres de filtre
			name := r.URL.Query().Get("name")
			categoryStr := r.URL.Query().Get("category")
			category, _ := strconv.Atoi(categoryStr)

			// Filtrer les items en fonction des paramètres
			filteredItems := filterItems(db.Items, name, category)

			// Renvoyer la liste filtrée en tant que réponse JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(filteredItems)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/reserve", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Récupérer les données de réservation depuis le corps de la requête
			var reservation struct {
				ItemName string `json:"item_name"`
				Size     string `json:"size"`
			}
			err := json.NewDecoder(r.Body).Decode(&reservation)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Réserver l'item dans la taille spécifiée
			err = reserveItem(&db, reservation.ItemName, reservation.Size)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Renvoyer une réponse de succès
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Reservation successful")

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Démarrer le serveur HTTP
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Fonction de filtrage des items par nom et catégorie
func filterItems(items []Item, search string, category int) []Item {
	filtered := make([]Item, 0)

	for _, item := range items {
		if item.Type == category || category == 0 {
			if item.AvailableS > 0 || item.AvailableM > 0 || item.AvailableL > 0 {
				if search == "" || containsIgnoreCase(item.Name, search) {
					filtered = append(filtered, item)
				}
			}
		}
	}

	return filtered
}

// Fonction utilitaire pour vérifier si une chaîne de caractères contient une sous-chaîne (ignorant la casse)
func containsIgnoreCase(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}

// Fonction de réservation d'un item dans une taille spécifique
func reserveItem(db *Database, name, size string) error {
	for i, item := range db.Items {
		if item.Name == name {
			switch size {
			case "S":
				if item.AvailableS > 0 {
					db.Items[i].AvailableS--
					return nil
				}
			case "M":
				if item.AvailableM > 0 {
					db.Items[i].AvailableM--
					return nil
				}
			case "L":
				if item.AvailableL > 0 {
					db.Items[i].AvailableL--
					return nil
				}
			}
			return fmt.Errorf("Item is not available in size %s", size)
		}
	}

	return fmt.Errorf("Item not found")
}
