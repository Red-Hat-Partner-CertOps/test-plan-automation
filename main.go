package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Feature struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	IsPublishable bool      `json:"isPublishable,omitempty"`
	Features      []Feature `json:"features,omitempty"`
}

type Component struct {
	Name           string    `json:"name,omitempty"`
	OptionType     string    `json:"optionType,omitempty"`
	IsSupplemental bool      `json:"isSupplemental,omitempty"`
	Type           string    `json:"type,omitempty"`
	Features       []Feature `json:"features,omitempty"`
}

func classifyDevice(deviceName string) string {
	lower := strings.ToLower(deviceName)
	if strings.Contains(lower, "processor") {
		return "Processor"
	} else if strings.Contains(lower, "ethernet adapter") ||
		strings.Contains(lower, "network adapter") ||
		strings.Contains(lower, "connectx-6") ||
		strings.Contains(lower, "connectx-7") ||
		strings.Contains(lower, "gbe") ||
		strings.Contains(lower, "nic") || // <-- add this
		strings.Contains(lower, "gb") || // <-- add this
		strings.Contains(lower, "10gbase-t") {
		return "Ethernet"
	} else if strings.Contains(lower, "storage") || strings.Contains(lower, "raid") ||
		strings.Contains(lower, "sas/sata") || strings.Contains(lower, "fibre channel adapter") ||
		strings.Contains(lower, "pcie") || strings.Contains(lower, "hba") ||
		strings.Contains(lower, "32gb") || strings.Contains(lower, "16gb") {
		return "Storage"
	} else if strings.Contains(lower, "dvd") || strings.Contains(lower, "optical disk drive") {
		return "dvd"
	}
	return "Unknown"
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	deviceNames := strings.Split(string(data), "\n")
	var components []Component

	for _, deviceName := range deviceNames {
		deviceName = strings.TrimSpace(deviceName)
		if deviceName == "" {
			continue
		}

		deviceType := classifyDevice(deviceName)
		var component Component

		switch deviceType {
		case "Processor":
			component = Component{
				Name:           deviceName,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{Name: "CORE", Features: []Feature{{ID: 90, Name: "Max Logical CPU", IsPublishable: true}}},
					{Name: "FV_CORE", Features: []Feature{{ID: 326, Name: "Virtual Machine (Host)", IsPublishable: true}}},
					{Name: "FV_MEMORY", Features: []Feature{{ID: 326, Name: "Virtual Machine (Host)", IsPublishable: true}}},
					{Name: "MEMORY", Features: []Feature{{ID: 80, Name: "System Memory", IsPublishable: true}}},
					{Name: "CPUSCALING", Features: []Feature{{ID: 116, Name: "System Controlled Scaling", IsPublishable: true}}},
					{Name: "profiler_hardware_core", Features: []Feature{{ID: 2867, Name: "CPU Core Performance Counters", IsPublishable: true}}},
					{Name: "profiler_hardware_uncore", Features: []Feature{{ID: 2877, Name: "Uncore Performance Counters", IsPublishable: true}}},
					{Name: "FV_CPU_PINNING", Features: []Feature{{ID: 2406, Name: "CPU Pinning", IsPublishable: true}}},
					{Name: "FV_PCIE_STORAGE_PASSTHROUGH", Features: []Feature{{ID: 2576, Name: "PCIE Pass-Through Storage", IsPublishable: false}}},
					{Name: "FV_PCIE_NETWORK_PASSTHROUGH", Features: []Feature{{ID: 2586, Name: "PCIE Pass-Through Network", IsPublishable: false}}},
					{Name: "FV_USB_STORAGE_PASSTHROUGH", Features: []Feature{{ID: 2596, Name: "USB Pass-Through Storage", IsPublishable: false}}},
					{Name: "FV_USB_NETWORK_PASSTHROUGH", Features: []Feature{{ID: 2606, Name: "USB Pass-Through Network", IsPublishable: false}}},
					{Name: "FV_LIVE_MIGRATION", Features: []Feature{{ID: 2616, Name: "FV Live Migration", IsPublishable: false}}},
				},
			}
		case "Ethernet":
			speeds := []struct {
				Keywords  []string
				Feature   Feature
				SpeedName string
			}{
				{[]string{"1gbe", "1gb", "1 gigabit", "1gbe", "1gbit", "1gbe"}, Feature{ID: 606, Name: "1 Gigabit Ethernet", IsPublishable: true}, "1GigEthernet"},
				{[]string{"2.5gbe", "2.5gb", "2.5 gigabit", "2.5gbe", "2.5gbit"}, Feature{ID: 616, Name: "2.5 Gigabit Ethernet", IsPublishable: true}, "2.5GigEthernet"},
				{[]string{"10gbe", "10gb", "10gbase-t", "10 gigabit", "10gbe", "10gbit"}, Feature{ID: 636, Name: "10 Gigabit Ethernet", IsPublishable: true}, "10GigEthernet"},
				{[]string{"25gbe", "25gb", "25 gigabit", "25gbe", "25gbit"}, Feature{ID: 656, Name: "25 Gigabit Ethernet", IsPublishable: true}, "25GigEthernet"},
				{[]string{"40gbe", "40gb", "40 gigabit", "40gbe", "40gbit"}, Feature{ID: 666, Name: "40 Gigabit Ethernet", IsPublishable: true}, "40GigEthernet"},
				{[]string{"50gbe", "50gb", "50 gigabit", "50gbe", "50gbit"}, Feature{ID: 676, Name: "50 Gigabit Ethernet", IsPublishable: true}, "50GigEthernet"},
				{[]string{"100gbe", "100gb", "100 gigabit", "100gbe", "100gbit"}, Feature{ID: 686, Name: "100 Gigabit Ethernet", IsPublishable: true}, "100GigEthernet"},
				{[]string{"200gbe", "200gb", "200 gigabit", "200gbe", "200gbit"}, Feature{ID: 2306, Name: "200 Gigabit Ethernet", IsPublishable: true}, "200GigEthernet"},
				{[]string{"400gbe", "400gb", "400 gigabit", "400gbe", "400gbit"}, Feature{ID: 3057, Name: "400 Gigabit Ethernet", IsPublishable: true}, "400GigEthernet"},
			}

			lower := strings.ToLower(deviceName)
			var selectedFeature Feature
			var speedName string

			for _, s := range speeds {
				for _, kw := range s.Keywords {
					if strings.Contains(lower, kw) {
						selectedFeature = s.Feature
						speedName = s.SpeedName
						break
					}
				}
				if selectedFeature.Name != "" {
					break
				}
			}

			// Special handling for ConnectX-6/7 and 200GbE
			if (strings.Contains(lower, "connectx-6") || strings.Contains(lower, "connectx-7")) &&
				(strings.Contains(lower, "200gbe") || strings.Contains(lower, "200gb") || strings.Contains(lower, "200 gigabit")) {
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					Type:           "Network",
					IsSupplemental: false,
					Features: []Feature{
						{
							Name: "200GigEthernet",
							Features: []Feature{
								{ID: 2306, Name: "200 Gigabit Ethernet", IsPublishable: true},
							},
						},
						{
							Name: "200GigRoCE",
							Features: []Feature{
								{ID: 2336, Name: "200 Gigabit RoCE", IsPublishable: true},
							},
						},
						{
							Name: "Infiniband_HDR",
							Features: []Feature{
								{ID: 2316, Name: "HDR Infiniband", IsPublishable: true},
							},
						},
					},
				}
			} else if (strings.Contains(lower, "connectx-6") || strings.Contains(lower, "connectx-7")) &&
				(strings.Contains(lower, "100gbe") || strings.Contains(lower, "100gb") || strings.Contains(lower, "100 gigabit")) {
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					Type:           "Network",
					IsSupplemental: false,
					Features: []Feature{
						{
							Name: "100GigEthernet",
							Features: []Feature{
								{ID: 686, Name: "100 Gigabit Ethernet", IsPublishable: true},
							},
						},
						{
							Name: "100GigRoCE",
							Features: []Feature{
								{ID: 936, Name: "100 Gigabit RoCE", IsPublishable: true},
							},
						},
						{
							Name: "Infiniband_HDR",
							Features: []Feature{
								{ID: 2316, Name: "HDR Infiniband", IsPublishable: true},
							},
						},
					},
				}
			} else {
				if selectedFeature.Name == "" {
					// fallback generic Ethernet
					selectedFeature = Feature{ID: 606, Name: "Ethernet", IsPublishable: true}
					speedName = "Ethernet"
				}
				component = Component{
					Name:           deviceName,
					OptionType:     "Integrated",
					Type:           "Network",
					IsSupplemental: true,
					Features: []Feature{
						{Name: speedName, Features: []Feature{selectedFeature}},
					},
				}
			}

		case "Storage":
			lower := strings.ToLower(deviceName)
			switch {
			case strings.Contains(lower, "sas"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Integrated",
					IsSupplemental: false,
					Features:       []Feature{{Name: "STORAGE", Features: []Feature{{ID: 1036, Name: "SAS", IsPublishable: true}}}},
				}
			case strings.Contains(lower, "sata"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Integrated",
					IsSupplemental: false,
					Features:       []Feature{{Name: "STORAGE", Features: []Feature{{ID: 1016, Name: "SATA", IsPublishable: true}}}},
				}
			case strings.Contains(lower, "m.2") && strings.Contains(lower, "sata"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					IsSupplemental: false,
					Features:       []Feature{{Name: "M2_SATA", Features: []Feature{{ID: 996, Name: "M.2 SATA", IsPublishable: true}}}},
				}
			case strings.Contains(lower, "m.2") && strings.Contains(lower, "nvme"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					IsSupplemental: false,
					Features:       []Feature{{Name: "M2_NVMe", Features: []Feature{{ID: 1057, Name: "M.2 NVMe", IsPublishable: true}}}},
				}
			case strings.Contains(lower, "u.2") && strings.Contains(lower, "nvme"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					IsSupplemental: false,
					Features:       []Feature{{Name: "STORAGE", Features: []Feature{{ID: 1056, Name: "U.2 NVMe", IsPublishable: true}}}},
				}
			case strings.Contains(lower, "raid"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					IsSupplemental: false,
					Features: []Feature{{Name: "STORAGE", Features: []Feature{
						{ID: 1036, Name: "SAS", IsPublishable: true},
						{ID: 1076, Name: "Hardware Raid", IsPublishable: true},
						{ID: 2556, Name: "Flash Backed Cache", IsPublishable: true},
					}}},
				}
			case strings.Contains(lower, "32gb"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					IsSupplemental: false,
					Features:       []Feature{{Name: "STORAGE", Features: []Feature{{ID: 756, Name: "32 Gigabit Fibre Channel", IsPublishable: true}}}},
				}
			case strings.Contains(lower, "16gb"):
				component = Component{
					Name:           deviceName,
					OptionType:     "Optional",
					IsSupplemental: false,
					Features:       []Feature{{Name: "STORAGE", Features: []Feature{{ID: 746, Name: "16 Gigabit Fibre Channel", IsPublishable: true}}}},
				}
			}
		case "dvd":
			component = Component{
				Name:           deviceName,
				OptionType:     "Optional",
				IsSupplemental: false,
				Features:       []Feature{{Name: "DVD", Features: []Feature{{ID: 1216, Name: "DVD-RW", IsPublishable: true}}}},
			}
		}

		if component.Name != "" {
			components = append(components, component)
		}
	}

	result, err := json.MarshalIndent(components, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("output.json", result, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
