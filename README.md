# Simple Rest API with golang

* para inicializar o modulo go em qualquer pasta
  * go mod init github.com/droveda/api-rest

  
* https://github.com/gorilla/mux
  * para instalar o gorrilamux: 
    * go get -u github.com/gorilla/mux
  * para resolver o problema de CORS
    * instalar: go get github.com/gorilla/handlers
* https://gorm.io/index.html
  * gorm golang ORM library
  * para instalar:
    * go get -u gorm.io/gorm
  * para instalar o driver do postgres
    * go get gorm.io/driver/postgres


# Database
* use the docker-compose.yaml file
* execute the following sql scripts
```sql
create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Deodato Petit Wertheimer', 'Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas.'),
('Carmela Dutra', 'Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.');
```

## Front End - React
* cd frontend-react-personalidades-master
* npm install
* npm update
* npm start