package scs

import (
	"sync"
)

// CreativeBindResultTopDto 结构体
type CreativeBindResultTopDto struct {
	// upgradeReportInfoList
	UpgradeReportInfoList []ReportResultTopDto `json:"upgrade_report_info_list,omitempty" xml:"upgrade_report_info_list>report_result_top_dto,omitempty"`
	// adgroupName
	AdgroupName string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
	// campaignName
	CampaignName string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
	// 创意名称
	CreativeName string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
	// clickUrl
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 创意标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// titleTag
	TitleTag string `json:"title_tag,omitempty" xml:"title_tag,omitempty"`
	// imgUrl
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 审核原因
	AuditReason string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty"`
	// creativeSize
	CreativeSize string `json:"creative_size,omitempty" xml:"creative_size,omitempty"`
	// 审核状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 创意type
	CreativeType string `json:"creative_type,omitempty" xml:"creative_type,omitempty"`
	// 创意类型描述
	FormatName string `json:"format_name,omitempty" xml:"format_name,omitempty"`
	// templateData
	TemplateData string `json:"template_data,omitempty" xml:"template_data,omitempty"`
	// jsContentsMap
	JsContentsMap string `json:"js_contents_map,omitempty" xml:"js_contents_map,omitempty"`
	// displayUrl
	DisplayUrl string `json:"display_url,omitempty" xml:"display_url,omitempty"`
	// imgPath
	ImgPath string `json:"img_path,omitempty" xml:"img_path,omitempty"`
	// 视频创意checksum
	VideoCheckSum string `json:"video_check_sum,omitempty" xml:"video_check_sum,omitempty"`
	// adgroupId
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// campaignId
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// creativeId
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 图片类型 1用户自己上传 2主图 3附图
	ImgType int64 `json:"img_type,omitempty" xml:"img_type,omitempty"`
	// 审核状态
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 外部实体id
	OuterId int64 `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 创意类型
	Format int64 `json:"format,omitempty" xml:"format,omitempty"`
	// 模板id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 模板id
	TemplatePackageId int64 `json:"template_package_id,omitempty" xml:"template_package_id,omitempty"`
	// source
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// creativeLevel
	CreativeLevel int64 `json:"creative_level,omitempty" xml:"creative_level,omitempty"`
	// catId
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 包类型（0：非包，1：标尺包,2：包）
	PackageType int64 `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// canEdit
	CanEdit bool `json:"can_edit,omitempty" xml:"can_edit,omitempty"`
}

var poolCreativeBindResultTopDto = sync.Pool{
	New: func() any {
		return new(CreativeBindResultTopDto)
	},
}

// GetCreativeBindResultTopDto() 从对象池中获取CreativeBindResultTopDto
func GetCreativeBindResultTopDto() *CreativeBindResultTopDto {
	return poolCreativeBindResultTopDto.Get().(*CreativeBindResultTopDto)
}

// ReleaseCreativeBindResultTopDto 释放CreativeBindResultTopDto
func ReleaseCreativeBindResultTopDto(v *CreativeBindResultTopDto) {
	v.UpgradeReportInfoList = v.UpgradeReportInfoList[:0]
	v.AdgroupName = ""
	v.CampaignName = ""
	v.CreativeName = ""
	v.ClickUrl = ""
	v.Title = ""
	v.TitleTag = ""
	v.ImgUrl = ""
	v.AuditReason = ""
	v.CreativeSize = ""
	v.Status = ""
	v.CreativeType = ""
	v.FormatName = ""
	v.TemplateData = ""
	v.JsContentsMap = ""
	v.DisplayUrl = ""
	v.ImgPath = ""
	v.VideoCheckSum = ""
	v.AdgroupId = 0
	v.CampaignId = 0
	v.CreativeId = 0
	v.ImgType = 0
	v.AuditStatus = 0
	v.OuterId = 0
	v.Format = 0
	v.TemplateId = 0
	v.TemplatePackageId = 0
	v.Source = 0
	v.CreativeLevel = 0
	v.CatId = 0
	v.PackageType = 0
	v.CanEdit = false
	poolCreativeBindResultTopDto.Put(v)
}
