# circuit_python program ran on the RP2040
# This creates an HID device as custom
# And can send a custom report to the host
# NO keyboard or mouse, just a CONSUMER_CONTROL device


import usb_hid
from adafruit_hid.consumer_control import ConsumerControl
from adafruit_hid.consumer_control_code import ConsumerControlCode
from time import sleep
import usb_hid


# Disable every default HID device
# This is to prevent the RP2040 from being a keyboard or mouse
usb_hid.disable()

# Custom HID report descriptor for a single-byte report.
CUSTOM_REPORT_DESCRIPTOR = bytes((
	0x06, 0x00, 0xff,   # USAGE_PAGE (Vendor Defined Page 1)
	0x09, 0x01,         # USAGE (Vendor Usage 1)
	0xA1, 0x01,         # COLLECTION (Application)
	0x85, 0x04,         #   Report ID (4)
	0x75, 0x08,         #   REPORT_SIZE (8 bits)
	0x95, 0x01,         #   REPORT_COUNT (1 byte)
	0x15, 0x00,         #   LOGICAL_MINIMUM (0)
	0x26, 0xff, 0x00,   #   LOGICAL_MAXIMUM (255)
	0x09, 0x01,         #   USAGE (Vendor Usage 1)
	0x81, 0x02,         #   INPUT (Data,Var,Abs)
	0xC0                # END_COLLECTION
))

custom_hid = usb_hid.Device(
	report_descriptor=CUSTOM_REPORT_DESCRIPTOR,
	usage_page=0x1234,      # Vendor Defined Page 1
	usage=0x01,             # Vendor Usage 1
	report_ids=(4,),        # Descriptor uses report ID 4.
	in_report_lengths=(1,), # This custom HID device sends 1 byte in its report.
	out_report_lengths=(0,),# It does not receive any reports.
	#manufacturer_string="The Purple Company",
	#product_string="Volumer",
)

while True:
	print("yep")
	custom_hid.send_report(b'\x55', 4)
	sleep(0.1)