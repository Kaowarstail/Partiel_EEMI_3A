from django.urls import path
from . import views

urlpatterns = [
    path('items', views.filter_items, name='filter_items'),
    path('reserve', views.reserve_item, name='reserve_item'),
]
