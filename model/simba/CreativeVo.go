package simba

import (
	"sync"
)

// CreativeVo 结构体
type CreativeVo struct {
	// 单品创意素材
	MaterialList []ItemMaterialVo `json:"material_list,omitempty" xml:"material_list>item_material_vo,omitempty"`
	// 跳转地址媒体类型,1:PC,2:无线,3:PC和无线
	ClickurlMediaTypeList []string `json:"clickurl_media_type_list,omitempty" xml:"clickurl_media_type_list>string,omitempty"`
	// 审核信息
	AdzoneAuditResultList []AuditVo `json:"adzone_audit_result_list,omitempty" xml:"adzone_audit_result_list>audit_vo,omitempty"`
	// 创意预览信息
	PreviewList []CreativeAdzonePreviewVo `json:"preview_list,omitempty" xml:"preview_list>creative_adzone_preview_vo,omitempty"`
	// 单品创意视频
	ItemVideos []ItemVideoVo `json:"item_videos,omitempty" xml:"item_videos>item_video_vo,omitempty"`
	// 创意关联宝贝
	ItemIdList []int64 `json:"item_id_list,omitempty" xml:"item_id_list>int64,omitempty"`
	// 关键词推广素材
	Children []CreativeChildrenVo `json:"children,omitempty" xml:"children>creative_children_vo,omitempty"`
	// 创意名称
	CreativeName string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
	// 前端创意中心id
	CreativeCenterId string `json:"creative_center_id,omitempty" xml:"creative_center_id,omitempty"`
	// 从海棠获取的地址
	JsInHtml string `json:"js_in_html,omitempty" xml:"js_in_html,omitempty"`
	// 从海棠获取的templateData
	TemplateData string `json:"template_data,omitempty" xml:"template_data,omitempty"`
	// 创意尺寸
	CreativeSize string `json:"creative_size,omitempty" xml:"creative_size,omitempty"`
	// 创意开始时间，用户设置
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 创意结束时间，用户设置
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 扩展数据
	ExtendData string `json:"extend_data,omitempty" xml:"extend_data,omitempty"`
	// 推广标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 一跳地址
	ImagePath string `json:"image_path,omitempty" xml:"image_path,omitempty"`
	// 二跳地址
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 视频时长
	VideoDuration string `json:"video_duration,omitempty" xml:"video_duration,omitempty"`
	// 视频地址
	VideoPath string `json:"video_path,omitempty" xml:"video_path,omitempty"`
	// 视频封面地址
	VideoImagePath string `json:"video_image_path,omitempty" xml:"video_image_path,omitempty"`
	// tab类型,101:竖版大图800x1200,102:竖版视频800x1200,103:竖版长图513x750,104:竖版视频512x750,105:竖版模板513x750,113:方图800x800,114:方视频800x800,115:横版大图320x90,116:横版大图240x100,117:横版大图140x120,118:横版大图730x350,119:横版大图846x220,120:横版大图780x210,121:横版大图840x90,122:横版大图180x100,124:横版大图980x200,125:横版大图820x90,126:竖版视频750x1000,127:竖版视频720x960,128:竖版视频1080x1440,129:竖版视频720x1280,130:竖版视频1080x1920
	TabType string `json:"tab_type,omitempty" xml:"tab_type,omitempty"`
	// tab类型名称,101:竖版大图800x1200,102:竖版视频800x1200,103:竖版长图513x750,104:竖版视频512x750,105:竖版模板513x750,113:方图800x800,114:方视频800x800,115:横版大图320x90,116:横版大图240x100,117:横版大图140x120,118:横版大图730x350,119:横版大图846x220,120:横版大图780x210,121:横版大图840x90,122:横版大图180x100,124:横版大图980x200,125:横版大图820x90,126:竖版视频750x1000,127:竖版视频720x960,128:竖版视频1080x1440,129:竖版视频720x1280,130:竖版视频1080x1920
	TabTypeName string `json:"tab_type_name,omitempty" xml:"tab_type_name,omitempty"`
	// 创意类型名称,1:自定义创意,1:主副图创意,1:主图视频创意,5:智能创意
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 创意类型名称,1:自定义创意,1:主副图创意,1:主图视频创意,5:智能创意
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 创意是否为主图视频创意，异步创建,1:是,其他不是
	MainPicVideo string `json:"main_pic_video,omitempty" xml:"main_pic_video,omitempty"`
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意分类,1:静态创意,2:智能创意,3:PSD创意,4:动态文案创意,5:自动化创意,6:智能创意,7:综合静态创意,8:综合智能创意,9:自动直播创意,10:元素化创意,11:素材类创意,12:钻展多图创意
	CreativeType int64 `json:"creative_type,omitempty" xml:"creative_type,omitempty"`
	// 创意来源,1:本地上传,2:BannerMaker,3:CreativeCenter,4:LuBan,5:直播广场,6:活动招商,7:海棠,8:原生创意
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 创意格式,1:文字,2:图片,3:Flash,4:视频,5:文字链,9:FLASH不遮盖,10:创意模板,11:视频直播,12:微视频,203:智能创意-三图创意
	Format int64 `json:"format,omitempty" xml:"format,omitempty"`
	// 创意模板ID,
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 外部实体id,eg:前端创意中心id
	OuterId int64 `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 审核信息
	Audit *AuditVo `json:"audit,omitempty" xml:"audit,omitempty"`
	// 创意状态,1:ON,-1:ARCHIVE
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 推广主体
	CreativePromotionEntity *CreativePromotionEntityVo `json:"creative_promotion_entity,omitempty" xml:"creative_promotion_entity,omitempty"`
	// 物料图片信息
	MaterialImageInfo *MaterialImageInfoVo `json:"material_image_info,omitempty" xml:"material_image_info,omitempty"`
	// 视频id
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
}

var poolCreativeVo = sync.Pool{
	New: func() any {
		return new(CreativeVo)
	},
}

// GetCreativeVo() 从对象池中获取CreativeVo
func GetCreativeVo() *CreativeVo {
	return poolCreativeVo.Get().(*CreativeVo)
}

// ReleaseCreativeVo 释放CreativeVo
func ReleaseCreativeVo(v *CreativeVo) {
	v.MaterialList = v.MaterialList[:0]
	v.ClickurlMediaTypeList = v.ClickurlMediaTypeList[:0]
	v.AdzoneAuditResultList = v.AdzoneAuditResultList[:0]
	v.PreviewList = v.PreviewList[:0]
	v.ItemVideos = v.ItemVideos[:0]
	v.ItemIdList = v.ItemIdList[:0]
	v.Children = v.Children[:0]
	v.CreativeName = ""
	v.CreativeCenterId = ""
	v.JsInHtml = ""
	v.TemplateData = ""
	v.CreativeSize = ""
	v.StartTime = ""
	v.EndTime = ""
	v.CreateTime = ""
	v.UpdateTime = ""
	v.ExtendData = ""
	v.Title = ""
	v.ImagePath = ""
	v.ClickUrl = ""
	v.VideoDuration = ""
	v.VideoPath = ""
	v.VideoImagePath = ""
	v.TabType = ""
	v.TabTypeName = ""
	v.TypeName = ""
	v.Type = ""
	v.MainPicVideo = ""
	v.CreativeId = 0
	v.CreativeType = 0
	v.Source = 0
	v.Format = 0
	v.TemplateId = 0
	v.OuterId = 0
	v.Audit = nil
	v.OnlineStatus = 0
	v.CreativePromotionEntity = nil
	v.MaterialImageInfo = nil
	v.VideoId = 0
	poolCreativeVo.Put(v)
}
