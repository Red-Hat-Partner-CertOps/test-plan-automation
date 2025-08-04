# Hardware Classification System

A Go-based system that automatically classifies hardware components and generates structured JSON output with detailed feature mappings. The system processes device names and maps them to standardized hardware categories with appropriate feature IDs and specifications.

## Features

- **Automatic Hardware Classification**: Classifies devices into categories (Processor, Ethernet, Storage, DVD)
- **Intelligent Feature Mapping**: Maps hardware specifications to standardized feature IDs
- **JSON Output Generation**: Creates structured JSON with detailed component information
- **Extensible Architecture**: Easy to add new device types and feature mappings
- **Pattern Recognition**: Uses keyword matching and pattern detection for classification

## Supported Hardware Categories

### Processors
- Multi-core CPUs (Intel, AMD)
- Server processors (Xeon, EPYC)
- Desktop processors (Core i-series, Ryzen)

### Network Cards
- Ethernet adapters (1Gb, 10Gb, 25Gb, 40Gb, 50Gb, 100Gb, 200Gb, 400Gb)
- ConnectX-6/7 cards with RoCE and InfiniBand support
- Multi-port network interfaces

### Storage Devices
- SATA SSDs and HDDs
- NVMe M.2 and U.2 drives
- RAID controllers
- Fibre Channel adapters (16Gb, 32Gb)

### Optical Drives
- DVD-RW drives

## 🛠 Installation & Usage

### Prerequisites
- Go 1.21 or higher

### Setup
1. Clone or download the project
2. Ensure you have Go installed
3. Navigate to the project directory

### Running the System
```bash
go run .
```

### Input Format
Create an `input.txt` file with device names, one per line:

```
AMD EPYC 9965 192-Core Processor  384 threads
RAID-LSI-9560-LP-16i-8GB
1gb UN-NIC-GE-4P-360T-B2
10gb NIC-ETH531F-LP-2P
25gb UN-NIC-620F-B2-25Gb-2P-1
100gb  UN-IB-MCX653105A-ECAT-100G-1P
SATA M.2 UN-SSD-480G-SATA-S4520-M.2
U.2 NVMe UN-SSD-960G-NVMe-PM9A3
NVMe M.2 UN-SSD-480G-NVMe-7450PRO-M.2
SATA_SSD: UN-SSD-1.92T-SATA-S4620-UCS
SAS_SSD: Seagate ST1800MM0129 SAS HDD 1.8T
```

### Output
The system generates `output.json` with structured component data:

```json
[
  {
    "name": "AMD EPYC 9965 192-Core Processor  384 threads",
    "optionType": "Integrated",
    "features": [
      {
        "name": "CORE",
        "features": [
          {
            "id": 90,
            "name": "Max Logical CPU",
            "isPublishable": true
          }
        ]
      },
      {
        "name": "FV_CORE",
        "features": [
          {
            "id": 326,
            "name": "Virtual Machine (Host)",
            "isPublishable": true
          }
        ]
      }
    ]
  }
]
```

## 🔧 Configuration

### Adding New Device Types
To add support for new hardware categories:

1. **Update Classification Logic** in `classifyDevice()` function:
```go
func classifyDevice(deviceName string) string {
    lower := strings.ToLower(deviceName)
    if strings.Contains(lower, "gpu") || strings.Contains(lower, "graphics") {
        return "Graphics"
    }
    // ... existing logic
}
```

2. **Add Feature Mapping** in the main switch statement:
```go
case "Graphics":
    component = Component{
        Name:           deviceName,
        OptionType:     "Optional",
        Type:           "Graphics",
        IsSupplemental: false,
        Features: []Feature{
            {Name: "GRAPHICS", Features: []Feature{
                {ID: 9001, Name: "GPU Acceleration", IsPublishable: true},
            }},
        },
    }
```

## Project Structure

```
test-plan-automation/
├── main.go          # Main classification logic
├── input.txt        # Input device names
├── output.json      # Generated JSON output
├── go.mod          # Go module definition
└── README.md       # This file
```

## Core Components

### Data Structures
- **Feature**: Represents individual hardware capabilities
- **Component**: Complete hardware component with features and metadata

### Classification Process
1. **Input Processing**: Reads device names from `input.txt`
2. **Device Classification**: Categorizes devices using keyword matching
3. **Feature Mapping**: Maps device capabilities to feature IDs
4. **JSON Generation**: Creates structured output with all component data

### Special Handling
- **ConnectX Cards**: Advanced network cards with RoCE and InfiniBand features
- **Storage Technologies**: Differentiation between SATA, SAS, NVMe interfaces
- **Network Speeds**: Automatic detection of bandwidth capabilities

## How It Works

### Classification Algorithm
The system uses a rule-based approach with keyword matching:

```go
func classifyDevice(deviceName string) string {
    lower := strings.ToLower(deviceName)
    if strings.Contains(lower, "processor") {
        return "Processor"
    } else if strings.Contains(lower, "nic") || strings.Contains(lower, "gb") {
        return "Ethernet"
    }
    // ... additional logic
}
```

### Feature Detection
- **Speed Extraction**: Detects network speeds (1gb, 10gb, 25gb, etc.)
- **Technology Recognition**: Identifies storage technologies (SATA, NVMe, SAS)
- **Port Counting**: Extracts port information for network cards

### Extensibility
The system is designed for easy extension:
- Add new device patterns in `classifyDevice()`
- Define new feature mappings in the main switch statement
- Extend data structures for additional metadata

## Sample Output

The system processes hardware lists and generates comprehensive JSON with:
- Device categorization
- Feature ID mappings
- Metadata (option types, supplemental flags)
- Hierarchical feature organization

Example for a complete server configuration:
- 1x AMD EPYC processor with virtualization features
- Multiple network cards (1Gb, 10Gb, 25Gb, 100Gb)
- Various storage devices (SATA SSD, NVMe, RAID controller)

## Contributing

To contribute to this project:
1. Fork the repository
2. Add new classification rules or feature mappings
3. Test with sample hardware configurations
4. Submit pull requests with improvements

## License

This project is available for use and modification as needed for hardware classification tasks.

---

**Note**: This system provides a foundation for hardware classification that can be extended with additional AI/ML capabilities for improved accuracy and automated learning from classification results. 