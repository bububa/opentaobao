package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消轻任务 APIRequest
taobao.qianniu.task.cancel

由任务发起者调用
*/
type TaobaoQianniuTaskCancelRequest struct {
    model.Params

    // 任务元数据ID
    metaId   int64 

    // 任务备注
    memo   string 

}

func NewTaobaoQianniuTaskCancelRequest() *TaobaoQianniuTaskCancelRequest{
    return &TaobaoQianniuTaskCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTaskCancelRequest) GetApiMethodName() string {
    return "taobao.qianniu.task.cancel"
}

func (r TaobaoQianniuTaskCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTaskCancelRequest) SetMetaId(metaId int64) error {
    r.metaId = metaId
    r.Set("meta_id", metaId)
    return nil
}

func (r TaobaoQianniuTaskCancelRequest) GetMetaId() int64 {
    return r.metaId
}

func (r *TaobaoQianniuTaskCancelRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

func (r TaobaoQianniuTaskCancelRequest) GetMemo() string {
    return r.memo
}

