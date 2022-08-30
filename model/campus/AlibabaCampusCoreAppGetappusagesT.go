package campus

// AlibabaCampusCoreAppGetappusagesT 结构体
type AlibabaCampusCoreAppGetappusagesT struct {
	// 应用名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 应用图标
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 应用打开模式
	OpenMode string `json:"open_mode,omitempty" xml:"open_mode,omitempty"`
	// 前台url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 后台管理url
	AdminUrl string `json:"admin_url,omitempty" xml:"admin_url,omitempty"`
	// 应用key
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// notifyUrl
	NotifyUrl string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	// 应用使用关系id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 应用id
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 园区ID
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// app排序
	AppOrder int64 `json:"app_order,omitempty" xml:"app_order,omitempty"`
	// 是否收费类应用
	IsCharge bool `json:"is_charge,omitempty" xml:"is_charge,omitempty"`
}
