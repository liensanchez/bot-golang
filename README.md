
# Payaso bot - Clown Bot!

<img align="right" alt="DiscordGo logo" src="https://i.ibb.co/ZKSZD5y/DALL-E-2023-05-10-15-01-13-make-me-a-pixel-art-clown.png" width="150">
Este proyecto es un bot de Discord que utiliza el paquete DiscordGo en Go. El bot se conecta a Discord y proporciona chistes, frases y refranes a los usuarios a través de una API personalizada. 

---


This project is a Discord bot that utilizes the DiscordGo package in Go. The bot connects to Discord and provides jokes, phrases, and proverbs to users through a custom API, currently available only in Spanish.



 
## Tech Stack

* Golang
* [DiscordGo](https://github.com/bwmarrin/discordgo)
* [JokesApi](https://github.com/liensanchez/rest-api)


## Uso/Ejemplo -Usage/Examples

#### Que diga una Frase - Get a phrase 
```http
  !frase
```
Solicita una frase aleatoria de una base de datos que contiene 85 frases.

Request a random phrase from a database that contains 85 phrases.

#### Que diga un Chiste - Get a Joke

```http
  !chiste
```
Solicita un chiste aleatorio de una base de datos que contiene 256 chistes.

Request a random joke from a database that contains 256 jokes.

#### Que diga un refran - Get a Proverb

```http
  !refran
```
Solicita un refran aleatorio de una base de datos que contiene 2780 refranes.

Request a random proverb from a database that contains 2780 proverbs.

