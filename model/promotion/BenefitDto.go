package promotion

import (
	"sync"
)

// BenefitDto 结构体
type BenefitDto struct {
	// 权益code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 权益类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 权益名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 权益描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 权益状态，online为上线，offline为下线，invalid为失效
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 权益发放开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 权益发放结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 发放模式,fixedAmount,dynamicAmount
	SendMode string `json:"send_mode,omitempty" xml:"send_mode,omitempty"`
	// 创建者昵称
	CreatorUserName string `json:"creator_user_name,omitempty" xml:"creator_user_name,omitempty"`
	// 修改者昵称
	ModifierUserName string `json:"modifier_user_name,omitempty" xml:"modifier_user_name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 总库存
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 发放量
	Bestow int64 `json:"bestow,omitempty" xml:"bestow,omitempty"`
	// 关联的权益池id
	BenefitPoolId int64 `json:"benefit_pool_id,omitempty" xml:"benefit_pool_id,omitempty"`
	// 总个数，当sendMode为dynamicAmount时表示已发放面额，单位为分
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 权益id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolBenefitDto = sync.Pool{
	New: func() any {
		return new(BenefitDto)
	},
}

// GetBenefitDto() 从对象池中获取BenefitDto
func GetBenefitDto() *BenefitDto {
	return poolBenefitDto.Get().(*BenefitDto)
}

// ReleaseBenefitDto 释放BenefitDto
func ReleaseBenefitDto(v *BenefitDto) {
	v.Code = ""
	v.Type = ""
	v.Name = ""
	v.Description = ""
	v.Status = ""
	v.StartDate = ""
	v.EndDate = ""
	v.Feature = ""
	v.SendMode = ""
	v.CreatorUserName = ""
	v.ModifierUserName = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Total = 0
	v.Bestow = 0
	v.BenefitPoolId = 0
	v.TotalNum = 0
	v.Id = 0
	poolBenefitDto.Put(v)
}
