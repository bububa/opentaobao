package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加活动参与的商品 API返回值 
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
*/
type TaobaoPromotionmiscActivityRangeAddAPIResponse struct {
    model.CommonResponse
    TaobaoPromotionmiscActivityRangeAddAPIResponseModel
}

// 增加活动参与的商品 成功返回结果
type TaobaoPromotionmiscActivityRangeAddAPIResponseModel struct {
    XMLName xml.Name `xml:"promotionmisc_activity_range_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 增加商品范围是否成功。
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
