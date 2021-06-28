package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交资质接口 APIResponse
taobao.tanx.qualification.add

dsp客户提交客户资质和行业资质
*/
type TaobaoTanxQualificationAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_qualification_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"