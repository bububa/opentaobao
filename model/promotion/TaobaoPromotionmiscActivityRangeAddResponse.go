package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
增加活动参与的商品 APIResponse
taobao.promotionmisc.activity.range.add

增加活动参与的商品，部分商品参与的活动，最大支持指定150个商品。
*/
type TaobaoPromotionmiscActivityRangeAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"promotionmisc_activity_range_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 增加商品范围是否成功。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"