package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspayregistercreateAPIResponse 提交支付服务开通 API返回值
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
type TaobaoalitripaxintranspayregistercreateAPIResponse struct {
	model.CommonResponse
	TaobaoalitripaxintranspayregistercreateAPIResponseModel
}

// TaobaoalitripaxintranspayregistercreateAPIResponseModel is 提交支付服务开通 成功返回结果
type TaobaoalitripaxintranspayregistercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoalitripaxintranspayregistercreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
