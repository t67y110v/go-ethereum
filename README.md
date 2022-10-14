# CLI Приложение

## Tech Stack
- [Go](https://go.dev)
- [go-promt](https://pkg.go.dev/github.com/c-bata/go-prompt)
- [go-ethereum](https://github.com/ethereum/go-ethereum)
- [crypto/sha256](https://pkg.go.dev/crypto/sha256)
- [crypto/aes](https://pkg.go.dev/crypto/aes)

## Команды
Команды доступные в CLI

| Command | Description |
| :---:   | :---: |
| **`reload`** | Проверяет наличие созданных аккаунтов, и выводит их в консоль   | 
| **`newWallet [password]`** | Создает новый аккаунт, пароль хешируется crypto/sha256, и перед сохранением в файл шифруется с помощью crypto/aes  | 
| **`signIn [wallet] [password]`** |  Из списка аккаунтов выбирается [wallet] и вводится пароль. Пароль хешируется crypto/sha256, и шифруется с crypto/aes,  после чего сравнивается с паролем в файле и если они совпадают пользователю дается доступ к кошельку   | 
| **`exit`** | Выходит из CLI c кодом 0   | 


## Использование 
Для запуска использовать команды:

```go mod download```

```go run cmd/cli/main.go```

Также должен быть установлен [gcc](https://jmeubank.github.io/tdm-gcc/)

## Демонстрация
- При первом запуске аккаунтов нет, выводится предложение о создании нового аккаунта

![image](https://user-images.githubusercontent.com/46971653/195937502-09c42ce5-08a6-497f-a656-f7eeca348c7c.png)

- Вводим команду newWallet [password], и видим в списке появляется созданный кошелек, и предложение войти в систему

![image](https://user-images.githubusercontent.com/46971653/195938309-d4d37e2c-c1b2-49f3-977d-4666ab3cc0ba.png)

- Вводим команду signIn [wallet] [password] - пароли совпадают - получаем доступ и информацию о кошельке 

![image](https://user-images.githubusercontent.com/46971653/195942894-5f7191cf-38a5-4e97-92bb-78e4803838a7.png)


