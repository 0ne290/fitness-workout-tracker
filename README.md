## Памятка (привет, тавтология):
Я использую errors исключительно для отдачи клиентам в том виде, в каком они были созданы. Я не использую их для отлова ошибок разработки - для этого я использую panic.

## TODO:
1. Web API с маршрутизацией и авторизацией по ролям через JWT (использовать заголовок Authorization: Bearer {token}).
2. Добавить полноценный Application Layer с неким подобием CQRS. На самом деле основная задача не в том, чтобы разделять Use Cases на Reading и Writing типы, а в том, чтобы выделять каждый Use Case в отдельную самостоятельну сущность, изолированную от других Use Cases и четко разделенную на данные и поведение - таким образом код будет более структурирован и менее размазан.
3. Логирование в Elastick Search. метрики в Prometheus. Ко всему этому добру прицепить Grafana.