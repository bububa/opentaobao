package tvpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAuthApplyAPIResponse tv支付申请设备授权 API返回值
// taobao.tvpay.auth.apply
//
// 为用户在指定设备上申请支付授权
type TaobaoTvpayAuthApplyAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayAuthApplyAPIResponseModel
}

// TaobaoTvpayAuthApplyAPIResponseModel is tv支付申请设备授权 成功返回结果
type TaobaoTvpayAuthApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_auth_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
