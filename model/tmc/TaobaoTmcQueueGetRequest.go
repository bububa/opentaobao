package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取消息队列积压情况 APIRequest
taobao.tmc.queue.get

根据appkey和groupName获取消息队列积压情况
*/
type TaobaoTmcQueueGetRequest struct {
    model.Params

    // TMC组名
    groupName   string 

}

func NewTaobaoTmcQueueGetRequest() *TaobaoTmcQueueGetRequest{
    return &TaobaoTmcQueueGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcQueueGetRequest) GetApiMethodName() string {
    return "taobao.tmc.queue.get"
}

func (r TaobaoTmcQueueGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcQueueGetRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r TaobaoTmcQueueGetRequest) GetGroupName() string {
    return r.groupName
}

