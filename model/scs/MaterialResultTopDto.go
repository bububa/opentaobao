package scs

// MaterialResultTopDto 结构体
type MaterialResultTopDto struct {
	// itemIdListInLive
	ItemIdListInLiveList []int64 `json:"item_id_list_in_live_list,omitempty" xml:"item_id_list_in_live_list>int64,omitempty"`
	// multiImageURL
	MultiImageURLList []string `json:"multi_image_u_r_l_list,omitempty" xml:"multi_image_u_r_l_list>string,omitempty"`
	// recommendTagList
	RecommendTagList []string `json:"recommend_tag_list,omitempty" xml:"recommend_tag_list>string,omitempty"`
	// showTagInfoList
	ShowTagInfoList []ShowTagTopDto `json:"show_tag_info_list,omitempty" xml:"show_tag_info_list>show_tag_top_dto,omitempty"`
	// mediaTypeNames
	MediaTypeNamesList []string `json:"media_type_names_list,omitempty" xml:"media_type_names_list>string,omitempty"`
	// reportInfoList
	ReportInfoList []ReportResultTopDto `json:"report_info_list,omitempty" xml:"report_info_list>report_result_top_dto,omitempty"`
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// linkUrl
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// imgUrl
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// liveStartTime
	LiveStartTime string `json:"live_start_time,omitempty" xml:"live_start_time,omitempty"`
	// liveEndTime
	LiveEndTime string `json:"live_end_time,omitempty" xml:"live_end_time,omitempty"`
	// publishTime
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// darenName
	DarenName string `json:"daren_name,omitempty" xml:"daren_name,omitempty"`
	// wirelessLongImageURL
	WirelessLongImageUrl string `json:"wireless_long_image_url,omitempty" xml:"wireless_long_image_url,omitempty"`
	// categoryId
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// price
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// starts
	Starts string `json:"starts,omitempty" xml:"starts,omitempty"`
	// createTime
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// modifyTime
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// totalOrderAmount
	TotalOrderAmount string `json:"total_order_amount,omitempty" xml:"total_order_amount,omitempty"`
	// materialId
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// materialType
	MaterialType int64 `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// contentEntityType
	ContentEntityType int64 `json:"content_entity_type,omitempty" xml:"content_entity_type,omitempty"`
	// contentBizType
	ContentBizType int64 `json:"content_biz_type,omitempty" xml:"content_biz_type,omitempty"`
	// quantity
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// bidCount
	BidCount int64 `json:"bid_count,omitempty" xml:"bid_count,omitempty"`
	// categoryLevel1
	CategoryLevel1 int64 `json:"category_level1,omitempty" xml:"category_level1,omitempty"`
	// recommendScore
	RecommendScore int64 `json:"recommend_score,omitempty" xml:"recommend_score,omitempty"`
	// videoDuration
	VideoDuration int64 `json:"video_duration,omitempty" xml:"video_duration,omitempty"`
	// launchCount
	LaunchCount int64 `json:"launch_count,omitempty" xml:"launch_count,omitempty"`
	// timingStart
	TimingStart bool `json:"timing_start,omitempty" xml:"timing_start,omitempty"`
	// unlockRptIndex
	UnlockRptIndex bool `json:"unlock_rpt_index,omitempty" xml:"unlock_rpt_index,omitempty"`
}
