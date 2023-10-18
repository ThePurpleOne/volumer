import toml

from volume_control import *
from pcb_to_pc import list_available_ports, connect_to_pcb


def establish_connection():
	"""
	Establish a connection to the PCB and return the right serial object 
	"""
	ports = list_available_ports()
	pcb_ser, ok = connect_to_pcb(ports)

	if not ok:
		print("LOG : Could not connect to PCB")
		exit(-1)
	else:
		print(f"Connection to the volumer established on {pcb_ser.name}")

	return pcb_ser

def get_valid_data(ser):
	while True:
		data = ser.readline().decode('utf-8').strip()
		if data != "" and "RE" in data:
			print(data)
			return data


def main():
	show_open_apps()

	config = toml.load('config.toml')
	knob_apps = {
		"RE1": config.get("knob1").get("apps"),
		"RE2": config.get("knob2").get("apps"),
		"RE3": config.get("knob3").get("apps"),
	}

	ser = establish_connection()

	while True:
		data = get_valid_data(ser)
		print(data)

		knob, delta = data.split(":")

		if knob in knob_apps:
			if delta == "1":
				raise_volume_apps(knob_apps[knob])
			elif delta == "-1":
				lower_volume_apps(knob_apps[knob])
			else:
				print(f"LOG : Invalid delta {delta}")


if __name__ == '__main__':
	main()
