package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取消息队列积压情况 API请求
taobao.tmc.queue.get

根据appkey和groupName获取消息队列积压情况
*/
type TaobaoTmcQueueGetRequest struct {
    model.Params
    // TMC组名
    _groupName   string
}

// 初始化TaobaoTmcQueueGetRequest对象
func NewTaobaoTmcQueueGetRequest() *TaobaoTmcQueueGetRequest{
    return &TaobaoTmcQueueGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcQueueGetRequest) GetApiMethodName() string {
    return "taobao.tmc.queue.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcQueueGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupName Setter
// TMC组名
func (r *TaobaoTmcQueueGetRequest) SetGroupName(_groupName string) error {
    r._groupName = _groupName
    r.Set("group_name", _groupName)
    return nil
}

// GroupName Getter
func (r TaobaoTmcQueueGetRequest) GetGroupName() string {
    return r._groupName
}
