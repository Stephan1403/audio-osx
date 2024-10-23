package audio

/*
#cgo CFLAGS: -I.
#include <stdio.h>

// Function to simulate listing audio devices
int list_devices() {
	return 42;
}

// Function to simulate changing the volume
void set_volume(int volume) {
    printf("Volume set to: %d\n", volume);
}
*/
import "C"

func GetFoo() string {
	return string(C.list_devices())
}
