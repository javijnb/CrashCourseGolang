## 01 - Hola Mundo ##

1. Se necesita importar el paquete main para usar la función main
2. La función no es necesaria declararla y luego definirla como en C
3. Para I/O se usa el import FMT
4. Las variables que creemos SE TIENEN que usar

### Para compilar: ###

> go run main.go

### Link para referenciar según el tipo de la variable: ###

https://golang.org/pkg/fmt/

## 02 - Variables ##

1. Las hay variables y locales, tipadas o no tipadas explícitamente, variables o constantes
2. Los valores de las constantes no pueden ser modificados
3. No tiene por quñe indicarse un tipo de dato en una variable
4. Un atajo para declarar variables, pero solo en local:
> nombre := "Javi"
5. Para comprobar el formato de una variable:
> fmt.Printf("%T\n", nombre)
