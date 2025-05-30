# Sistema de Usuarios en Go con Gin

Este proyecto es una API REST construida con el lenguaje de programación **Go** y el framework **Gin**. Está diseñado para gestionar usuarios con diferentes roles y permisos, como Superadministrador, Administrador, Investigador, Usuario, Anónimo, Miembro de Comunidad y Administrador de Comunidad.

## 📂 Estructura del Proyecto

## Estructura del proyecto
- `cmd/`: Punto de entrada del programa.
- `config/`: Configuración de la base de datos.
- `controllers/`: Controladores para autenticación y usuarios.
- `middleware/`: Autenticación JWT y control de roles.
- `models/`: Definición de entidades.
- `routes/`: Rutas y agrupación por permisos.
- `services/`: Lógica de negocio.
- `utils/`: Utilidades para JWT.


## 🧪 Funcionalidades

- Login de usuario con JWT
- Creación de usuarios por roles
- Listado y obtención de usuarios


## 👤 Roles de Usuario

- **Superadministrador:** Crea administradores
- **Administrador:** Crea, visualiza, elimina usuarios, bloquea/desbloquea usuarios y reasigna contraseñas
- **Otros roles:** Acceso restringido según su perfil (Investigador, Comunidad, Anónimo, etc.)

## ⚙️ Instalación y Ejecución

### Requisitos

- Go 1.20 o superior
- Docker y Docker Compose
- PostgreSQL

## Instrucciones

Para iniciar el proyecto, ejecutar los siguientes comandos:
```bash
docker-compose up -d

go mod tidy

go run cmd/main.go
```
### Insertar usuario principal

- INSERT INTO users (name,email, password, role) VALUES ('admin','pepe@gmail.com', 'admin123', 'admin');


### Seguridad
Autenticación basada en JWT



### Construido con
Go

Gin Web Framework

PostgreSQL

Docker