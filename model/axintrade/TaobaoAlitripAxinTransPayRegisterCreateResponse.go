package axintrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交支付服务开通 APIResponse
taobao.alitrip.axin.trans.pay.register.create

阿信供销平台-提交支付服务开通接口
*/
type TaobaoAlitripAxinTransPayRegisterCreateAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripAxinTransPayRegisterCreateResponse
}

type TaobaoAlitripAxinTransPayRegisterCreateResponse struct {
    XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoAlitripAxinTransPayRegisterCreateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
