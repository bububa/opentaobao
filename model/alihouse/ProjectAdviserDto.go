package alihouse

// ProjectAdviserDto 结构体
type ProjectAdviserDto struct {
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 公司ID
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 门店ID
	OuterShopId string `json:"outer_shop_id,omitempty" xml:"outer_shop_id,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 公司简称
	CompanyShortName string `json:"company_short_name,omitempty" xml:"company_short_name,omitempty"`
	// 中文标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 工作年限
	WorkYear int64 `json:"work_year,omitempty" xml:"work_year,omitempty"`
	// 公司名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 状态：0无效，1有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 默认0，1内场顾问，2外场分销
	Role int64 `json:"role,omitempty" xml:"role,omitempty"`
	// 400分机号/虚拟号
	SubPhone string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
	// 主机号
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 头像图片
	HeadUrl string `json:"head_url,omitempty" xml:"head_url,omitempty"`
	// 姓名
	AdviserUserName string `json:"adviser_user_name,omitempty" xml:"adviser_user_name,omitempty"`
	// 手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 类型 1-置业顾问 2-小区专家 3-经纪人
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 顾问id
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
}
