package qianniu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新任务元数据 APIResponse
taobao.qianniu.taskmeta.update

由任务发起者调用
*/
type TaobaoQianniuTaskmetaUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoQianniuTaskmetaUpdateResponse
}

type TaobaoQianniuTaskmetaUpdateResponse struct {
    XMLName xml.Name `xml:"qianniu_taskmeta_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
