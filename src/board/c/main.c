#include <stdio.h>
#include "pico/stdlib.h"
#include "pico/binary_info.h"
#include "pico/unique_id.h"
#include "hardware/gpio.h"
//#include "hardware/irq.h"
//#include "hardware/pio.h"
#include "usb/usb.h"
#include "usb/hid.h"

// Define the VID and PID for your custom HID device
#define VENDOR_ID 0x1234
#define PRODUCT_ID 0x5678

// Custom report IDs
#define CUSTOM_REPORT_ID_1 0x01
#define CUSTOM_REPORT_ID_2 0x02
#define CUSTOM_REPORT_ID_3 0x03
#define CUSTOM_REPORT_ID_4 0x04

void usb_hid_init() 
{
    stdio_init_all();

    gpio_set_function(0, GPIO_FUNC_USB);
    gpio_set_function(1, GPIO_FUNC_USB);

    if (usb_init()) 
	{
        printf("USB init failed\n");
        return;
    }

    if (usb_add_string_descriptor("RP2040 Volumer", 14, 1)) 
	{
        printf("Failed to add string descriptor\n");
        return;
    }

    usb_hid_set_report_descriptor(NULL, 0); // No custom descriptor for simplicity


	// Register the custom report IDs
    if (usb_start(VENDOR_ID, PRODUCT_ID, "The Purple Company", "volumer", "12345", "1.0", usb_default_handle_interrupt)) 
	{
        printf("Failed to start USB\n");
        return;
    }
}

void send_custom_message(uint8_t report_id, uint8_t* data, size_t length) 
{
    if (length > 64) 
	{
        printf("Data length exceeds report size\n");
        return;
    }

    hid_message_t message = {
        .report_id = report_id,
        .length = length,
    };

    memcpy(message.data, data, length);
    usb_hid_send_message(&message);
}

int main() 
{
    bi_decl(bi_program_name("Custom HID Program"));

    usb_hid_init();

    while (1) 
	{
        uint8_t msg1[] = "UP";
        uint8_t msg2[] = "DN";
        uint8_t msg3[] = "LF";
        uint8_t msg4[] = "RI";

        send_custom_message(CUSTOM_REPORT_ID_1, msg1, sizeof(msg1));
        //send_custom_message(CUSTOM_REPORT_ID_2, msg2, sizeof(msg2));
        //send_custom_message(CUSTOM_REPORT_ID_3, msg3, sizeof(msg3));
        //send_custom_message(CUSTOM_REPORT_ID_4, msg4, sizeof(msg4));

        sleep_ms(1000);
    }

    return 0;
}
