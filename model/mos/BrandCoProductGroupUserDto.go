package mos

// BrandCoProductGroupUserDto 结构体
type BrandCoProductGroupUserDto struct {
	// 数据创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 数据修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 品牌名称
	GzNameCn string `json:"gz_name_cn,omitempty" xml:"gz_name_cn,omitempty"`
	// 主责任工号
	MainWorkNo string `json:"main_work_no,omitempty" xml:"main_work_no,omitempty"`
	// 主责任姓名
	MainName string `json:"main_name,omitempty" xml:"main_name,omitempty"`
	// 组员姓名
	MemberName string `json:"member_name,omitempty" xml:"member_name,omitempty"`
	// 组员工号
	MemberWorkNo string `json:"member_work_no,omitempty" xml:"member_work_no,omitempty"`
	// 承包方式
	ContractType string `json:"contract_type,omitempty" xml:"contract_type,omitempty"`
	// 开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 组员数据id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
