package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑专题 APIResponse
yunos.tvpubadmin.manage.topic.edit

编辑专题
*/
type YunosTvpubadminManageTopicEditAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicEditResponse
}

type YunosTvpubadminManageTopicEditResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_edit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
