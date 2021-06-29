package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加任务接收人接口 API请求
taobao.qianniu.task.increase

根据任务元id增加任务接收人
*/
type TaobaoQianniuTaskIncreaseRequest struct {
    model.Params
    // 任务元id
    metadataId   int64
    // 任务列表，JSON格式，例如： tasks =[{ "receiver_uid" : 123, "receiver_nick" : "nick"}, { "receiver_uid" : 456, "receiver_nick" : "nick2"} ]
    tasks   string
}

// 初始化TaobaoQianniuTaskIncreaseRequest对象
func NewTaobaoQianniuTaskIncreaseRequest() *TaobaoQianniuTaskIncreaseRequest{
    return &TaobaoQianniuTaskIncreaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskIncreaseRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.increase"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskIncreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MetadataId Setter
// 任务元id
func (r *TaobaoQianniuTaskIncreaseRequest) SetMetadataId(metadataId int64) error {
    r.metadataId = metadataId
    r.Set("metadata_id", metadataId)
    return nil
}

// MetadataId Getter
func (r TaobaoQianniuTaskIncreaseRequest) GetMetadataId() int64 {
    return r.metadataId
}
// Tasks Setter
// 任务列表，JSON格式，例如： tasks =[{ "receiver_uid" : 123, "receiver_nick" : "nick"}, { "receiver_uid" : 456, "receiver_nick" : "nick2"} ]
func (r *TaobaoQianniuTaskIncreaseRequest) SetTasks(tasks string) error {
    r.tasks = tasks
    r.Set("tasks", tasks)
    return nil
}

// Tasks Getter
func (r TaobaoQianniuTaskIncreaseRequest) GetTasks() string {
    return r.tasks
}
