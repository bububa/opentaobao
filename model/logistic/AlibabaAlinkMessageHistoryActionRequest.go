package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
操作历史消息 APIRequest
alibaba.alink.message.history.action

阿里智能操作历史消息
*/
type AlibabaAlinkMessageHistoryActionRequest struct {
    model.Params

    // 消息id
    index   string 

    // 删除：delete,已读：read
    action   string 

}

func NewAlibabaAlinkMessageHistoryActionRequest() *AlibabaAlinkMessageHistoryActionRequest{
    return &AlibabaAlinkMessageHistoryActionRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkMessageHistoryActionRequest) GetApiMethodName() string {
    return "alibaba.alink.message.history.action"
}

func (r AlibabaAlinkMessageHistoryActionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkMessageHistoryActionRequest) SetIndex(index string) error {
    r.index = index
    r.Set("index", index)
    return nil
}

func (r AlibabaAlinkMessageHistoryActionRequest) GetIndex() string {
    return r.index
}

func (r *AlibabaAlinkMessageHistoryActionRequest) SetAction(action string) error {
    r.action = action
    r.Set("action", action)
    return nil
}

func (r AlibabaAlinkMessageHistoryActionRequest) GetAction() string {
    return r.action
}

