package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获取本地模板信息 APIResponse
taobao.tanx.nativetemplates.get

根据传入的本地模板ID批量返回本地模板
*/
type TaobaoTanxNativetemplatesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxNativetemplatesGetResponse `json:"taobao_tanx_nativetemplates_get_response,omitempty"`
}

type TaobaoTanxNativetemplatesGetResponse struct {

    // 是否成功
    IsOk   bool `json:"is_ok,omitempty"`

    // 本地模板列表
    NativeTemplateList   []NativeTemplateDto `json:"native_template_list,omitempty"`

}