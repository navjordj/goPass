# GoPass

Simple CLI password manager written in GO. Saves AES256 encrypted passwords in a database locally.


## Installation
```console
git clone https://github.com/navjordj/goPass
cd goPass
go install
```

## Usage


### Insert 
```console
goPass insert
```

### Get Password 
```console
goPass get
```

### Update Password 
```console
goPass update
```



## Ågjøre
 - [X] Generer innsatt passord random (insert.go)
 - [X] Hente ut passord fra databasen og decrypte
 - [X] Teste i praksis
 - [X] Skrive en README 
 - [X] Finne et heftig navn
 - [ ] Samme passord for alt (masterpassord)
 - [ ] La bruker bestemme passordinstillinger (stor bokstav, lengde osv.)
 - [X] Legge til mail
 - [X] Mulighet til å oppdatere passord
 - [ ] Mulighet til å oppdatere email
 - [ ] Slette passord i databasen
 - [ ] Droppe sqlite for postgres
 - [ ] Skrive passord til clipboard (https://godoc.org/github.com/atotto/clipboard) 
 - [ ] Fjerne cobratekst
 - [ ] Tester
 - [ ] Eksporter all data
 - [ ] Importer data
 - [ ] Error handling