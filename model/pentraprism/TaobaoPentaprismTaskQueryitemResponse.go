package pentraprism

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务当前进度 APIResponse
taobao.pentaprism.task.queryitem

外网用户查询五棱镜任务系统当前进度
*/
type TaobaoPentaprismTaskQueryitemAPIResponse struct {
    model.CommonResponse
    TaobaoPentaprismTaskQueryitemResponse
}

type TaobaoPentaprismTaskQueryitemResponse struct {
    XMLName xml.Name `xml:"pentaprism_task_queryitem_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // TOP接口标准出参
    
    Result   *TaskResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
