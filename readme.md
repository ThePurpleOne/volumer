# Volumer
Volume Mixer Hardware module for easy access to individual application-wise volume control. 


## PCB

### V0.1
[V0.1](./pcb/volumer_v0.1/)

Version 0.1 was a prototype 
- Rotary encoders are not really satisfying to turn
- you can only turn one at the time
- Wrong rotary choice (really hard to turn and not clicky)
- Encoders too close to each other, cant pass a finger between them

### V0.2
[V0.2](./pcb/volumer_v0.2/)

Version 0.2 is the first version that is actually usable:
- Linear pots are nice and satisfying
- Size of pots is good
- Place between pots is good
- LEDs are useless
- PCB could be a lot smaller (pots mounted on other side of PCB)


## Board Firmware 
The board firmware is written in circuitpython for now.


Available in [firmware](./src/board/cpy/)


## RP2040 C SDK
```bash
yay -S cmake arm-none-eabi-gcc arm-none-eabi-newlib arm-none-eabi-gdb arm-none-eabi-binutils
```

```bash
cmake -DCMAKE_TOOLCHAIN_FILE=path/to/your/toolchain-file.cmake ..
```