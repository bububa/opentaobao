package axintrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信创建退款单 API返回值 
taobao.alitrip.axin.trans.refund.create

阿信供销平台-创建退款单服务
*/
type TaobaoAlitripAxinTransRefundCreateAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripAxinTransRefundCreateAPIResponseModel
}

// 阿信创建退款单 成功返回结果
type TaobaoAlitripAxinTransRefundCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_axin_trans_refund_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoAlitripAxinTransRefundCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
