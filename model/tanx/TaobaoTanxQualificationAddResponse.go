package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提交资质接口 APIResponse
taobao.tanx.qualification.add

dsp客户提交客户资质和行业资质
*/
type TaobaoTanxQualificationAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxQualificationAddResponse `json:"taobao_tanx_qualification_add_response,omitempty"`
}

type TaobaoTanxQualificationAddResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 返回id对应列表
    IdList   []Number `json:"id_list,omitempty"`

}
