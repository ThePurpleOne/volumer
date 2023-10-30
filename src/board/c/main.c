#include <stdio.h>
#include "pico/stdlib.h"
#include "hardware/gpio.h"
#include "hardware/adc.h"

int main() 
{
	stdio_init_all();
	printf("ADC Example, measuring multiple channels\n");

	adc_init();

	uint pins[] = {26, 27, 28, 29};  // Replace with the actual GPIO pins you want to use

	for (int i = 0; i < sizeof(pins) / sizeof(pins[0]); i++) 
		adc_gpio_init(pins[i]);

	while (1) 
	{
		int scaled_vals[sizeof(pins) / sizeof(pins[0])];

		for (int i = 0; i < sizeof(pins) / sizeof(pins[0]); i++) 
		{
			adc_select_input(i);  // Select the ADC input for this iteration
			scaled_vals[i] = adc_read() >> 2;
		}

		printf("%d|%d|%d|%d|%d\n", 	scaled_vals[0], 
									scaled_vals[1], 
									scaled_vals[2], 
									scaled_vals[3], 
									scaled_vals[0]);

		sleep_ms(10);
	}
}
