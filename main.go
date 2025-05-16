/*package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Feature struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name"`
	IsPublishable bool      `json:"isPublishable,omitempty"`
	Features      []Feature `json:"features,omitempty"`
}

type Component struct {
	Name           string    `json:"name"`
	OptionType     string    `json:"optionType"`
	IsSupplemental bool      `json:"isSupplemental"`
	Features       []Feature `json:"features"`
}

func main() {
	input := []string{
		"Ethernet Controller: Intel I350T2 1G",
		"Ethernet Controller: ZTE NS212 10G",
		"Ethernet Controller: ZTE NS312 20G",

		"Video Card: ASPEED AST2600",

		"Storage Controller: Seagate ST4000NM024B-2TF103 *1 RAID RS241",
		"Storage Controller :Samsung: SAMSUNG MZ7L3960HCJR-00B7C *1 RAID RS241",
		"Storage Controller: Micron	Micron_7450_MTFDKCC3T2TFS *1",
		"Storage Controller: Intel INTEL SSDSC2KG960G8",

		"USB: 4 x USB 3.0",
	}

	jsonData := []Component{}

	for _, line := range input {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) < 2 {
			continue
		}
		category, value := parts[0], parts[1]
		if category == "Video Card" {
			videoComponent := Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{
						Name: "VIDEO",
						Features: []Feature{
							{ID: 406, Name: "Graphic Console", IsPublishable: true},
						},
					},
					{
						Name: "VIDEO_DRM",
						Features: []Feature{
							{ID: 2836, Name: "Basic GPU Graphics", IsPublishable: true},
						},
					},
				},
			}
			jsonData = append(jsonData, videoComponent)
		} else if category == "Ethernet Controller" {
			ethernetComponent := Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{
						Name: "1GIGETHERNET",
						Features: []Feature{
							{ID: 606, Name: "1 Gigabit Ethernet", IsPublishable: true},
						},
						{
							Name: "10GIGETHERNET",
							Features: []Features{
								{ ID: 636,Name: "10 Gigabit Ethernet", IsPublishable: true},
						},
						{
								Name: "20GigEthernet",
								Features: []Features{
									{ ID: 646,Name: "20 Gigabit Ethernet",isPublishable: true},
						},
			},
		},
	},
},
			jsonData = append(jsonData, ethernetComponent)
		} else if category == "USB" {
			usbComponent := Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{
						Name: "USB3_5Gbps",
						Features: []Feature{
							{ID: 376, Name: "USB C (5 Gigabit) Ports", IsPublishable: true},
							{ID: 2476, Name: "USB 3 (5 Gigabit) Ports", IsPublishable: true},
						},
					},
				},
			}
			jsonData = append(jsonData, usbComponent)
		} else if category == "Storage Controller" {
			storageComponent := Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{
						Name: "STORAGE",
						Features: []Feature{
							{ID: 986, Name: "M.2 NVMe", IsPublishable: true},
					},
					{
						Name: "STORAGE",
						Features: []Feature{
							{ID: 1056, Name: "U.2 NVMe", IsPublishable: true},
					},
					{   Name: "STORAGE",
						Features: []Feature{
							{ID: 1006, Name: "PCIE NVMe", IsPublishable: true},
					},
				},
			},
		},
	},
	jsonData = append(jsonData, storageComponent)
}
		result, _ := json.MarshalIndent(jsonData, "", "  ")
		file, err := os.Create("output.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		_, err = file.Write(result)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			}
		}
	}*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Feature struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name"`
	IsPublishable bool      `json:"isPublishable,omitempty"`
	Features      []Feature `json:"features,omitempty"`
}

type Component struct {
	Name           string    `json:"name"`
	OptionType     string    `json:"optionType"`
	IsSupplemental bool      `json:"isSupplemental"`
	Type           string    `json:"type,omitempty"`
	Features       []Feature `json:"features"`
}

func main() {
	input := []string{
		"Processor: ARL CPU U15 i7 non-Vpro,1.7GHZ,12C",
		"Processor: ARL CPU H45 i9 Vpro,2.7GHZ,16C",
		"Storage: Integrated PCH in CPU",
		"Storage: PCIe_NVMe SSD",
		"TB :2x TBT4 Type C",
		"USB: 1x USB Type C MF 10 Gbps",
		"USB: 1x USB Type-A 3.2 G1 5Gbps",

		"Ethernet: Intel I219LM 1Gig",
		"Audio: HP Audio, dual stereo speakers, dual array digital microphones, functions keys for volume up and down, combo microphone/headphone jack",
		"HDMI: Audio HDMI",
		"Storage: SD Card reader",

		"Fingerprint: Fingerprint Reader",

		"Wirless: Intel AX211 Wi-Fi 6E +BT 5.3 M.2 and Intel BE201 Wi-Fi 7 +BT 5.4 M.2",

		"Video: Discrete Graphic Cards: RTX A500 Ada Generation",
		"Display: eDP: 300 TOP, 400nits and 500nits",
	}

	jsonData := []Component{}

	for _, line := range input {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) < 2 {
			continue
		}
		category, value := parts[0], parts[1]

		switch category {
		case "Processor":
			jsonData = append(jsonData, Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{Name: "CORE", Features: []Feature{{ID: 90, Name: "Max Logical CPU", IsPublishable: true}}},
					{Name: "FV_CORE", Features: []Feature{{ID: 326, Name: "Virtual Machine (Host)", IsPublishable: true}}},
					{Name: "FV_MEMORY", Features: []Feature{{ID: 326, Name: "Virtual Machine (Host)", IsPublishable: true}}},
					{Name: "MEMORY", Features: []Feature{{ID: 80, Name: "System Memory", IsPublishable: true}}},
					{Name: "CPUSCALING", Features: []Feature{{ID: 116, Name: "System Controlled Scaling", IsPublishable: true}}},
					{Name: "profiler_hardware_core", Features: []Feature{{ID: 2867, Name: "CPU Core Performance  Counters", IsPublishable: true}}},
					{Name: "profiler_hardware_uncore", Features: []Feature{{ID: 2877, Name: "Uncore Performance Counters", IsPublishable: true}}},

					{Name: "FV_CPU_PINNING", Features: []Feature{{ID: 2406, Name: "CPU Pinning", IsPublishable: true}}},

					{Name: "FV_PCIE_STORAGE_PASSTHROUGH", Features: []Feature{{ID: 2576, Name: "PCIE Pass-Through Storage", IsPublishable: false}}},
					{Name: "FV_PCIE_NETWORK_PASSTHROUGH", Features: []Feature{{ID: 2586, Name: "PCIE Pass-Through Network", IsPublishable: false}}},

					{Name: "FV_USB_STORAGE_PASSTHROUGH", Features: []Feature{{ID: 2596, Name: "USB Pass-Through Storage", IsPublishable: false}}},
					{Name: "FV_USB_NETWORK_PASSTHROUGH", Features: []Feature{{ID: 2606, Name: "USB Pass-Through Network", IsPublishable: false}}},

					{Name: "FV_LIVE_MIGRATION", Features: []Feature{{ID: 2616, Name: "FV Live Migration", IsPublishable: false}}},
				},
			})
		case "Video":
			jsonData = append(jsonData, Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{Name: "VIDEO", Features: []Feature{{ID: 406, Name: "Graphic Console", IsPublishable: true}}},
					{Name: "VIDEO_DRM", Features: []Feature{{ID: 2836, Name: "Basic GPU Graphics", IsPublishable: true}}},
				},
			})

		case "Ethernet":
			ethernetSpeeds := map[string]Feature{
				"1GbE":   {ID: 606, Name: "1 Gigabit Ethernet", IsPublishable: true},
				"2.5GbE": {ID: 616, Name: "2.5 Gigabit Ethernet", IsPublishable: true},
				"10GbE":  {ID: 636, Name: "10 Gigabit Ethernet", IsPublishable: true},
				"25GbE":  {ID: 656, Name: "25 Gigabit Ethernet", IsPublishable: true},
				"40GbE":  {ID: 666, Name: "40 Gigabit Ethernet", IsPublishable: true},
				"50GbE":  {ID: 676, Name: "50 Gigabit Ethernet", IsPublishable: true},
				"100GbE": {ID: 686, Name: "100 Gigabit Ethernet", IsPublishable: true},
				"200GbE": {ID: 2306, Name: "200 Gigabit Ethernet", IsPublishable: true},
				"400GbE": {ID: 3057, Name: "400 Gigabit Ethernet", IsPublishable: true},
			}

			var selectedFeature Feature
			var speedName string

			for speed, feature := range ethernetSpeeds {
				if strings.Contains(value, speed) {
					selectedFeature = feature
					speedName = strings.Replace(speed, "GbE", "GigEthernet", 1) // Converts "1GbE" -> "1GigEthernet"
					break
				}
			}

			if selectedFeature.Name != "" {
				jsonData = append(jsonData, Component{
					Name:           value,
					OptionType:     "Optional",
					Type:           "Network",
					IsSupplemental: false,
					Features: []Feature{
						{Name: speedName, Features: []Feature{selectedFeature}},
					},
				})
			}

		case "USB":
			jsonData = append(jsonData, Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{Name: "USB3_5Gbps", Features: []Feature{
						{ID: 376, Name: "USB C (5 Gigabit) Ports", IsPublishable: true},
						{ID: 2476, Name: "USB 3 (5 Gigabit) Ports", IsPublishable: true},
					}},
				},
			})

		case "Storage":
			jsonData = append(jsonData, Component{
				Name:           value,
				OptionType:     "Integrated",
				IsSupplemental: false,
				Features: []Feature{
					{Name: "STORAGE", Features: []Feature{{ID: 986, Name: "M.2 NVMe", IsPublishable: true}}},
					{Name: "STORAGE", Features: []Feature{{ID: 1056, Name: "U.2 NVMe", IsPublishable: true}}},
					{Name: "STORAGE", Features: []Feature{{ID: 1006, Name: "PCIE NVMe", IsPublishable: true}}},
				},
			})
		}
	}

	result, _ := json.MarshalIndent(jsonData, "", "  ")
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
