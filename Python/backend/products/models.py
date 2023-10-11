from django.db import models

class Item(models.Model):
    name = models.CharField(max_length=100)
    type = models.IntegerField()
    description = models.TextField()
    img = models.CharField(max_length=256)
    available_s = models.IntegerField()
    available_m = models.IntegerField()
    available_l = models.IntegerField()
