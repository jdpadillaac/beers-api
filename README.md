# Welcome to BeersApp!

En este README encontrará los pasos para ejecutar la aplicación

## Requerimientos
- Debe tener docker y docker compose instalado en su máquina.

## Pasos
- Debe crear un archivo .env en la raiz del proyecto con las siguientes variables

  `MONGO_CNN=string con conexion a base de datos en mongo`  
  `BASE_URL=/api/v1`   
  `CRLAYER_URL=http://api.currencylayer.com/`  
  `CRLAYER_KEY=your api access key`  
  `PORT=8101`   
  `MONGO_INITDB_DATABASE=beers_db`    
  `MONGO_INITDB_ROOT_USERNAME=user`    
  `MONGO_INITDB_ROOT_PASSWORD=password`  


- Ahora ejecute el comando en la terminal `docker-compose up -d --build`


### Consideraciones
- para conectarse a la base de datos local debe usar como valor en la variable `MONGO_CNN` la siguiente cadena `mongodb://user:password@mongoapp:27017/beers_db?authSource=admin` y remplazar user y password por los valores ingresados en su archivo .env


- User la acces key del api de currencylayer que se ha enviado en el archivo anexo, ya que esta está configurada con lo necesario para funcionar correctamente
