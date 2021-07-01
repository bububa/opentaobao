package lstvending

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售货机订单物流信息回传 API返回值 
alibaba.lst.vending.order.update

零售通与设备供应商进行订单对接，通过此接口回流订单物流信息。
*/
type AlibabaLstVendingOrderUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaLstVendingOrderUpdateAPIResponseModel
}

// 自动售货机订单物流信息回传 成功返回结果
type AlibabaLstVendingOrderUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_lst_vending_order_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
    // 错误码
    ErrorNo   string `json:"error_no,omitempty" xml:"error_no,omitempty"`
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 是否执行了更新操作
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
}
