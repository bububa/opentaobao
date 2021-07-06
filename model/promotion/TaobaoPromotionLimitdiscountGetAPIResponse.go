package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionLimitdiscountGetAPIResponse 限时打折查询 API返回值
// taobao.promotion.limitdiscount.get
//
// 分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。
type TaobaoPromotionLimitdiscountGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionLimitdiscountGetAPIResponseModel
}

// TaobaoPromotionLimitdiscountGetAPIResponseModel is 限时打折查询 成功返回结果
type TaobaoPromotionLimitdiscountGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_limitdiscount_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 限时打折列表。
	LimitDiscountList []LimitDiscount `json:"limit_discount_list,omitempty" xml:"limit_discount_list>limit_discount,omitempty"`
	// 满足该查询条件的限时打折总数量。
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
