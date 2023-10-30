# This has to be named code.py to run on the board at boot

import board
import analogio
from time import sleep

# Offset to subtract from the final value of ADC
# Depends on your pots
OFFSET = 4  

def main():
    chans = [board.A0, board.A1, board.A2, board.A3]
    adc = [analogio.AnalogIn(chan) for chan in chans]

    while True:
        vals = [chan.value for chan in adc]
        scaled_vals = [int((val / 65535) * 1023 - OFFSET) for val in vals]
        print(f"{scaled_vals[0]}|{scaled_vals[1]}|{scaled_vals[2]}|{scaled_vals[3]}|{scaled_vals[0]}")
        sleep(0.01)

main()
