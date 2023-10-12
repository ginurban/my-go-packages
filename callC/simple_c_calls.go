package callC

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(int x) {
//   printf("integer has value %d\n", x);
// }
// static void add_one(int *x){
//	*x = *x + 1;
// }
import "C"

func CallC() {
	x := C.int(4)
	C.myprint(x)
	C.add_one(&x)
	C.myprint(x)
}
