package cola

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T]{}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T]{}
