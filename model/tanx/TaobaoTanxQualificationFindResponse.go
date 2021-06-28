package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资质查询接口 APIResponse
taobao.tanx.qualification.find

资质查询接口
*/
type TaobaoTanxQualificationFindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tanx_qualification_find_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"