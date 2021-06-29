package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审核创意状态 API返回值 
taobao.tanx.creative.get

获取单个审核创意状态
*/
type TaobaoTanxCreativeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxCreativeGetResponse
}

// 获取单个审核创意状态 成功返回结果
type TaobaoTanxCreativeGetResponse struct {
    XMLName xml.Name `xml:"tanx_creative_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 调用返回码
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`
    // 是否成功
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
    // 创意查询返回结果列表
    Result   *CreativeAuditDTO `json:"result,omitempty" xml:"result,omitempty"`
}
