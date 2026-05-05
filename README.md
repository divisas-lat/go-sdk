# Divisas.lat Go SDK

El SDK oficial para interactuar con la API de Divisas.lat utilizando Go (Golang).
Este SDK proporciona un cliente idiomático, fuertemente tipado ("Zero Dependencies" más allá de la biblioteca estándar) y seguro para concurrencia.

## Requisitos
- Go 1.21 o superior

## Instalación

```bash
go get github.com/divisas-lat/go-sdk
```

## Características
- **Tipado Fuerte**: Modelos nativos (`structs`) y Tipos Fuertes para parámetros de países y monedas.
- **Fluent API Builder**: Interfaz intuitiva y encadenable (`divisas.Query().ForCountry()`).
- **Cero Dependencias**: Usa nativamente `net/http` y `encoding/json`.
- **Caché Concurrente**: Incluye un mecanismo de caché en memoria de alto rendimiento y `thread-safe`.
- **Soporte Context**: Totalmente compatible con `context.Context` de Go para timeouts y cancelación.

## Uso Básico

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/divisas-lat/go-sdk"
	"github.com/divisas-lat/go-sdk/enums"
)

func main() {
	// Inicializa el cliente. 
	// Automáticamente leerá la variable de entorno DIVISAS_API_KEY
	client := divisas.NewClient(
		divisas.WithAPIKey("tu_api_key_aqui"),
	)

	ctx := context.Background()

	// Obtener el tipo de cambio de hoy para Guatemala
	rates, err := client.Query().
		ForCountry(enums.Guatemala).
		GetToday(ctx)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Moneda: %s, Compra: %.2f, Venta: %.2f\n", 
		rates.Rate.CurrencyCode, 
		rates.Rate.Buy, 
		rates.Rate.Sell,
	)

	// Conversión rápida
	conversion, err := client.Query().
		ForCountry(enums.Mexico).
		WithCurrency(enums.USD).
		Convert(ctx, enums.MXN, 100.50)
		
	if err == nil {
		fmt.Printf("100.50 USD son %.2f MXN\n", conversion.Result)
	}
}
```

## Herramienta CLI

El SDK también incluye una utilidad de consola lista para usar:

```bash
# Instalar globalmente
go install github.com/divisas-lat/go-sdk/cmd/divisas@latest

# Obtener tipo de cambio de Guatemala
divisas today GT USD

# Convertir 100 Dólares a Quetzales
divisas convert 100 USD to GTQ in GT
```
