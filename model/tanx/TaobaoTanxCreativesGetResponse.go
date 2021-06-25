package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获取DSP用户的创意审核结果 APIResponse
taobao.tanx.creatives.get

批量获取DSP用户的创意审核结果
*/
type TaobaoTanxCreativesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxCreativesGetResponse `json:"taobao_tanx_creatives_get_response,omitempty"`
}

type TaobaoTanxCreativesGetResponse struct {

    // 调用的成功信息或失败信息
    Message   string `json:"message,omitempty"`

    // 调用返回码
    TanxErrorCode   int64 `json:"tanx_error_code,omitempty"`

    // 调用是否成功
    IsOk   bool `json:"is_ok,omitempty"`

    // 返回的创意列表
    Results   []CreativeDto `json:"results,omitempty"`

}
