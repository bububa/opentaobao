package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改资质接口 APIResponse
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改
*/
type TaobaoTanxQualificationModifyAPIResponse struct {
    model.CommonResponse
    TaobaoTanxQualificationModifyResponse
}

type TaobaoTanxQualificationModifyResponse struct {
    XMLName xml.Name `xml:"tanx_qualification_modify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
