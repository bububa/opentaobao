package alihouse

import (
	"sync"
)

// SyncProjectNewReviewDraftDto 结构体
type SyncProjectNewReviewDraftDto struct {
	// 头像
	Head string `json:"head,omitempty" xml:"head,omitempty"`
	// 外部测评id
	OuterReviewId string `json:"outer_review_id,omitempty" xml:"outer_review_id,omitempty"`
	// 发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 项目不足
	Defect string `json:"defect,omitempty" xml:"defect,omitempty"`
	// 项目亮点
	Lights string `json:"lights,omitempty" xml:"lights,omitempty"`
	// 公共交通
	PublicTraffic string `json:"public_traffic,omitempty" xml:"public_traffic,omitempty"`
	// 样板房细节
	ModelHouse string `json:"model_house,omitempty" xml:"model_house,omitempty"`
	// 户型分析描述
	HouseTypeAnalysis string `json:"house_type_analysis,omitempty" xml:"house_type_analysis,omitempty"`
	// 价格潜力
	PotentialPrice string `json:"potential_price,omitempty" xml:"potential_price,omitempty"`
	// 价格现状
	NowPrice string `json:"now_price,omitempty" xml:"now_price,omitempty"`
	// 其他配套
	OtherSource string `json:"other_source,omitempty" xml:"other_source,omitempty"`
	// 医疗配套
	MedicalSource string `json:"medical_source,omitempty" xml:"medical_source,omitempty"`
	// 教育配套
	EduSource string `json:"edu_source,omitempty" xml:"edu_source,omitempty"`
	// 自驾道路
	MainRoad string `json:"main_road,omitempty" xml:"main_road,omitempty"`
	// 轨道交通
	SubwayDesc string `json:"subway_desc,omitempty" xml:"subway_desc,omitempty"`
	// 板块描述
	BlockDesc string `json:"block_desc,omitempty" xml:"block_desc,omitempty"`
	// 区域描述
	DistrictDesc string `json:"district_desc,omitempty" xml:"district_desc,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 自定义配套
	CustomSource string `json:"custom_source,omitempty" xml:"custom_source,omitempty"`
	// 生活配套描述
	LiveSourceDesc string `json:"live_source_desc,omitempty" xml:"live_source_desc,omitempty"`
	// 交通出行描述
	TrafficTravelDesc string `json:"traffic_travel_desc,omitempty" xml:"traffic_travel_desc,omitempty"`
	// 产品特色
	ProductFeature string `json:"product_feature,omitempty" xml:"product_feature,omitempty"`
	// 产品特色描述
	ProductFeatureDesc string `json:"product_feature_desc,omitempty" xml:"product_feature_desc,omitempty"`
	// 区域板块描述
	DistrictBlockDesc string `json:"district_block_desc,omitempty" xml:"district_block_desc,omitempty"`
	// 适合人群
	SuitablePeople string `json:"suitable_people,omitempty" xml:"suitable_people,omitempty"`
	// 项目描述
	ProjectDesc string `json:"project_desc,omitempty" xml:"project_desc,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 描述
	Describe string `json:"describe,omitempty" xml:"describe,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 户型需要支持多组（1-5组）
	HouseMultiple string `json:"house_multiple,omitempty" xml:"house_multiple,omitempty"`
	// 价格描述
	PriceDesc string `json:"price_desc,omitempty" xml:"price_desc,omitempty"`
	// 公园配套
	ParkSource string `json:"park_source,omitempty" xml:"park_source,omitempty"`
	// 是否是测试数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 外部视频id
	OuterVideoId int64 `json:"outer_video_id,omitempty" xml:"outer_video_id,omitempty"`
	// 是否展示价格走势模块开关
	PriceTrend int64 `json:"price_trend,omitempty" xml:"price_trend,omitempty"`
	// 唯一id
	UniqueId int64 `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
}

var poolSyncProjectNewReviewDraftDto = sync.Pool{
	New: func() any {
		return new(SyncProjectNewReviewDraftDto)
	},
}

// GetSyncProjectNewReviewDraftDto() 从对象池中获取SyncProjectNewReviewDraftDto
func GetSyncProjectNewReviewDraftDto() *SyncProjectNewReviewDraftDto {
	return poolSyncProjectNewReviewDraftDto.Get().(*SyncProjectNewReviewDraftDto)
}

// ReleaseSyncProjectNewReviewDraftDto 释放SyncProjectNewReviewDraftDto
func ReleaseSyncProjectNewReviewDraftDto(v *SyncProjectNewReviewDraftDto) {
	v.Head = ""
	v.OuterReviewId = ""
	v.PublishTime = ""
	v.Defect = ""
	v.Lights = ""
	v.PublicTraffic = ""
	v.ModelHouse = ""
	v.HouseTypeAnalysis = ""
	v.PotentialPrice = ""
	v.NowPrice = ""
	v.OtherSource = ""
	v.MedicalSource = ""
	v.EduSource = ""
	v.MainRoad = ""
	v.SubwayDesc = ""
	v.BlockDesc = ""
	v.DistrictDesc = ""
	v.OuterId = ""
	v.CustomSource = ""
	v.LiveSourceDesc = ""
	v.TrafficTravelDesc = ""
	v.ProductFeature = ""
	v.ProductFeatureDesc = ""
	v.DistrictBlockDesc = ""
	v.SuitablePeople = ""
	v.ProjectDesc = ""
	v.NickName = ""
	v.Describe = ""
	v.Name = ""
	v.HouseMultiple = ""
	v.PriceDesc = ""
	v.ParkSource = ""
	v.IsTest = 0
	v.Status = 0
	v.CityId = 0
	v.OuterVideoId = 0
	v.PriceTrend = 0
	v.UniqueId = 0
	poolSyncProjectNewReviewDraftDto.Put(v)
}
