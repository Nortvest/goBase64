## Сервис для Encode и Decode строк алгоритмом Base64

### Endpoints:
1. GET "/api/v1/encode" - Обязательный (единственный) параметр - "text".
   Возвращает JSON {"text": "<Введенный текст>", "result": "<Резельтат encode>"}

2. GET "/api/v1/decode" - Обязательный (единственный) параметр - "text".
   Возвращает JSON {"text": "<Введенный текст>", "result": "<Резельтат decode>"}
