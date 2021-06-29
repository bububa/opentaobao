package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店接单接口 API返回值 
taobao.omniorder.store.accpeted

ISV Pos端门店接单，通知星盘
*/
type TaobaoOmniorderStoreAccpetedAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreAccpetedResponse
}

// Pos端门店接单接口 成功返回结果
type TaobaoOmniorderStoreAccpetedResponse struct {
    XMLName xml.Name `xml:"omniorder_store_accpeted_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误内容
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
