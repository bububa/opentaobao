package campus

// Building 结构体
type Building struct {
	// 所属公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 所属园区Code
	CampusCode string `json:"campus_code,omitempty" xml:"campus_code,omitempty"`
	// 所属园区名称
	CampusName string `json:"campus_name,omitempty" xml:"campus_name,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// modifier
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// creator
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 所属公司
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 状态，0停用，1启用
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 楼宇排序号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 所属园区ID
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否删除，0未删除，1删除
	IsDelete bool `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
}
