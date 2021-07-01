package axintrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付入驻重新申请 API返回值 
taobao.alitrip.axin.trans.pay.register.reapply

阿信支付入驻重新申请
用于支付平台驳回，商户提交时的业务场景
*/
type TaobaoAlitripAxinTransPayRegisterReapplyAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel
}

// 阿信支付入驻重新申请 成功返回结果
type TaobaoAlitripAxinTransPayRegisterReapplyAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_axin_trans_pay_register_reapply_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *BaseResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
