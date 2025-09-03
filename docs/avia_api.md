# API Авиабилетов (Bukhara Integration)

API для работы с авиабилетами через интеграцию с Bukhara API.

## Базовый URL
```
http://localhost:8080/avia
```

## Аутентификация
Все запросы к Bukhara API выполняются автоматически с использованием предустановленных учетных данных.

## Endpoints

### 🔍 Поиск и справочники

#### Поиск авиабилетов
```http
POST /avia/search
```

**Тело запроса:**
```json
{
  "directions": [
    {
      "departure_airport": "TAS",
      "arrival_airport": "IST",
      "date": "2024-01-15"
    }
  ],
  "service_class": "E",
  "adults": 2,
  "children": 0,
  "infants": 0,
  "infants_with_seat": 0
}
```

**Ответ:**
```json
{
  "success": true,
  "result": {
    "offers": [
      {
        "id": "offer_id_123",
        "type": "AV",
        "price": {
          "amount": 1250000,
          "currency": "UZS"
        },
        "directions": [...]
      }
    ]
  },
  "message": "Поиск выполнен успешно"
}
```

#### Список аэропортов
```http
GET /avia/airports
```

#### Классы обслуживания
```http
GET /avia/service-classes
```

#### Типы пассажиров
```http
GET /avia/passenger-types
```

### ✈️ Офферы

#### Получение информации об оффере
```http
GET /avia/offers/{offer_id}
```

#### Правила тарифа
```http
GET /avia/offers/{offer_id}/rules
```

### 🎫 Бронирование

#### Создание бронирования
```http
POST /avia/offers/{offer_id}/booking
```

**Тело запроса:**
```json
{
  "payer_name": "Иван Иванов",
  "payer_email": "ivan@example.com",
  "payer_tel": "+998901234567",
  "passengers": [
    {
      "first_name": "Иван",
      "last_name": "Иванов",
      "middle_name": "Иванович",
      "age": "adt",
      "birthdate": "1990-01-01",
      "gender": "M",
      "citizenship": "UZ",
      "tel": "+998901234567",
      "doc_type": "A",
      "doc_number": "AA1234567",
      "doc_expire": "2030-01-01"
    }
  ]
}
```

#### Информация о бронировании
```http
GET /avia/booking/{booking_id}
```

#### Оплата бронирования
```http
POST /avia/booking/{booking_id}/payment
```

#### Отмена бронирования
```http
POST /avia/booking/{booking_id}/cancel
```

### 🏥 Системные

#### Проверка состояния
```http
GET /avia/health
```

## Примеры использования

### Поиск рейса Ташкент-Стамбул

```bash
curl -X POST http://localhost:8080/avia/search \
  -H "Content-Type: application/json" \
  -d '{
    "directions": [
      {
        "departure_airport": "TAS",
        "arrival_airport": "IST",
        "date": "2024-01-15"
      }
    ],
    "service_class": "E",
    "adults": 1,
    "children": 0,
    "infants": 0,
    "infants_with_seat": 0
  }'
```

### Создание бронирования

```bash
curl -X POST http://localhost:8080/avia/offers/offer_id_123/booking \
  -H "Content-Type: application/json" \
  -d '{
    "payer_name": "Иван Иванов",
    "payer_email": "ivan@example.com",
    "payer_tel": "+998901234567",
    "passengers": [
      {
        "first_name": "Иван",
        "last_name": "Иванов",
        "middle_name": "Иванович",
        "age": "adt",
        "birthdate": "1990-01-01",
        "gender": "M",
        "citizenship": "UZ",
        "tel": "+998901234567",
        "doc_type": "A",
        "doc_number": "AA1234567",
        "doc_expire": "2030-01-01"
      }
    ]
  }'
```

### Оплата бронирования

```bash
curl -X POST http://localhost:8080/avia/booking/booking_id_456/payment
```

## Коды ошибок

API возвращает стандартные HTTP коды состояния:

- `200` - Успешный запрос
- `400` - Неверные параметры запроса
- `500` - Внутренняя ошибка сервера
- `503` - Сервис недоступен

## Ограничения

- Максимум 9 пассажиров на один запрос
- Количество младенцев не может превышать количество взрослых
- Токен Bukhara API автоматически обновляется каждые 28 дней
- Все запросы выполняются к тестовой среде Bukhara API

## Поддержка

При возникновении проблем обращайтесь к логам сервера или проверьте статус сервиса через `/avia/health`. 