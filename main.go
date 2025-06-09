// main.go
package main

import (
	"bufio"    // Para leer la entrada del usuario línea por línea
	"encoding/csv" // Para leer y escribir archivos CSV
	"fmt"      // Para formatear la salida (imprimir en consola)
	"io"       // Interfaz básica de I/O
	"os"       // Para interactuar con el sistema operativo (archivos)
	"strconv"  // Para convertir strings a números y viceversa
	"strings"  // Para manipular strings
)

const inventoryFile = "inventory.csv" // Nombre del archivo CSV para el inventario

// Product representa un artículo en el inventario
type Product struct {
	Name     string
	Quantity int
	Price    float64
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Lector para la entrada del usuario

	for {
		fmt.Println("\n--- Gestor de Inventario ---")
		fmt.Println("1. Añadir producto")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Actualizar producto")
		fmt.Println("4. Eliminar producto")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		input, _ := reader.ReadString('\n') // Lee la línea completa hasta el salto de línea
		choice := strings.TrimSpace(input)  // Elimina espacios en blanco y el salto de línea

		switch choice {
		case "1":
			addProduct(reader)
		case "2":
			listProducts()
		case "3":
			updateProduct(reader)
		case "4":
			deleteProduct(reader)
		case "5":
			fmt.Println("¡Adiós!")
			return // Sale del programa
		default:
			fmt.Println("Opción inválida. Por favor, intenta de nuevo.")
		}
	}
}

// loadProducts carga los productos desde el archivo CSV
func loadProducts() ([]Product, error) {
	file, err := os.OpenFile(inventoryFile, os.O_RDONLY|os.O_CREATE, 0644) // Abre el archivo en modo lectura, lo crea si no existe
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo de inventario: %w", err)
	}
	defer file.Close() // Asegura que el archivo se cierre al finalizar la función

	reader := csv.NewReader(file)
	// Opcional: Si tu CSV tiene un encabezado y quieres omitirlo
	// _, err = reader.Read() // Lee la primera línea (encabezado)
	// if err != nil && err != io.EOF {
	// 	return nil, fmt.Errorf("error al leer encabezado CSV: %w", err)
	// }

	var products []Product
	for {
		record, err := reader.Read() // Lee una fila del CSV
		if err == io.EOF {
			break // Se llegó al final del archivo
		}
		if err != nil {
			return nil, fmt.Errorf("error al leer registro CSV: %w", err)
		}

		if len(record) != 3 { // Asegúrate de que la fila tiene 3 columnas esperadas
			fmt.Printf("Advertencia: Fila CSV mal formada, se esperaba 3 columnas, se encontró %d: %v\n", len(record), record)
			continue // Salta esta fila y continúa con la siguiente
		}

		quantity, err := strconv.Atoi(record[1]) // Convierte la cantidad de string a int
		if err != nil {
			fmt.Printf("Advertencia: Cantidad inválida '%s' en el registro: %v\n", record[1], record)
			continue
		}
		price, err := strconv.ParseFloat(record[2], 64) // Convierte el precio de string a float64
		if err != nil {
			fmt.Printf("Advertencia: Precio inválido '%s' en el registro: %v\n", record[2], record)
			continue
		}

		products = append(products, Product{
			Name:     record[0],
			Quantity: quantity,
			Price:    price,
		})
	}
	return products, nil
}

// saveProducts guarda la lista actual de productos en el archivo CSV
func saveProducts(products []Product) error {
	file, err := os.Create(inventoryFile) // Crea/sobrescribe el archivo
	if err != nil {
		return fmt.Errorf("error al crear/sobrescribir archivo de inventario: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // Asegura que los datos se escriban al disco

	// Opcional: Escribe un encabezado CSV
	// if err := writer.Write([]string{"Nombre", "Cantidad", "Precio"}); err != nil {
	// 	return fmt.Errorf("error al escribir encabezado CSV: %w", err)
	// }

	for _, p := range products {
		record := []string{
			p.Name,
			strconv.Itoa(p.Quantity),    // Convierte int a string
			strconv.FormatFloat(p.Price, 'f', 2, 64), // Convierte float64 a string con 2 decimales
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error al escribir registro CSV: %w", err)
		}
	}
	return nil
}

// addProduct permite al usuario añadir un nuevo producto
func addProduct(reader *bufio.Reader) {
	fmt.Print("Introduce el nombre del producto: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	products, err := loadProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Comprobar si el producto ya existe (por nombre)
	for _, p := range products {
		if strings.EqualFold(p.Name, name) { // EqualFold ignora mayúsculas/minúsculas
			fmt.Println("Error: Ya existe un producto con ese nombre. Usa 'Actualizar' si quieres modificarlo.")
			return
		}
	}

	fmt.Print("Introduce la cantidad: ")
	quantityStr, _ := reader.ReadString('\n')
	quantity, err := strconv.Atoi(strings.TrimSpace(quantityStr))
	if err != nil || quantity < 0 {
		fmt.Println("Cantidad inválida. Debe ser un número entero positivo.")
		return
	}

	fmt.Print("Introduce el precio: ")
	priceStr, _ := reader.ReadString('\n')
	price, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err != nil || price < 0 {
		fmt.Println("Precio inválido. Debe ser un número positivo.")
		return
	}

	products = append(products, Product{Name: name, Quantity: quantity, Price: price})
	if err := saveProducts(products); err != nil {
		fmt.Println("Error al guardar producto:", err)
		return
	}
	fmt.Println("Producto añadido con éxito.")
}

// listProducts muestra todos los productos en el inventario
func listProducts() {
	products, err := loadProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(products) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	fmt.Println("\n--- Inventario Actual ---")
	fmt.Printf("%-20s %-10s %-10s\n", "Nombre", "Cantidad", "Precio")
	fmt.Println(strings.Repeat("-", 42)) // Línea separadora
	for _, p := range products {
		fmt.Printf("%-20s %-10d %-10.2f\n", p.Name, p.Quantity, p.Price)
	}
}

// findProductIndex busca un producto por nombre y devuelve su índice
func findProductIndex(products []Product, name string) int {
	for i, p := range products {
		if strings.EqualFold(p.Name, name) {
			return i
		}
	}
	return -1 // No encontrado
}

// updateProduct permite al usuario actualizar un producto existente
func updateProduct(reader *bufio.Reader) {
	fmt.Print("Introduce el nombre del producto a actualizar: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	products, err := loadProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	index := findProductIndex(products, name)
	if index == -1 {
		fmt.Println("Producto no encontrado.")
		return
	}

	fmt.Println("Producto encontrado:", products[index].Name)
	fmt.Print("Introduce la nueva cantidad (deja en blanco para no cambiar): ")
	newQuantityStr, _ := reader.ReadString('\n')
	newQuantityStr = strings.TrimSpace(newQuantityStr)
	if newQuantityStr != "" {
		newQuantity, err := strconv.Atoi(newQuantityStr)
		if err != nil || newQuantity < 0 {
			fmt.Println("Cantidad inválida. No se actualizó la cantidad.")
		} else {
			products[index].Quantity = newQuantity
		}
	}

	fmt.Print("Introduce el nuevo precio (deja en blanco para no cambiar): ")
	newPriceStr, _ := reader.ReadString('\n')
	newPriceStr = strings.TrimSpace(newPriceStr)
	if newPriceStr != "" {
		newPrice, err := strconv.ParseFloat(newPriceStr, 64)
		if err != nil || newPrice < 0 {
			fmt.Println("Precio inválido. No se actualizó el precio.")
		} else {
			products[index].Price = newPrice
		}
	}

	if err := saveProducts(products); err != nil {
		fmt.Println("Error al actualizar producto:", err)
		return
	}
	fmt.Println("Producto actualizado con éxito.")
}

// deleteProduct permite al usuario eliminar un producto
func deleteProduct(reader *bufio.Reader) {
	fmt.Print("Introduce el nombre del producto a eliminar: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	products, err := loadProducts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	index := findProductIndex(products, name)
	if index == -1 {
		fmt.Println("Producto no encontrado.")
		return
	}

	// Eliminar el producto usando slicing
	products = append(products[:index], products[index+1:]...)

	if err := saveProducts(products); err != nil {
		fmt.Println("Error al eliminar producto:", err)
		return
	}
	fmt.Println("Producto eliminado con éxito.")
}