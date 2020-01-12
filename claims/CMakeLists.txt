cmake_minimum_required(VERSION 3.5)
project(claims VERSION 1.0.0)
set(CMAKE_CXX_STANDARD 14)
set(EOSIO_WASM_OLD_BEHAVIOR "Off")

if(EOSIO_CDT_ROOT STREQUAL "" OR NOT EOSIO_CDT_ROOT)
    MESSAGE("EOSIO_CDT_ROOT variable not found")
    set(EOSIO_CDT_ROOT "/home/tjoma/dev/eos/eosio.cdt")
endif()

include_directories(${EOSIO_CDT_ROOT}/libraries/eosiolib/contracts)

set(SRC_LIST src/claims.cpp)
add_executable(claims ${SRC_LIST})