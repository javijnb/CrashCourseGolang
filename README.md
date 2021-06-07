# Crash Course Golang - Traversy Media #

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
3. No tiene por qué indicarse un tipo de dato en una variable
4. Un atajo para declarar variables, pero solo en local:
> nombre := "Javi"
5. Para comprobar el formato de una variable:
> fmt.Printf("%T\n", nombre)

## 03 - Paquetes ##

1. Los paquetes se separan entre espacios, y los que no se usan, se eliminan solos
2. Creamos una función en otro paquete para invertir un string que reciba por parámetro

## 15 + 16 - Control Flow ##

Defer - panic - recover

1. La sentencia *defer* permite posponer una función para el final. Cuando acabe la ejecución normal del resto de funciones (main por ejemplo), antes de cerrar el programa, comprueba si hay funciones *defer* pendientes y las ejecuta. Sigue política LIFO, último en entrar, primero en salir - Pila. Este comportamiento se puede comprobar poniendo *defer* a todas las funciones
2. IMPORTANTE: la acción que vaya a ejecutar la sentencia que tenga *defer* será la misma que la que haría en la ejecución correcta, y no la ejecución que cabría esperar cuando haya acabado todo.
3. *Panic*: funciona de forma similar a lanzar excepciones de forma voluntaria en Java
4. Las sentencias *defer* son ejecutadas **después** de mostrar los mensaje de *panic*
5. *Recover*: se emplean para capturar *panics* y continuar con la ejecución correcta del programa
