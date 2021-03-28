cmake_minimum_required (VERSION 3.9)
project (gmni VERSION 0.0.0 LANGUAGES C)

include (GNUInstallDirs)
find_package (OpenSSL REQUIRED)

set (libgmni_src
	src/client.c
	src/escape.c
	src/tofu.c
	src/url.c
	src/util.c
	src/parser.c
	)

add_library (libgmni_obj OBJECT ${libgmni_src})
set_target_properties (libgmni_obj PROPERTIES POSITION_INDEPENDENT_CODE ON)
target_include_directories (libgmni_obj PRIVATE include)

add_library (libgmni_shared SHARED $<TARGET_OBJECTS:libgmni_obj>)
set_target_properties (libgmni_shared PROPERTIES OUTPUT_NAME "gmni")
target_include_directories (libgmni_shared
	PUBLIC
	$<INSTALL_INTERFACE:include>
	$<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
	)

add_library (libgmni_static STATIC $<TARGET_OBJECTS:libgmni_obj>)
set_target_properties (libgmni_static PROPERTIES OUTPUT_NAME "gmni")
target_include_directories (libgmni_static
	PUBLIC
	$<INSTALL_INTERFACE:include>
	$<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
	)

configure_file (libgmni.pc.in libgmni.pc @ONLY)
find_program (SCDOC scdoc)

add_executable (gmni src/gmni.c)
target_link_libraries (gmni PUBLIC libgmni_shared OpenSSL::SSL)
add_custom_command (
	TARGET gmni
	COMMAND ${SCDOC} < ${CMAKE_SOURCE_DIR}/doc/gmni.scd | gzip > ${CMAKE_BINARY_DIR}/gmni.1.gz
	)

add_executable (gmnlm src/gmnlm.c)
target_link_libraries (gmnlm PUBLIC libgmni_shared OpenSSL::SSL)
add_custom_command (
	TARGET gmnlm
	COMMAND ${SCDOC} < ${CMAKE_SOURCE_DIR}/doc/gmnlm.scd | gzip > ${CMAKE_BINARY_DIR}/gmnlm.1.gz
	)


install (
	TARGETS libgmni_shared libgmni_static gmni gmnlm
	LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
	ARCHIVE DESTINATION ${CMAKE_INSTALL_LIBDIR}
	RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
	)
install (
	DIRECTORY include/gmni/
	DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/gmni
	)
install (
	FILES ${CMAKE_CURRENT_BINARY_DIR}/libgmni.pc
	DESTINATION ${CMAKE_INSTALL_PREFIX}/libdata/pkgconfig
	)
install (
	FILES ${CMAKE_CURRENT_BINARY_DIR}/gmni.1.gz ${CMAKE_CURRENT_BINARY_DIR}/gmnlm.1.gz
	DESTINATION ${CMAKE_INSTALL_MANDIR}/man1
	)

