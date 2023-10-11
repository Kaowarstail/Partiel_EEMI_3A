# Projet Backend Symfony

Ce projet est un backend développé avec Symfony qui offre une API permettant de filtrer et de réserver des produits à partir d'une base de données JSON.

## Prérequis

Avant de commencer, assurez-vous d'avoir les éléments suivants installés sur votre machine :

- PHP 8.2 ou supérieur
- Composer
- Docker

## Accédez au répertoire du projet :

cd Symfony/symfony_backend

## Construisez l'image Docker :

docker build -t symfony_backend .

# Utilisation

Lancez le conteneur Docker :

docker run -p 8000:80 symfony_backend

