# Тестовое задание SkillsRock

## Локальный запуск
Скопировать проект
```bash
  git clone https://github.com/BabyJhon/skillsrock-test-task
```

Заполнить файлы:
- .env
- /configs/config.yml

Перейти в директорию с проектом
```bash
  cd skillsrock-test-task
```

Собрать и запустить проект
```bash
  docker-compose up test-task
```

Выполнить миграции БД(при необходимости изменить данные)
```bash
  migrate -path ./migrations -database postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable up
```

## Скриншоты работоспособности

- Создание задачи
  ![image](https://github.com/user-attachments/assets/3b94cf7a-6fe6-471c-8e7a-e4bc86ce5323)
  ![image](https://github.com/user-attachments/assets/ac948258-54c5-4d97-b5bc-d18704bcac27)
- Удаление задачи
  ![image](https://github.com/user-attachments/assets/8da83462-56c8-429a-9850-009f29265772)
  ![image](https://github.com/user-attachments/assets/f72d5725-81b0-41bf-ac10-a3f81f3ad06d)
- Редактирование задачи
  ![image](https://github.com/user-attachments/assets/b46014a2-b4de-48de-b864-af4bbd920202)
  ![image](https://github.com/user-attachments/assets/6e47a6a6-7aac-4b86-a8fb-125c8f0f7ef3)
- Вывод всех задач
  ![image](https://github.com/user-attachments/assets/ad966867-a492-4022-8ec3-d234b33c2cb3)




