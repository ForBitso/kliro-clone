# 🛫 Тестирование API подсказок аэропортов

## 📍 **Endpoint: GET /avia/airport-hints**

### **Параметры:**
- `phrase` (обязательный) - поисковая фраза
- `limit` (опциональный) - максимальное количество результатов (по умолчанию 8)

---

## 🧪 **Примеры тестирования:**

### **1. Поиск по "tash" (Ташкент):**
```bash
curl -X GET "http://localhost:8080/avia/airport-hints?phrase=tash&limit=8" | jq '.'
```

### **2. Поиск по "dub" (Дубай):**
```bash
curl -X GET "http://localhost:8080/avia/airport-hints?phrase=dub&limit=5" | jq '.'
```

### **3. Поиск по "mos" (Москва):**
```bash
curl -X GET "http://localhost:8080/avia/airport-hints?phrase=mos&limit=10" | jq '.'
```

### **4. Поиск по "lon" (Лондон):**
```bash
curl -X GET "http://localhost:8080/avia/airport-hints?phrase=lon&limit=3" | jq '.'
```

### **5. Без параметра phrase (должна быть ошибка):**
```bash
curl -X GET "http://localhost:8080/avia/airport-hints" | jq '.'
```

---

## 🔍 **Ожидаемый ответ:**

```json
{
  "success": true,
  "result": [
    {
      "code": "TAS",
      "title": "Ташкент",
      "city": "Ташкент",
      "country": "Узбекистан"
    },
    {
      "code": "TAS",
      "title": "Ташкент Южный",
      "city": "Ташкент", 
      "country": "Узбекистан"
    }
  ],
  "message": "Найдено 2 подсказок аэропортов",
  "phrase": "tash",
  "limit": 8
}
```

---

## ⚠️ **Обработка ошибок:**

### **Ошибка валидации:**
```json
{
  "success": false,
  "error": "Параметр 'phrase' обязателен для поиска"
}
```

### **Ошибка API:**
```json
{
  "success": false,
  "error": "Ошибка получения подсказок: failed to get airport hints: ошибка API: ..."
}
```

---

## 🚀 **Использование во фронтенде:**

```javascript
// Пример автодополнения
async function searchAirports(query) {
  try {
    const response = await fetch(`/avia/airport-hints?phrase=${query}&limit=8`);
    const data = await response.json();
    
    if (data.success) {
      return data.result; // Массив аэропортов
    } else {
      console.error('Ошибка:', data.error);
      return [];
    }
  } catch (error) {
    console.error('Ошибка запроса:', error);
    return [];
  }
}

// Использование
searchAirports('tash').then(airports => {
  console.log('Найденные аэропорты:', airports);
});
```

---

## 📝 **Примечания:**

- API работает как прокси к Bukhara API
- Автоматически обновляет токен при необходимости
- Возвращает данные в том же формате что и Bukhara
- Поддерживает ограничение количества результатов
- Включает валидацию входных параметров 