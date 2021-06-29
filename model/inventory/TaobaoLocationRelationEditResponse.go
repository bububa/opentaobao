package inventory

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
地点关联关系增量编辑 APIResponse
taobao.location.relation.edit

地点关联关系增量编辑
*/
type TaobaoLocationRelationEditAPIResponse struct {
    model.CommonResponse
    TaobaoLocationRelationEditResponse
}

type TaobaoLocationRelationEditResponse struct {
    XMLName xml.Name `xml:"location_relation_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误码
    
    Errorcode   string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`

    
}
