cmake_minimum_required(VERSION 3.20)
project(dhcp_tftp_server NONE)

# ---- configuration ----
set(GO_EXECUTABLE go)
set(BINARY_NAME server)

# ---- source files (for tracking only) ----
set(GO_CLI_SOURCES
    cmd/cli/main.go
)

set(CMAKE_CLI_SOURCE_DIR ${PROJECT_SOURCE_DIR}/cmd/cli)
set(CMAKE_BUILD_DIR ${PROJECT_SOURCE_DIR}/build)

# ---- build command ----
add_custom_target(build ALL
    # For Windows
    COMMAND ${GO_EXECUTABLE} build -o ${CMAKE_BUILD_DIR}/${BINARY_NAME}.exe

    # For Linux
    COMMAND ${GO_EXECUTABLE} build -o ${CMAKE_BUILD_DIR}/${BINARY_NAME}

    WORKING_DIRECTORY ${CMAKE_CLI_SOURCE_DIR}
    SOURCES ${GO_CLI_SOURCES}
)

# ---- clean helper ----
add_custom_target(clean-go
    # For Linux
    COMMAND ${CMAKE_COMMAND} -E remove ${CMAKE_BUILD_DIR}/${BINARY_NAME}

    # For Windows
    COMMAND ${CMAKE_COMMAND} -E remove ${CMAKE_BUILD_DIR}/${BINARY_NAME}.exe
)
