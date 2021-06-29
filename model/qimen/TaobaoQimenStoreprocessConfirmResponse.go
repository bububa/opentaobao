package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单确认接口 API返回值 
taobao.qimen.storeprocess.confirm

WMS调用奇门的接口,回传仓内加工单创建情况
*/
type TaobaoQimenStoreprocessConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStoreprocessConfirmResponse
}

// 仓内加工单确认接口 成功返回结果
type TaobaoQimenStoreprocessConfirmResponse struct {
    XMLName xml.Name `xml:"qimen_storeprocess_confirm_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`
}
