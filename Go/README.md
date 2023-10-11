# Projet Backend Go

Ce projet est un backend simple en Go qui utilise une base de données JSON pour fournir une liste filtrée d'articles et permettre la réservation d'un article dans une taille spécifique.

## Prérequis

- Docker : Assurez-vous d'avoir Docker installé sur votre système.

## Accédez au répertoire du projet :

cd Go

## Construisez l'image Docker :

docker build -t go_backend .

# Utilisation

Lancez le conteneur Docker :

docker run -p 8080:8080 go_backend
