# Utilisez l'image de base Python
FROM python:3.9

# Copiez les fichiers source de votre application dans le conteneur
COPY . /app

# Définissez le répertoire de travail
WORKDIR /app

# Installez Django
RUN pip install django
RUN pip install django-cors-headers

# Exposez le port 8000 pour le serveur Django
EXPOSE 8000

# Démarrez le serveur Django une fois le conteneur exécuté
CMD ["python", "backend/manage.py", "runserver", "0.0.0.0:8000"]
