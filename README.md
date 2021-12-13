# Welcome to BeersApp!

En este README encontrará los pasos para ejecutar la aplicación

## Requerimientos
- Debe tener docker y docker compose instalado en su máquina.

## Pasos
- Debe crear un archivo .env en la raiz del proyecto con las siguientes variables



  `MONGO_CNN=string con conexión a base de datos en mongo`

  `BASE_URL=/api/v1`  

  `CRLAYER_URL=http://api.currencylayer.com/`  

  `CRLAYER_KEY=your api key`  

  `PORT=8101`  

  `MONGO_INITDB_DATABASE=beers_db`  

  `MONGO_INITDB_ROOT_USERNAME=useradmin`  

  `MONGO_INITDB_ROOT_PASSWORD=example`



- Ahora ejecute el comando en la terminal `docker-compose up -d --build`
