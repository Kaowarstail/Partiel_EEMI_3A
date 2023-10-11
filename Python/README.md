# Projet Backend Python Django

Ce projet est un backend simple en Python Django qui utilise une base de données JSON pour fournir une liste filtrée d'articles et permettre la réservation d'un article dans une taille spécifique.

## Prérequis

- Docker : Assurez-vous d'avoir Docker installé sur votre système.

## Accédez au répertoire du projet :

cd Python

## Construisez l'image Docker :

docker build -t python_backend .

# Utilisation

Lancez le conteneur Docker :

docker run -p 8000:8000 python_backend
