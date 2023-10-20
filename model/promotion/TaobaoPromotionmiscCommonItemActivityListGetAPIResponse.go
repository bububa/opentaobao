package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscCommonItemActivityListGetAPIResponse 查询通用单品优惠活动列表 API返回值
// taobao.promotionmisc.common.item.activity.list.get
//
// 查询通用单品优惠活动列表。
type TaobaoPromotionmiscCommonItemActivityListGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscCommonItemActivityListGetAPIResponseModel
}

// TaobaoPromotionmiscCommonItemActivityListGetAPIResponseModel is 查询通用单品优惠活动列表 成功返回结果
type TaobaoPromotionmiscCommonItemActivityListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_activity_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动列表
	ActivityList []CommonItemActivity `json:"activity_list,omitempty" xml:"activity_list>common_item_activity,omitempty"`
	// 数据总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否查询成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
