package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审核创意状态 APIResponse
taobao.tanx.creative.get

获取单个审核创意状态
*/
type TaobaoTanxCreativeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxCreativeGetResponse `json:"taobao_tanx_creative_get_response,omitempty"`
}

type TaobaoTanxCreativeGetResponse struct {

    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty"`

    // 调用返回码
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`

    // 是否成功
    IsOk   bool `json:"is_ok,omitempty"`

    // 创意查询返回结果列表
    Result   *CreativeAuditDto `json:"result,omitempty"`

}
