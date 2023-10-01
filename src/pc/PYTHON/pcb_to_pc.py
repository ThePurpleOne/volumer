# This file contains the serial communication between the PCB and the PC
import serial
import time


def list_available_ports():
	import serial.tools.list_ports
	ports = serial.tools.list_ports.comports()

	out = []
	for port, _, _ in sorted(ports):
		out.append(f"{port}")
	return out

def conect_to_port(port):
	ser = serial.Serial(port, 115200, timeout=2)
	return ser

def wait_for_connection(ser) -> bool:
	"""
	The PCB should spam PING until it receives PONG from the PC
	so we'll check for a PING and respond with a PONG
	(It only checks if it contains "PING" not an exact match, to avoid crlf issues)
	"""

	for _ in range(2):
		if ser.in_waiting:
			read = ser.readline().decode('utf-8').strip()

			if "PING" in read:
				ser.write("PONG\n".encode('utf-8'))
				return True
		time.sleep(1)
	return False

def connect_to_pcb(ports):
	ports = list_available_ports()

	# Try to connect to each port
	ser = None
	for port in ports:
		print(f"testing port {port}")
		try:
			ser = conect_to_port(port)
		except Exception as e:
			print(f"Exception: {e}")
			continue

		ok = wait_for_connection(ser)
		break

	if ser is None:
		raise Exception("Could not connect to any serial port")

	return ser, ok

if __name__ == "__main__":
	ports = list_available_ports()
	pcb_ser, ok = connect_to_pcb(ports)

	if not ok:
		print("LOG : Could not connect to PCB")
		exit(-1)

	print("Connected to PCB")
	print(pcb_ser.name)