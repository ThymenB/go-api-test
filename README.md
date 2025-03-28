# Recepten API  

Deze API maakt het mogelijk om recepten toe te voegen, op te halen en te verwijderen.  

## Installatie  

1. Zorg ervoor dat je Go hebt geïnstalleerd.  
2. clone deze repository.  
3. Open de terminal en ga naar de map van het project.  
4. Start de API met: `go run main.go`


## API Endpoints  

### 1. Recepten ophalen  
**http://localhost:8080** `/recipes`  
**Methode:** `GET`  
**Beschrijving:** Hiermee haal je een lijst op van alle recepten.  

### 2. Recept toevoegen  
**http://localhost:8080** `/recipes`  
**Methode:** `POST`  
**Beschrijving:** Hiermee voeg je een nieuw recept toe.  
**Benodigde gegevens (JSON-formaat):**  
```json
{
  "name": "NAAM",
  "ingredients": "JOUWINGEDIENTEN, JOUWINGEDIENTEN, JOUWINGEDIENTEN",
  "instructions": "JOUWINSTRUCTIES, JOUWINSTRUCTIES"
}
```  

### 3. Recept verwijderen  
**http://localhost:8080** `/recipe/{id}`  
**Methode:** `DELETE`  
**Beschrijving:** Hiermee verwijder je een recept op basis van het ID.  

## Extra informatie  

- De API draait standaard op `http://localhost:8080`.  
- Zorg ervoor dat je een tool zoals Postman of insomnia gebruikt om API-aanvragen te testen.  


## Database Importeren



Volg de onderstaande stappen om de geëxporteerde PostgreSQL-database in je lokale omgeving te importeren.

Verander DB_PASSWORD=je_database_wachtwoord naar je eigen wachtwoord van PostgreSQL

Vereisten:
## 1. PostgreSQL moet geïnstalleerd zijn. Als je PostgreSQL nog niet hebt geïnstalleerd, kun je dit downloaden van de officiële PostgreSQL website (https://www.postgresql.org/download/).
   
## 2. Het gedumpte bestand (`souschef_backup.dump`) moet beschikbaar zijn. Dit bestand bevat een back-up van de database en moet worden gedownload naar je lokale machine.

Stappen om de database te importeren:

1. Maak een nieuwe database aan:
   Open de commandoregel (Command Prompt of PowerShell) en log in op PostgreSQL:
   `psql -U postgres`

   Maak vervolgens een nieuwe database aan:
  `CREATE DATABASE souschef;`
   Vervang `souschef` door de naam van de database die je wilt gebruiken, als dit anders is.

## 2. Plaats het gedumpte bestand:
   Plaats het bestand `souschef_backup.dump` op een toegankelijke locatie op je computer, bijvoorbeeld op je bureaublad of in een andere map.

## 3. Implementeer de gedumpte database:
   Gebruik de `pg_restore` tool om de gedumpte gegevens in je nieuwe database te importeren. Voer het volgende commando uit in de commandoregel, waarbij je het pad naar het gedumpte bestand vervangt:
   `pg_restore -U postgres -d souschef "C:\path\to\souschef_backup.dump"`
   
   Vervang `"C:\path\to\souschef_backup.dump"` door het werkelijke pad naar het bestand op je computer (bijvoorbeeld `C:\Users\your_username\Desktop\souschef_backup.dump`).

## 4. Controleer de geïmporteerde gegevens:
   Nadat de import is voltooid, kun je de gegevens controleren door de database in te loggen:
   `psql -U postgres -d souschef`
   
   Voer vervolgens een query uit om te controleren of de gegevens correct zijn geïmporteerd, bijvoorbeeld:
   `SELECT * FROM recipes;`
   
   Vervang `recipes` door de naam van de tabel die je hebt geïmporteerd.

Problemen oplossen:
- Als je problemen ondervindt bij het uitvoeren van de bovenstaande stappen, zorg er dan voor dat PostgreSQL correct is geïnstalleerd en dat het pad naar het gedumpte bestand correct is opgegeven.
