// https://docs.scientiamobile.com/documentation/infuze/infuze-c-api-user-guide

#include <stdlib.h>
#include <wurfl/wurfl.h>

// @NOTE: wurfl_device_capability_* methods are deprecated but there is no examples of using wurfl_enum_* method
// evean their own C++ library don't use it
static char** wurfl_device_capability_enumerate(wurfl_device_capability_enumerator_handle enumerator, char** result) {
  int offset = 0;
  const int maxElements = 1000 * 2;

	while (wurfl_device_capability_enumerator_is_valid(enumerator) != 0) {
		const char* name = wurfl_device_capability_enumerator_get_name(enumerator);
		const char* val = wurfl_device_capability_enumerator_get_value(enumerator);

		if (name != NULL && val != NULL) {
      if (result == NULL) {
        result = (char**)malloc(sizeof(char*) * maxElements + 1);
      }

      if (offset < maxElements) {
        result[offset++] = (char *)name;
        result[offset++] = (char *)val;
      } else {
        break;
      }
		}

		wurfl_device_capability_enumerator_move_next(enumerator);
	}

  if (result != NULL) {
    result[offset] = NULL;
  }

  return result;
}

static char** wurfl_device_get_capabilities(wurfl_device_handle hwurfldevice, char** result) {
	wurfl_device_capability_enumerator_handle enumerator = wurfl_device_get_capability_enumerator(hwurfldevice);
	result = wurfl_device_capability_enumerate(enumerator, result);
	wurfl_device_capability_enumerator_destroy(enumerator);
  return result;
}

static char** wurfl_device_get_virtual_capabilities(wurfl_device_handle hwurfldevice, char** result) {
	wurfl_device_capability_enumerator_handle enumerator = wurfl_device_get_virtual_capability_enumerator(hwurfldevice);
	result = wurfl_device_capability_enumerate(enumerator, result);
	wurfl_device_capability_enumerator_destroy(enumerator);
  return result;
}
