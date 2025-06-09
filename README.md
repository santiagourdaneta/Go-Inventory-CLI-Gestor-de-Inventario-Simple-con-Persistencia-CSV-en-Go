
Go-Inventory-CLI es una aplicación de consola desarrollada en Go que te permite gestionar un inventario básico de productos. Este programa facilita la adición, listado, actualización y eliminación de productos, manteniendo toda la información persistente gracias al uso de archivos CSV (Comma Separated Values). Es una herramienta ideal para quienes buscan una solución ligera para el control de inventario o para estudiantes que desean explorar el manejo de archivos, estructuras de datos y la interactividad básica de consola en Go.

# Go-Inventory-CLI: Gestor de Inventario Simple con Persistencia CSV en Go

## 🚀 Descripción del Proyecto

**Go-Inventory-CLI** es una aplicación de línea de comandos diseñada para ayudarte a gestionar un inventario de productos de manera sencilla y eficiente. Desarrollada en **Go**, esta herramienta permite a los usuarios añadir nuevos productos, listar los existentes, actualizar sus cantidades o precios, y eliminar entradas no deseadas. Todos los datos del inventario se guardan y cargan desde un archivo **CSV**, asegurando la persistencia de la información entre sesiones. Es un proyecto excelente para familiarizarse con las capacidades de Go en el manejo de archivos y la interactividad de consola.

## ✨ Características

* **Gestión CRUD:**
    * **Crear:** Añade nuevos productos al inventario (nombre, cantidad, precio).
    * **Leer:** Lista todos los productos disponibles con sus detalles.
    * **Actualizar:** Modifica la cantidad o el precio de un producto existente por su nombre.
    * **Eliminar:** Borra un producto del inventario por su nombre.
* **Persistencia de Datos:** Todos los cambios se guardan automáticamente en un archivo CSV (`inventory.csv`) y se cargan al iniciar la aplicación.
* **Interfaz de Consola:** Interacción simple y directa a través de la terminal.
* **Manejo de Errores:** Incluye validaciones básicas para la entrada de usuario y el manejo de archivos.

## 🛠️ Tecnologías Utilizadas

* **Go Lang** (versión 1.16 o superior recomendada)

## 🚀 Cómo Ejecutar el Proyecto

Sigue estos pasos para poner en marcha el Gestor de Inventario:

### Localmente (Si tienes Go instalado en tu máquina):

1.  Asegúrate de tener **Go** instalado en tu sistema. Puedes descargarlo desde el [sitio web oficial de Go](https://go.dev/dl/).
2.  Crea una carpeta para tu proyecto (ej. `inventario_go`).
3.  Dentro de esa carpeta, crea un archivo `main.go` y pega el código fuente proporcionado.
4.  Abre tu línea de comandos o terminal (CMD/PowerShell en Windows, Terminal en macOS/Linux).
5.  Navega hasta el directorio de tu proyecto:
   
    cd ruta/a/tu/inventario_go/
  
6.  **Ejecuta** la aplicación:
    
    go run main.go
    
    El programa se iniciará en tu terminal.

## 📁 Estructura del Archivo CSV

El programa crea y utiliza un archivo llamado `inventory.csv` en el mismo directorio donde se ejecuta el programa. Cada línea de este archivo representa un producto con sus valores separados por comas, en el siguiente orden:

`NombreDelProducto,Cantidad,Precio`

**Ejemplo:**

Manzanas,100,0.75
Leche,50,1.20
Pan,20,2.50

¡Genial! Aquí tienes el título, la descripción y un archivo README.md completo para tu Gestor de Inventario Simple en Go. Esto te servirá para documentar tu proyecto y compartirlo de forma clara.

Gestor de Inventario Simple (Go con CSV)
Título del Proyecto:
Go-Inventory-CLI: Gestor de Inventario Simple con Persistencia CSV en Go

Descripción del Proyecto:
Go-Inventory-CLI es una aplicación de consola desarrollada en Go que te permite gestionar un inventario básico de productos. Este programa facilita la adición, listado, actualización y eliminación de productos, manteniendo toda la información persistente gracias al uso de archivos CSV (Comma Separated Values). Es una herramienta ideal para quienes buscan una solución ligera para el control de inventario o para estudiantes que desean explorar el manejo de archivos, estructuras de datos y la interactividad básica de consola en Go.

README.md para Go-Inventory-CLI:
Markdown

# Go-Inventory-CLI: Gestor de Inventario Simple con Persistencia CSV en Go

## 🚀 Descripción del Proyecto

**Go-Inventory-CLI** es una aplicación de línea de comandos diseñada para ayudarte a gestionar un inventario de productos de manera sencilla y eficiente. Desarrollada en **Go**, esta herramienta permite a los usuarios añadir nuevos productos, listar los existentes, actualizar sus cantidades o precios, y eliminar entradas no deseadas. Todos los datos del inventario se guardan y cargan desde un archivo **CSV**, asegurando la persistencia de la información entre sesiones. Es un proyecto excelente para familiarizarse con las capacidades de Go en el manejo de archivos y la interactividad de consola.

## ✨ Características

* **Gestión CRUD:**
    * **Crear:** Añade nuevos productos al inventario (nombre, cantidad, precio).
    * **Leer:** Lista todos los productos disponibles con sus detalles.
    * **Actualizar:** Modifica la cantidad o el precio de un producto existente por su nombre.
    * **Eliminar:** Borra un producto del inventario por su nombre.
* **Persistencia de Datos:** Todos los cambios se guardan automáticamente en un archivo CSV (`inventory.csv`) y se cargan al iniciar la aplicación.
* **Interfaz de Consola:** Interacción simple y directa a través de la terminal.
* **Manejo de Errores:** Incluye validaciones básicas para la entrada de usuario y el manejo de archivos.

## 🛠️ Tecnologías Utilizadas

* **Go Lang** (versión 1.16 o superior recomendada)

## 🚀 Cómo Ejecutar el Proyecto

Sigue estos pasos para poner en marcha el Gestor de Inventario:

### En Codespaces (Recomendado para un inicio rápido):

1.  **Asegúrate de tener un Codespace configurado para Go.** Si no lo tienes, añade un archivo `.devcontainer/devcontainer.json` similar a este en la raíz de tu repositorio y **reconstruye tu Codespace**:

    ```json
    // .devcontainer/devcontainer.json
    {
      "name": "Go Inventory App",
      "image": "[mcr.microsoft.com/devcontainers/go:0-1.22](https://mcr.microsoft.com/devcontainers/go:0-1.22)", // Puedes usar una versión de Go más reciente
      "features": {
        "ghcr.io/devcontainers/features/git:1": {
          "version": "latest"
        }
      },
      "customizations": {
        "vscode": {
          "extensions": [
            "golang.go"
          ]
        }
      },
      "remoteUser": "vscode"
    }
    ```
2.  Una vez que tu Codespace esté activo, crea una carpeta llamada `inventario_go/` y, dentro de ella, un archivo `main.go`.
3.  Copia y pega el código fuente completo de `main.go` en ese archivo.
4.  Abre la **terminal** integrada en VS Code (puedes ir a `Terminal` > `New Terminal`).
5.  Navega hasta el directorio de tu proyecto:
    ```bash
    cd inventario_go/
    ```
6.  **Ejecuta** la aplicación:
    ```bash
    go run main.go
    ```
    El gestor de inventario se iniciará y te presentará el menú principal en la terminal.

### Localmente (Si tienes Go instalado en tu máquina):

1.  Asegúrate de tener **Go** instalado en tu sistema. Puedes descargarlo desde el [sitio web oficial de Go](https://go.dev/dl/).
2.  Crea una carpeta para tu proyecto (ej. `inventario_go`).
3.  Dentro de esa carpeta, crea un archivo `main.go` y pega el código fuente proporcionado.
4.  Abre tu línea de comandos o terminal (CMD/PowerShell en Windows, Terminal en macOS/Linux).
5.  Navega hasta el directorio de tu proyecto:
    ```bash
    cd ruta/a/tu/inventario_go/
    ```
6.  **Ejecuta** la aplicación:
    ```bash
    go run main.go
    ```
    El programa se iniciará en tu terminal.

## 📁 Estructura del Archivo CSV

El programa crea y utiliza un archivo llamado `inventory.csv` en el mismo directorio donde se ejecuta el programa. Cada línea de este archivo representa un producto con sus valores separados por comas, en el siguiente orden:

`NombreDelProducto,Cantidad,Precio`

**Ejemplo:**
```csv
Manzanas,100,0.75
Leche,50,1.20
Pan,20,2.50

💡 Uso
Al ejecutar el programa, se te presentará un menú de opciones. Simplemente ingresa el número correspondiente a la acción que deseas realizar y sigue las instrucciones en pantalla.

🤝 Contribuciones
Las contribuciones son bienvenidas. Si tienes ideas para mejorar este gestor de inventario (por ejemplo, añadir más campos, implementar una interfaz de usuario web, etc.), no dudes en abrir un "issue" o enviar un "pull request" en el repositorio.
