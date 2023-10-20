package nrt

import (
	"sync"
)

// NrtBenefitDto 结构体
type NrtBenefitDto struct {
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 中奖概率
	Probability string `json:"probability,omitempty" xml:"probability,omitempty"`
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 权益名称
	BenefitName string `json:"benefit_name,omitempty" xml:"benefit_name,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 权益图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 关联的权益模板id
	OutRelateId int64 `json:"out_relate_id,omitempty" xml:"out_relate_id,omitempty"`
	// 权益id
	BenefitId int64 `json:"benefit_id,omitempty" xml:"benefit_id,omitempty"`
	// 总数
	TotalQuantity int64 `json:"total_quantity,omitempty" xml:"total_quantity,omitempty"`
	// 个人限领数
	PersonLimitCount int64 `json:"person_limit_count,omitempty" xml:"person_limit_count,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 使用渠道
	UseChannel int64 `json:"use_channel,omitempty" xml:"use_channel,omitempty"`
}

var poolNrtBenefitDto = sync.Pool{
	New: func() any {
		return new(NrtBenefitDto)
	},
}

// GetNrtBenefitDto() 从对象池中获取NrtBenefitDto
func GetNrtBenefitDto() *NrtBenefitDto {
	return poolNrtBenefitDto.Get().(*NrtBenefitDto)
}

// ReleaseNrtBenefitDto 释放NrtBenefitDto
func ReleaseNrtBenefitDto(v *NrtBenefitDto) {
	v.GmtModified = ""
	v.BenefitType = ""
	v.Probability = ""
	v.BizCode = ""
	v.Description = ""
	v.GmtCreate = ""
	v.BenefitName = ""
	v.Extra = ""
	v.StartTime = ""
	v.EndTime = ""
	v.ImgUrl = ""
	v.OutRelateId = 0
	v.BenefitId = 0
	v.TotalQuantity = 0
	v.PersonLimitCount = 0
	v.Status = 0
	v.UseChannel = 0
	poolNrtBenefitDto.Put(v)
}
