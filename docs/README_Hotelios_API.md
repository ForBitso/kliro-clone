# 🏨 Hotelios API - Инструкция для фронтендера

## 📁 Файлы документации

1. **`Hotelios_API_Complete.postman_collection.json`** - Postman коллекция со всеми 26 API
2. **`hotelios_api_documentation.md`** - Подробная документация в Markdown
3. **`README_Hotelios_API.md`** - Эта инструкция

## 🚀 Быстрый старт

### 1. Импорт в Postman

1. Откройте Postman
2. Нажмите **Import** (импорт)
3. Выберите файл `Hotelios_API_Complete.postman_collection.json`
4. Коллекция будет импортирована со всеми 26 API

### 2. Настройка переменных

В Postman коллекции уже настроена переменная:
- `base_url` = `http://localhost:8080`

Если нужно изменить URL сервера:
1. Откройте коллекцию
2. Перейдите на вкладку **Variables**
3. Измените значение `base_url`

### 3. Тестирование API

#### Справочные API (работают без дополнительных параметров):
- `POST /hotels/countries` - Список стран
- `POST /hotels/types` - Типы отелей  
- `POST /hotels/facilities` - Список удобств
- `POST /hotels/currencies` - Список валют

#### Booking-Flow API (требуют данные в body):
- `POST /hotels/search` - Поиск отелей
- `POST /hotels/booking/read` - Детали бронирования

## 📋 Структура коллекции

### 📚 Справочные API (v1.0) - 20 endpoints
```
📋 Справочные API (v1.0)/
├── GetCountryList - Список стран
├── GetRegionList - Список регионов  
├── GetCityList - Список городов
├── GetHotelTypeList - Типы отелей
├── GetHotelList - Список отелей
├── GetHotelPhotoList - Фотографии отеля
├── GetHotelRoomTypeList - Типы номеров
├── GetHotelRoomTypesPhotoList - Фото номеров
├── GetFacilityList - Список удобств
├── GetHotelFacilityList - Удобства отеля
├── GetEquipmentList - Список оборудования
├── GetRoomTypeEquipmentList - Оборудование номеров
├── GetPriceRange - Диапазон цен
├── GetStarList - Список звезд
├── GetNearbyPlacesTypeList - Типы ближайших мест
├── GetHotelNearbyPlacesList - Ближайшие места отеля
├── GetServicesInRoomList - Услуги в номере
├── GetHotelServicesInRoomList - Услуги в номере отеля
├── GetBedTypeList - Типы кроватей
└── GetCurrencyList - Список валют
```

### 🏨 Booking-Flow API (v1.1.0) - 6 endpoints
```
🏨 Booking-Flow API (v1.1.0)/
├── BookingFlowSearch - Поиск отелей
├── BookingFlowQuote - Актуальные цены
├── BookingFlowCreate - Создание бронирования
├── BookingFlowConfirm - Подтверждение бронирования
├── BookingFlowCancel - Отмена бронирования
└── BookingFlowRead - Детали бронирования
```

## ✅ Автоматическая авторизация

**ВСЕ API работают БЕЗ ручной передачи credentials!**

Сервер автоматически добавляет:
- **login**: `api-0002-001`
- **password**: `d5f12e53a182c062b6bf30c1445153faff12269a`
- **access_key**: `377edeb5-f452-4b10-a24d-67b977892ea9`

**Фронтендеру НЕ нужно передавать эти данные!**

## 🧪 Примеры тестирования

### Тест 1: Список стран
```bash
curl -X POST "http://localhost:8080/hotels/countries"
```

### Тест 2: Поиск отелей
```bash
curl -X POST "http://localhost:8080/hotels/search" \
  -H "Content-Type: application/json" \
  -d '{
    "login": "api-0002-001",
    "password": "d5f12e53a182c062b6bf30c1445153faff12269a",
    "access_key": "377edeb5-f452-4b10-a24d-67b977892ea9",
    "data": {
      "city_id": 67,
      "check_in": "2025/11/25 14:00",
      "check_out": "2025/11/27 12:00",
      "occupancies": [{"adults": 2, "children_ages": []}],
      "currency": "uzs"
    }
  }'
```

## 📖 Дополнительная документация

Для подробной информации о каждом API смотрите файл `hotelios_api_documentation.md`.

## ⚠️ Важные замечания

1. **Сервер должен быть запущен** на `http://localhost:8080`
2. **Все запросы POST** - даже для получения данных
3. **Credentials автоматически добавляются** - не нужно их указывать вручную
4. **Booking-Flow API** требуют передачи данных в теле запроса
5. **Справочные API** работают с пустым телом `{}`

## 🆘 Поддержка

При возникновении проблем:
1. Проверьте, что сервер запущен
2. Убедитесь в правильности URL
3. Проверьте формат данных в запросе
4. Обратитесь к подробной документации в `hotelios_api_documentation.md`
