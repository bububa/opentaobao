package axintrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintranspaysigncheckAPIResponse 阿信支付宝验签服务 API返回值
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
type TaobaoalitripaxintranspaysigncheckAPIResponse struct {
	model.CommonResponse
	TaobaoalitripaxintranspaysigncheckAPIResponseModel
}

// TaobaoalitripaxintranspaysigncheckAPIResponseModel is 阿信支付宝验签服务 成功返回结果
type TaobaoalitripaxintranspaysigncheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_pay_sign_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
