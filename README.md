# djangolang
djangolang

## how to run project without docker
1. Create the following files/folders in root folder
    * file : (.env), 
2. Add environment variables in .env file
    1. DB_URL
3. Install all dependencies (command: go mod download)
4. Run project (command: go run main.go)

## reference for tech documentation
* List each endpoint in file endpoints.md
* List each database tables in file dbtables.md
* Database script files are located in folder path resources/sql