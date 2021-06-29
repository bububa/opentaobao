package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专题 APIResponse
yunos.tvpubadmin.manage.topic.add

专题新增
*/
type YunosTvpubadminManageTopicAddAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminManageTopicAddResponse
}

type YunosTvpubadminManageTopicAddResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_manage_topic_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
