package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发品资质校验 API返回值 
taobao.item.permit.check

对淘宝商品发品、编辑前的预校验接口
*/
type TaobaoItemPermitCheckAPIResponse struct {
    model.CommonResponse
    TaobaoItemPermitCheckResponse
}

// 发品资质校验 成功返回结果
type TaobaoItemPermitCheckResponse struct {
    XMLName xml.Name `xml:"item_permit_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否成功
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
    // 错误码
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
}
