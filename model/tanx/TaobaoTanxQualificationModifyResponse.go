package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改资质接口 APIResponse
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改
*/
type TaobaoTanxQualificationModifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTanxQualificationModifyResponse `json:"taobao_tanx_qualification_modify_response,omitempty"`
}

type TaobaoTanxQualificationModifyResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
