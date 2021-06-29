package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分组任务是否完成 APIResponse
taobao.crm.grouptask.check

检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组<br/>若分组上有任务则该属性不能被操作。
*/
type TaobaoCrmGrouptaskCheckAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrouptaskCheckResponse
}

type TaobaoCrmGrouptaskCheckResponse struct {
    XMLName xml.Name `xml:"crm_grouptask_check_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步任务是否完成，true表示完成
    
    IsFinished   bool `json:"is_finished,omitempty" xml:"is_finished,omitempty"`

    
}
