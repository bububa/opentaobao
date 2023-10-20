package drugtrace

// AlibabaalihealthdrugkytdrgetproteminfoModel 结构体
type AlibabaalihealthdrugkytdrgetproteminfoModel struct {
	// 存储温度
	StorageTemperatureList []StorageTemperatureList `json:"storage_temperature_list,omitempty" xml:"storage_temperature_list>storage_temperature_list,omitempty"`
	// 运输温度
	TransportTemperatureList []TransportTemperatureList `json:"transport_temperature_list,omitempty" xml:"transport_temperature_list>transport_temperature_list,omitempty"`
}
