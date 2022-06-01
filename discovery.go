package main

import (
	"fmt"
	"net"
)

const HDHOMERUN_DEVICE_TYPE_WILDCARD = 0xFFFFFFFF
const HDHOMERUN_DEVICE_TYPE_TUNER = 0x00000001
const HDHOMERUN_DEVICE_TYPE_STORAGE = 0x00000005
const HDHOMERUN_DEVICE_ID_WILDCARD = 0xFFFFFFFF

// Main function
func main() {
	discoverPrint(nil)
}

func discoverPrint(targetIpStr *string) (int, error) {

	// Decode possible IP address, leave as nil if non-existant
	var ip net.IP = nil
	if targetIpStr != nil {
		ip = net.ParseIP(*targetIpStr)
		if ip == nil {
			return -1, fmt.Errorf("invalid ip address: %s", *targetIpStr)
		}
	}

	/*
	   static int discover_print(char *target_ip_str)
	   {
	           uint32_t target_ip = 0;
	           if (target_ip_str) {
	                   target_ip = parse_ip_addr(target_ip_str);
	                   if (target_ip == 0) {
	                           fprintf(stderr, "invalid ip address: %s\n", target_ip_str);
	                           return -1;
	                   }
	           }

	           struct hdhomerun_discover_device_t result_list[64];
	           int result_count = hdhomerun_discover_find_devices_custom_v2(target_ip, HDHOMERUN_DEVICE_TYPE_WILDCARD, HDHOMERUN_DEVICE_ID_WILDCARD, result_list, 64);
	           if (result_count < 0) {
	                   fprintf(stderr, "error sending discover request\n");
	                   return -1;
	           }

	           struct hdhomerun_discover_device_t *result = result_list;
	           struct hdhomerun_discover_device_t *result_end = result_list + result_count;

	           int valid_count = 0;
	           while (result < result_end) {
	                   if (result->device_id == 0) {
	                           result++;
	                           continue;
	                   }

	                   printf("hdhomerun device %08X found at %u.%u.%u.%u\n",
	                           (unsigned int)result->device_id,
	                           (unsigned int)(result->ip_addr >> 24) & 0x0FF, (unsigned int)(result->ip_addr >> 16) & 0x0FF,
	                           (unsigned int)(result->ip_addr >> 8) & 0x0FF, (unsigned int)(result->ip_addr >> 0) & 0x0FF
	                   );

	                   valid_count++;
	                   result++;
	           }

	           if (valid_count == 0) {
	                   printf("no devices found\n");
	           }

	           return valid_count;
	   }
	*/
	return 0, nil
}

type hdhomerunDiscoverDevice struct {
	ipAddr IP
}

/*
struct hdhomerun_discover_device_t {
        uint32_t ip_addr;
        uint32_t device_type;
        uint32_t device_id;
        uint8_t tuner_count;
        bool is_legacy;
        char device_auth[25];
        char base_url[29];
};

struct hdhomerun_discover_device_v3_t {
        uint32_t ip_addr;
        uint32_t device_type;
        uint32_t device_id;
        uint8_t tuner_count;
        bool is_legacy;
        char device_auth[25];
        char base_url[29];

        char storage_id[37];
        char lineup_url[128];
        char storage_url[128];
};
*/
