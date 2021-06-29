package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取本地模板信息 APIResponse
taobao.tanx.nativetemplates.get

根据传入的本地模板ID批量返回本地模板
*/
type TaobaoTanxNativetemplatesGetAPIResponse struct {
    model.CommonResponse
    TaobaoTanxNativetemplatesGetResponse
}

type TaobaoTanxNativetemplatesGetResponse struct {
    XMLName xml.Name `xml:"tanx_nativetemplates_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsOk   bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`

    
    // 本地模板列表
    
    NativeTemplateList   []NativeTemplateDto `json:"native_template_list,omitempty" xml:"native_template_list>native_template_dto,omitempty"`
    
    
}
