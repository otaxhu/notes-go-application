# notes-go-application

## Bienvenidos a mi proyecto de Go

### Introduccion a este proyecto:

Este es un proyecto de API REST escrito en el lenguaje de programacion Go (Golang), el objetivo de este proyecto es mejorar mis habilidades y conocer mas acerca de este lenguaje en especifico, y saber mas acerca del Backend y todo lo que eso conlleva, estare usando varias librerias las cuales nombro a continuacion.

### Librerias:

1. [JWT](https://github.com/golang-jwt/jwt) (JSON Web Token para la Autenticacion de usuarios)
2. [UUID](https://github.com/google/uuid) (Universally Unique Identifier para la generacion de IDs aleatorios)
3. [gorilla/mux](https://github.com/gorilla/mux) (para los middlewares, el enrutamiento y el manejo de peticiones HTTP)
4. [godotenv](https://github.com/joho/godotenv) (para las variables de entorno las cuales se configuran en la carpeta environment/)
5. [crypto](https://golang.org/x/crypto) (para la encriptacion de las contraseñas y comparacion con la base de datos)
6. [GORM](https://gorm.io/gorm) (ORM para el mapeo de bases de datos relacionales orientado a objetos) me gusto mucho esta libreria, me parece interesante estudiarla
7. [GORM mysql driver](https://gorm.io/driver/mysql) (este es solo el driver mysql que viene con GORM)

### Caracteristicas de esta API REST

Cabe resaltar que en este punto aun no he implementado todas las funciones y caracteristicas que tengo en mente pero aun asi nombrare las que se me ocurrieron:

- Enrutador y manejador de peticiones HTTP
- Middleware de autenticacion usando los JWT, por ejemplo los usuarios que no tengan un token valido, solo podran acceder a las endpoints de /signup y /login
- Sistema de registro e inicio de sesion usando un email y contraseña implementando los JWT para la autenticacion
- Sistema de Creacion, Lectura, Actualizacion y Eliminacion de notas usando una base de datos mysql y el ORM de gorm. Cada usuario que se registre e inicie sesion, podra crear sus notas, pero al momento de querer interactuar con ellas, tendra que presentar su token en la cabecera de la peticion y solo podra leer, actualizar y borrar las notas que son de su propiedad (es decir las notas que creo usando su correo y contraseña)

### Como usar esta API?

Bueno, en realidad nose como se podria implementar en el frontend jaja, pero se que funciona con un cliente HTTP como Postman o Thunder Client.

1. GORM automaticamente configura las tablas de la base de datos asi que no debes preocuparte de las tablas, eso si, la base de datos debe ser mysql y tambien debes modificar las variables de entorno relacionadas a la base de datos para poder conectarte
2. Despues de haber configurado las variables de entorno, lo que sigue es ejecutar el programa e ir a tu cliente HTTP favorito
3. Vas a visitar la endpoint de /signup con el metodo "POST" y vas a poner tu email y contraseña de la siguiente manera:

![image](https://user-images.githubusercontent.com/93091522/220222470-56c68c39-7e58-49e4-932e-1c5da8bcf2d7.png)

4. Le das al boton de "Send" e inmediatamente te vas a la endpoint de /login y le das "Send" con los mismos datos y el mismo metodo "POST" (puedes probar a colocar datos incorrectos para testear la seguridad)

![image](https://user-images.githubusercontent.com/93091522/220223850-8248123b-21fb-4ba2-9ff0-71897364e1e4.png)

5. Como puedes observar en el response body has recibido tu token exitosamente, ahora debes copiar ese token y pegarlo en el header de "Authorization" de la siguiente manera:

![image](https://user-images.githubusercontent.com/93091522/220224726-e8356fa1-1f5d-4f19-89ac-89b692ab842e.png)

6. Ahora podras acceder a todas las endpoints disponibles en este backend, pero la que mas interesa endpoint de /api/notes puedes tratar de experimentar con esa endpoint todas las operaciones CRUD de las notas

### Gracias por visitar este repositorio!
