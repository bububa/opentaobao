package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
为已授权的用户开通消息服务 API请求
taobao.tmc.user.permit

为已授权的用户开通消息服务，授权消息使用。<br/><span style="color:red">注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic</span>
*/
type TaobaoTmcUserPermitRequest struct {
    model.Params
    // 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
    topics   []string
}

// 初始化TaobaoTmcUserPermitRequest对象
func NewTaobaoTmcUserPermitRequest() *TaobaoTmcUserPermitRequest{
    return &TaobaoTmcUserPermitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserPermitRequest) GetApiMethodName() string {
    return "taobao.tmc.user.permit"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserPermitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Topics Setter
// 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
func (r *TaobaoTmcUserPermitRequest) SetTopics(topics []string) error {
    r.topics = topics
    r.Set("topics", topics)
    return nil
}

// Topics Getter
func (r TaobaoTmcUserPermitRequest) GetTopics() []string {
    return r.topics
}
