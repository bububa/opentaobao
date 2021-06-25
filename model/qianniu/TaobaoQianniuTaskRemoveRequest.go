package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
轻任务删除接口 APIRequest
taobao.qianniu.task.remove

轻任务删除接口。
*/
type TaobaoQianniuTaskRemoveRequest struct {
    model.Params

    // 对于发起人删除一个任务，请使用这个字段，同时清除所有处理人。
    metadataId   int64 

}

func NewTaobaoQianniuTaskRemoveRequest() *TaobaoQianniuTaskRemoveRequest{
    return &TaobaoQianniuTaskRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTaskRemoveRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.remove"
}

func (r TaobaoQianniuTaskRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTaskRemoveRequest) SetMetadataId(metadataId int64) error {
    r.metadataId = metadataId
    r.Set("metadata_id", metadataId)
    return nil
}

func (r TaobaoQianniuTaskRemoveRequest) GetMetadataId() int64 {
    return r.metadataId
}

