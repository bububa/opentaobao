package omniorder

// SdtQueryPackageResponse 
type SdtQueryPackageResponse struct {
    // 站点信息
    Stations   []SdtStationDto `json:"stations,omitempty" xml:"stations>sdt_station_dto,omitempty"`
}