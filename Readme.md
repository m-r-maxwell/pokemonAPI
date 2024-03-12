# Pokémon REST API

This project is a GoLang REST API built with the latest version of `net/http`. It retrieves data from a PostgreSQL database containing information about Pokémon, abilities, moves, and items. Additionally, the API includes middleware to log incoming requests. 

## Disclaimer

**Important:** This project is purely for educational and non-commercial purposes. It is not affiliated with or endorsed by Pokémon, Nintendo, or any related entities. The use of Pokémon data is solely for demonstration and learning purposes.

## Features

- **Pokémon Data:** Retrieve information about various Pokémon species, including their types, abilities, base stats, and more.
- **Abilities:** Get details about different abilities that Pokémon can possess.
- **Moves:** Access data on moves that Pokémon can learn and use in battles.
- **Items:** Retrieve information about items that Pokémon can hold or use during gameplay.
- **Middleware:** Includes middleware to log incoming requests for debugging and monitoring purposes.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/pokemon-api.git
    cd pokemon-api
    ```
    NOTE: Python is included since that is what I used to generate the sql files and move and rename them for the database. They are there for people who are curious.

2. Set up the PostgreSQL database using Docker:

    ```bash
    docker build . -t database
    docker run database
    ```

3. Run the GoLang API:

    ```bash
    go run main.go
    ```
    -- docker compose hopefully later

## Usage

### Endpoints

- **Pokémon:**
    - `GET /pokemon` - Retrieve a list of all Pokémon.
    - `GET /pokemon/{id}` - Retrieve details about a specific Pokémon by ID.
    - `POST /pokemon` - Add a new Pokémon to the database.
    - `PUT /pokemon/{id}` - Update information about an existing Pokémon.
    - `DELETE /pokemon/{id}` - Delete a Pokémon from the database.

- **Abilities, Moves, and Items:** Similar endpoints to manage abilities, moves, and items.

### Logging

The middleware logs incoming requests to help with debugging and monitoring. Logs can be found in the `logs` directory.

## Future Enhancements

- **Swagger Documentation:** Add Swagger documentation to describe the API endpoints and their usage comprehensively.
- **Authentication:** Implement authentication mechanisms to secure the API endpoints.
- **Pagination and Filtering:** Enhance endpoints to support pagination and filtering for better data retrieval.

