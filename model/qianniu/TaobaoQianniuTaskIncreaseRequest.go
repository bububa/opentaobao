package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加任务接收人接口 APIRequest
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

func NewTaobaoQianniuTaskIncreaseRequest() *TaobaoQianniuTaskIncreaseRequest{
    return &TaobaoQianniuTaskIncreaseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTaskIncreaseRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.increase"
}

func (r TaobaoQianniuTaskIncreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTaskIncreaseRequest) SetMetadataId(metadataId int64) error {
    r.metadataId = metadataId
    r.Set("metadata_id", metadataId)
    return nil
}

func (r TaobaoQianniuTaskIncreaseRequest) GetMetadataId() int64 {
    return r.metadataId
}

func (r *TaobaoQianniuTaskIncreaseRequest) SetTasks(tasks string) error {
    r.tasks = tasks
    r.Set("tasks", tasks)
    return nil
}

func (r TaobaoQianniuTaskIncreaseRequest) GetTasks() string {
    return r.tasks
}

