package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新任务元数据 APIRequest
taobao.qianniu.taskmeta.update

由任务发起者调用
*/
type TaobaoQianniuTaskmetaUpdateRequest struct {
    model.Params

    // 要更新的任务元数据，JSON格式，例如：<br/>meta= {<br/>		"id" : 1,<br/>		"title" : "xxx",<br/>		"content" : "yyyy",<br/>		"biz_sys_Id" : 12232,<br/>		"biz_sys_task_type" : 1212,<br/>		"start_time" : 1380173565480,<br/>		"end_time" : 1380173565480,<br/> "sender_uid" : 213123213,<br/>		"sender_nick" : "tbtest1063",<br/>		"reminder_flag" : 1,<br/>		"finish_strategy" : 1<br/>	}
    meta   string 

}

func NewTaobaoQianniuTaskmetaUpdateRequest() *TaobaoQianniuTaskmetaUpdateRequest{
    return &TaobaoQianniuTaskmetaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQianniuTaskmetaUpdateRequest) GetApiMethodName() string {
    return "taobao.qianniu.taskmeta.update"
}

func (r TaobaoQianniuTaskmetaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQianniuTaskmetaUpdateRequest) SetMeta(meta string) error {
    r.meta = meta
    r.Set("meta", meta)
    return nil
}

func (r TaobaoQianniuTaskmetaUpdateRequest) GetMeta() string {
    return r.meta
}

