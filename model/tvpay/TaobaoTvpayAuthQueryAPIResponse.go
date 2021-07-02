package tvpay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAuthQueryAPIResponse tv支付授权查询 API返回值
// taobao.tvpay.auth.query
//
// 查询该用户在指定设备上是否有支付授权
type TaobaoTvpayAuthQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayAuthQueryAPIResponseModel
}

// TaobaoTvpayAuthQueryAPIResponseModel is tv支付授权查询 成功返回结果
type TaobaoTvpayAuthQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_auth_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
