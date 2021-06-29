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
    TaobaoTanxQualificationAddResponse
}

type TaobaoTanxQualificationAddResponse struct {
    XMLName xml.Name `xml:"tanx_qualification_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回id对应列表
    
    IdList   []int64 `json:"id_list,omitempty" xml:"id_list>int64,omitempty"`
    
    
}
