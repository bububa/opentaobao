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
    TaobaoTanxQualificationFindResponse
}

type TaobaoTanxQualificationFindResponse struct {
    XMLName xml.Name `xml:"tanx_qualification_find_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 返回的资质内容dto
    
    QualificationList   []QualificationDto `json:"qualification_list,omitempty" xml:"qualification_list>qualification_dto,omitempty"`
    
    
    // 查询返回总条数
    
    Count   string `json:"count,omitempty" xml:"count,omitempty"`

    
}
