# Utilisez l'image de base Go
FROM golang:1.20

# Définissez le répertoire de travail
WORKDIR /app

# Copiez le fichier JSON contenant la base de données dans le conteneur
COPY database.json /app/database.json


# Copiez les fichiers source de votre application dans le conteneur
COPY . .

# Construisez l'application Go
RUN go build -o main .

# Exposez le port 8080 pour l'application Go
EXPOSE 8080

# Démarrez l'application une fois le conteneur exécuté
CMD ["./main"]
