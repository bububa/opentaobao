package usergrowth

// SuggestionContextParam 结构体
type SuggestionContextParam struct {
	// 不同媒体资源位
	SiteId string `json:"site_id,omitempty" xml:"site_id,omitempty"`
	// 天气代码，可空。1：晴, 2：晴（夜）, 3：多云, 4：多云（夜）, 5：阴, 6：阴（夜）, 7：雾, 8：雾（夜）, 9：霾, 10：霾（夜）, 11：雨, 12：雨（夜）, 13：雪, 14：雪（夜）
	Weather string `json:"weather,omitempty" xml:"weather,omitempty"`
	// 地域
	Region string `json:"region,omitempty" xml:"region,omitempty"`
	// 性别：m/f
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 附加参数，可用于传递一些非业务参数，如是否压测流量(&#34;isStressTest&#34;:true)等。格式为：key1:value1,key2:value2
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 设备信息
	DeviceId *DeviceIdParam `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 年龄
	Age int64 `json:"age,omitempty" xml:"age,omitempty"`
	// 温度，包括最高和最低
	Temperature *Temperature `json:"temperature,omitempty" xml:"temperature,omitempty"`
}
