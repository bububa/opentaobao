package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRefundRefusereasonGetAPIResponse
获取拒绝原因列表 API返回值
taobao.refund.refusereason.get

获取商家拒绝原因列表 */
type TaobaoRefundRefusereasonGetAPIResponse struct {
	model.CommonResponse
	TaobaoRefundRefusereasonGetAPIResponseModel
}

// TaobaoRefundRefusereasonGetAPIResponseModel is 获取拒绝原因列表 成功返回结果
type TaobaoRefundRefusereasonGetAPIResponseModel struct {
	XMLName xml.Name `xml:"refund_refusereason_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖家拒绝原因对象
	Reasons []Reason `json:"reasons,omitempty" xml:"reasons>reason,omitempty"`
	// 原因个数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
