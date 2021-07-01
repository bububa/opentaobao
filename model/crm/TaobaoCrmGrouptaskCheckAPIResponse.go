package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分组任务是否完成 API返回值 
taobao.crm.grouptask.check

检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组<br/>若分组上有任务则该属性不能被操作。
*/
type TaobaoCrmGrouptaskCheckAPIResponse struct {
    model.CommonResponse
    TaobaoCrmGrouptaskCheckAPIResponseModel
}

// 查询分组任务是否完成 成功返回结果
type TaobaoCrmGrouptaskCheckAPIResponseModel struct {
    XMLName xml.Name `xml:"crm_grouptask_check_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 异步任务是否完成，true表示完成
    IsFinished   bool `json:"is_finished,omitempty" xml:"is_finished,omitempty"`
}
