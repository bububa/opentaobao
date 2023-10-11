package trade

// AdvertiseInfoDto 结构体
type AdvertiseInfoDto struct {
	// 信息流投放转换追踪标识
	ConversionTracking string `json:"conversion_tracking,omitempty" xml:"conversion_tracking,omitempty"`
	// 信息流投放广告账户id
	AdvertiserId string `json:"advertiser_id,omitempty" xml:"advertiser_id,omitempty"`
}
