package shop

// OpenApiSearchRequest 结构体
type OpenApiSearchRequest struct {
	// 渠道版本 alipay,koubei,eleme,gaode,taobao
	App string `json:"app,omitempty" xml:"app,omitempty"`
	// 所选城市
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 客户端
	ClientOs string `json:"client_os,omitempty" xml:"client_os,omitempty"`
	// 客户端版本
	ClientVersion string `json:"client_version,omitempty" xml:"client_version,omitempty"`
	// 客户端上下文
	Context string `json:"context,omitempty" xml:"context,omitempty"`
	// 所选城市
	CurrentCity string `json:"current_city,omitempty" xml:"current_city,omitempty"`
	// 所选区
	CurrentDistrict string `json:"current_district,omitempty" xml:"current_district,omitempty"`
	// 省
	CurrentProvince string `json:"current_province,omitempty" xml:"current_province,omitempty"`
	// 固定值
	Forward bool `json:"forward,omitempty" xml:"forward,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 所在商圈
	LbsBusiAreaId string `json:"lbs_busi_area_id,omitempty" xml:"lbs_busi_area_id,omitempty"`
	// 所在城市
	LbsCityId string `json:"lbs_city_id,omitempty" xml:"lbs_city_id,omitempty"`
	// 所在区
	LbsDistrictId string `json:"lbs_district_id,omitempty" xml:"lbs_district_id,omitempty"`
	// 经纬度精度
	LocationAccuracy string `json:"location_accuracy,omitempty" xml:"location_accuracy,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 个数
	OriginalSize int64 `json:"original_size,omitempty" xml:"original_size,omitempty"`
	// OS 版本
	OsVersion string `json:"os_version,omitempty" xml:"os_version,omitempty"`
	// 扩展参数
	ParamsMap string `json:"params_map,omitempty" xml:"params_map,omitempty"`
	// 埋点ID
	PlaceId string `json:"place_id,omitempty" xml:"place_id,omitempty"`
	// 关键词
	Query string `json:"query,omitempty" xml:"query,omitempty"`
	// 请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 场景吗
	SceneId string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 搜索ID
	SearchId string `json:"search_id,omitempty" xml:"search_id,omitempty"`
	// 菜单项
	SelectedMenus string `json:"selected_menus,omitempty" xml:"selected_menus,omitempty"`
	// sessionid
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 个数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 排序
	Sort string `json:"sort,omitempty" xml:"sort,omitempty"`
	// 来源埋点
	SrcSpm string `json:"src_spm,omitempty" xml:"src_spm,omitempty"`
	// 分页
	Start int64 `json:"start,omitempty" xml:"start,omitempty"`
	// 搜索TOKEN
	TokenId string `json:"token_id,omitempty" xml:"token_id,omitempty"`
	// traceID
	Trace string `json:"trace,omitempty" xml:"trace,omitempty"`
	// 用户ID
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
}
