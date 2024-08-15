# Proyecto de practica de Clean Architecture

Hola mundo en go, proyecto de un Todo List con el fin de aprender Clean Architecture
y comenzar aprender Go
##  Swagger
> [!NOTE]
> Endpoint `/swagger/`
> Comando `swag init -g /main.go -d cmd,internal/entity,pkg/adapter/handler`
## Distribución de carpetas

```
└── 📁GO.TodoList
    └── .gitignore
    └── 📁cmd
        ├── server
    └── go.mod
    └── 📁internal
        └── 📁domain
            └── 📁entity
                └── task.go
        └── 📁repository
            └── task_repository.go
        └── 📁usecase
            └── task_usecase.go
    └── 📁pkg
        └── 📁adapter
            └── 📁db
                └── posgrest_repository.go
            ├── http
        ├── shared
    └── README.md
```
- **cmd** : La carpeta es la que levanta el servidor
- **internal** : La carpeta contiene la parte interna del sistema el *domain, usecase y repositorios* de la app
    - **Domain** : Dominio pose todas las entidades del app
    - **UseCase** : Casos de uso(logica de negocio)
    - **Repository** : Interfaces del repositorio
- **pkg** : La carpeta contiene lo que considero las partes que interactuan con lo externo del app eso quiere decir los *adapters he interfaces externas(UI, Web, DB, Devices)*
    - **adapter** : Contiene adaptades de la base de datos y de API HTTP
    - **shared** : Código compartido (helpers, utilidades)

## Nomenclatura de commits

| Description                          | Type     | Format | Example                                    |
|--------------------------------------|----------|--------|--------------------------------------------|
| **Requirements and Features**        | feat     | `:sparkles:`    | feat: :sparkles: Include new feature       |
| **Change Control**                   | feat     | `:boom:`        | feat: :boom: Service implementation        |
| **Defects and Incidents**            | fix      | `:construction:`| fix: :construction: Mapping is corrected   |
| **Fix bugs**                         | fix      | `:bug:`         | fix: :bug: Mapping order fix               |
| **Immediate correction is required** | fix      | `:ambulance:`   | fix: :ambulance: Fix flow bug              |
| **Phase or sprint implemented**      | feat     | `:package:`     | feat: :package: Feature is included        |
| **Add, update or pass tests**        | test     | `:white_check_mark:` | test: :white_check_mark: New tests added |
| **Add or update documentation**      | docs     | `:memo:`        | docs: :memo: Update doc                    |
| **Add or update UI styles**          | style    | `:lipstick:`    | style: :lipstick: Update UI                |
| **Write bad code needed review**     | refactor | `:poop:`        | refactor: :poop: Fix this please |
| **Remove files**                     | feat     | `:fire:`        | feat: :fire: Remove file                   |
| **Reverting changes**                | revert   | `:rewind:`      | revert: :rewind: I shouldn't do that again |
| **Improving Performance**            | perf     | `:zap:`         | perf: :zap: Optimizing code                |


Complejiodad algorimica es cuanto tiempo se demora en resolver el problema en base a la cantidad de datos