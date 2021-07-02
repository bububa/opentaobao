package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeActivitySyncAPIResponse 尖货交易活动信息同步 API返回值
// taobao.opentrade.activity.sync
//
// 尖货交易活动信息配置，创建或更新活动信息
// 在活动时间开始前，所有用户（包括标记可购买的用户），无法购买商品；
// 在活动时间内，标记可购买的用户可在小程序中跳转下单页，完成购买；
// 在活动结束后，对限购不再限制，平台开放购买，用户可在小程序内、商品详情、购物车下单购买；
type TaobaoOpentradeActivitySyncAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeActivitySyncAPIResponseModel
}

// TaobaoOpentradeActivitySyncAPIResponseModel is 尖货交易活动信息同步 成功返回结果
type TaobaoOpentradeActivitySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_activity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动关联的商品列表，使用逗号(,)分割
	ItemIds string `json:"item_ids,omitempty" xml:"item_ids,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 创建活动的appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 排队活动ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 最近修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
