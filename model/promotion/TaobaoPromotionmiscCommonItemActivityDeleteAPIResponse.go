package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmisccommonitemactivitydeleteAPIResponse 删除通用单品优惠活动 API返回值
// taobao.promotionmisc.common.item.activity.delete
//
// 删除通用单品优惠活动。
type TaobaopromotionmisccommonitemactivitydeleteAPIResponse struct {
	model.CommonResponse
	TaobaopromotionmisccommonitemactivitydeleteAPIResponseModel
}

// TaobaopromotionmisccommonitemactivitydeleteAPIResponseModel is 删除通用单品优惠活动 成功返回结果
type TaobaopromotionmisccommonitemactivitydeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_common_item_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否删除成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
