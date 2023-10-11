from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
import json

database = None


def load_database():
    global database
    with open('database.json', 'r') as f:
        database = json.load(f)


def filter_items(request):
    if request.method == 'GET':
        name = request.GET.get('name', '')
        category = request.GET.get('category', '')
        filtered_items = filter_items_by_name_and_category(name, category)
        return JsonResponse(filtered_items, safe=False)
    else:
        return JsonResponse({"error": "Method not allowed"}, status=405)


def reserve_item(request):
    if request.method == 'POST':
        data = json.loads(request.body)
        item_name = data.get('item_name', '')
        size = data.get('size', '')
        try:
            result = reserve_item_by_name_and_size(item_name, size)
            if result:
                return JsonResponse({"message": "Reservation successful"})
            else:
                return JsonResponse({"error": "Item not available in the specified size"}, status=400)
        except Exception as e:
            return JsonResponse({"error": str(e)}, status=400)
    else:
        return JsonResponse({"error": "Method not allowed"}, status=405)


def filter_items_by_name_and_category(name, category):
    filtered_items = []
    for item in database['items']:
        if category == '' or item['type'] == int(category):
            if item['available_s'] > 0 or item['available_m'] > 0 or item['available_l'] > 0:
                if name == '' or name.lower() in item['name'].lower():
                    filtered_items.append(item)
    return filtered_items


def reserve_item_by_name_and_size(item_name, size):
    for item in database['items']:
        if item['name'] == item_name:
            if size == 'S':
                if item['available_s'] > 0:
                    item['available_s'] -= 1
                    return True
            elif size == 'M':
                if item['available_m'] > 0:
                    item['available_m'] -= 1
                    return True
            elif size == 'L':
                if item['available_l'] > 0:
                    item['available_l'] -= 1
                    return True
            raise Exception("Item is not available in size {}".format(size))
    raise Exception("Item not found")


load_database()
