package campus

// DeviceStandardApiDto 结构体
type DeviceStandardApiDto struct {
	// 设备参数信息
	MetaPointDatas []MeatDataApiDto `json:"meta_point_datas,omitempty" xml:"meta_point_datas>meat_data_api_dto,omitempty"`
	// 设备标签列表
	TagInfoList []TagInfoApiDto `json:"tag_info_list,omitempty" xml:"tag_info_list>tag_info_api_dto,omitempty"`
	// 设备code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// 楼宇名称
	BuildingName string `json:"building_name,omitempty" xml:"building_name,omitempty"`
	// 楼层名称
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// 模板code
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 最后更新时间
	LastUpdateTime string `json:"last_update_time,omitempty" xml:"last_update_time,omitempty"`
	// 设备uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// 版本信息
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 空间单元名称
	SpaceName string `json:"space_name,omitempty" xml:"space_name,omitempty"`
	// 模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 设备别名
	Nickname string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 设备运行状态展示字段
	RunStatusName string `json:"run_status_name,omitempty" xml:"run_status_name,omitempty"`
	// campusId
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 楼宇id
	BuildingId int64 `json:"building_id,omitempty" xml:"building_id,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 模板id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 空间单元id
	SpaceId int64 `json:"space_id,omitempty" xml:"space_id,omitempty"`
	// 设备运行状态。0->在线 1->离线 2->故障
	RunStatus int64 `json:"run_status,omitempty" xml:"run_status,omitempty"`
	// 公司id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 设备是否运行
	BeRun bool `json:"be_run,omitempty" xml:"be_run,omitempty"`
	// 是否是逻辑设备
	BeLogic bool `json:"be_logic,omitempty" xml:"be_logic,omitempty"`
}
