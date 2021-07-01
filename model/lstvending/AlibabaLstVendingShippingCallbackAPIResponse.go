package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
售货机出货回传接口 API返回值 
alibaba.lst.vending.shipping.callback

零售通自动售货机商品出货回传接口，同步商品出库最新状态。
*/
type AlibabaLstVendingShippingCallbackAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendingShippingCallbackAPIResponseModel
}

// 售货机出货回传接口 成功返回结果
type AlibabaLstVendingShippingCallbackAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_lst_vending_shipping_callback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 执行结果
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
}
