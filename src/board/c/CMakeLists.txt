cmake_minimum_required(VERSION 3.12)

include($ENV{PICO_SDK_PATH}/external/pico_sdk_import.cmake)


project(volumer CXX C ASM)

set(CMAKE_C_STANDARD 11)
set(CMAKE_CXX_STANDARD 17)
set(PICO_SDK_PATH "/home/jonas/installs/pico/pico-sdk")


add_compile_options(-Wall)

# Initialize the RP2040 SDK
include(${PICO_SDK_PATH}/pico_sdk_init.cmake)

# Set the CMake project type to RP2040
pico_sdk_init()

add_executable(${PROJECT_NAME}
    main.c
)

target_link_libraries(${PROJECT_NAME}
    pico_stdlib
	hardware_adc
)

# Create a UF2 file for flashing the RP2040
pico_add_extra_outputs(${PROJECT_NAME})

# Set the output directory for the built binary (if desired)
set(EXECUTABLE_OUTPUT_PATH ${CMAKE_BINARY_DIR}/bin)

# Enable USB output, disable UART output
pico_enable_stdio_usb(${PROJECT_NAME} 1)
pico_enable_stdio_uart(${PROJECT_NAME} 0)
