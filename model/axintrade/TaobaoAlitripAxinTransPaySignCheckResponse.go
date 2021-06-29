package axintrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付宝验签服务 APIResponse
taobao.alitrip.axin.trans.pay.sign.check

阿信支付宝验签服务
*/
type TaobaoAlitripAxinTransPaySignCheckAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripAxinTransPaySignCheckResponse
}

type TaobaoAlitripAxinTransPaySignCheckResponse struct {
    XMLName xml.Name `xml:"alitrip_axin_trans_pay_sign_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
