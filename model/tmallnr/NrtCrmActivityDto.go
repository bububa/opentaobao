package tmallnr

import (
	"sync"
)

// NrtCrmActivityDto 结构体
type NrtCrmActivityDto struct {
	// 有价券DTO
	SceneActivityList []NrtCrmSceneActivityDto `json:"scene_activity_list,omitempty" xml:"scene_activity_list>nrt_crm_scene_activity_dto,omitempty"`
	// 头图
	BannerUrl string `json:"banner_url,omitempty" xml:"banner_url,omitempty"`
	// 活动状态
	StatusStr string `json:"status_str,omitempty" xml:"status_str,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 活动规则
	Rule string `json:"rule,omitempty" xml:"rule,omitempty"`
	// 权益信息
	CertificateRights string `json:"certificate_rights,omitempty" xml:"certificate_rights,omitempty"`
	// 同城站ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 活动状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 站外页面ID
	PageId int64 `json:"page_id,omitempty" xml:"page_id,omitempty"`
	// 电子凭证ID
	TmpCertificateId int64 `json:"tmp_certificate_id,omitempty" xml:"tmp_certificate_id,omitempty"`
	// 直播信息
	NrtCrmLiveDto *NrtCrmLiveDto `json:"nrt_crm_live_dto,omitempty" xml:"nrt_crm_live_dto,omitempty"`
	// 卖场ID
	MallId int64 `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
}

var poolNrtCrmActivityDto = sync.Pool{
	New: func() any {
		return new(NrtCrmActivityDto)
	},
}

// GetNrtCrmActivityDto() 从对象池中获取NrtCrmActivityDto
func GetNrtCrmActivityDto() *NrtCrmActivityDto {
	return poolNrtCrmActivityDto.Get().(*NrtCrmActivityDto)
}

// ReleaseNrtCrmActivityDto 释放NrtCrmActivityDto
func ReleaseNrtCrmActivityDto(v *NrtCrmActivityDto) {
	v.SceneActivityList = v.SceneActivityList[:0]
	v.BannerUrl = ""
	v.StatusStr = ""
	v.EndTime = ""
	v.StartTime = ""
	v.Title = ""
	v.Description = ""
	v.Rule = ""
	v.CertificateRights = ""
	v.CityId = 0
	v.Status = 0
	v.SellerId = 0
	v.ActivityId = 0
	v.PageId = 0
	v.TmpCertificateId = 0
	v.NrtCrmLiveDto = nil
	v.MallId = 0
	poolNrtCrmActivityDto.Put(v)
}
