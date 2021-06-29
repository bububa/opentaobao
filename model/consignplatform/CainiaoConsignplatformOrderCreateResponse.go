package consignplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟发货工作台创建订单 API返回值 
cainiao.consignplatform.order.create

菜鸟发货工作台，商家或者isv通过api进行订单写入操作
*/
type CainiaoConsignplatformOrderCreateAPIResponse struct {
    model.CommonResponse
    CainiaoConsignplatformOrderCreateResponse
}

// 菜鸟发货工作台创建订单 成功返回结果
type CainiaoConsignplatformOrderCreateResponse struct {
    XMLName xml.Name `xml:"cainiao_consignplatform_order_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 失败信息
    FailMessage   string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
    // 失败code
    FailCode   string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
    // 创建是否成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
