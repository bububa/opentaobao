package wdk

// LogisticsNodeFullInfo 结构体
type LogisticsNodeFullInfo struct {
	// 节点名称
	NodeName string `json:"node_name,omitempty" xml:"node_name,omitempty"`
	// 节点商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 浙江省
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 西集镇任郎路与供杜路交叉口西北200米
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 节点编码
	NodeCode string `json:"node_code,omitempty" xml:"node_code,omitempty"`
	// 110100
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 经纬度
	Poi string `json:"poi,omitempty" xml:"poi,omitempty"`
	// 区域ID
	AreaId string `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区域名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 省份id
	ProvId string `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// *      * 仓           WAREHOUSE(1, &#34;仓&#34;),      *      * 揽运站           COLLECT_DOCK(2, &#34;揽运站&#34;),      *      * 配送站           DELIVERY_DOCK(3, &#34;配送站&#34;),      *      * 近端履约中心           CFC(4, &#34;近端履约中心&#34;),     ;
	NodeType int64 `json:"node_type,omitempty" xml:"node_type,omitempty"`
}
