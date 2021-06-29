package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序消息推送 API请求
alibaba.lsy.miniapp.msg.push

零售云小程序消息推送，推送消息至零售云（喵零等）
*/
type AlibabaLsyMiniappMsgPushRequest struct {
    model.Params
    // 小程序ID
    appId   string
    // 消息ID
    msgId   int64
    // 摊位ID
    storeId   int64
    // 消息模板，miaoling_msg_isv_clue - 线索通知消息
    templateId   string
    // 消息参数
    params   string
}

// 初始化AlibabaLsyMiniappMsgPushRequest对象
func NewAlibabaLsyMiniappMsgPushRequest() *AlibabaLsyMiniappMsgPushRequest{
    return &AlibabaLsyMiniappMsgPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyMiniappMsgPushRequest) GetApiMethodName() string {
    return "alibaba.lsy.miniapp.msg.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyMiniappMsgPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppId Setter
// 小程序ID
func (r *AlibabaLsyMiniappMsgPushRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r AlibabaLsyMiniappMsgPushRequest) GetAppId() string {
    return r.appId
}
// MsgId Setter
// 消息ID
func (r *AlibabaLsyMiniappMsgPushRequest) SetMsgId(msgId int64) error {
    r.msgId = msgId
    r.Set("msg_id", msgId)
    return nil
}

// MsgId Getter
func (r AlibabaLsyMiniappMsgPushRequest) GetMsgId() int64 {
    return r.msgId
}
// StoreId Setter
// 摊位ID
func (r *AlibabaLsyMiniappMsgPushRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaLsyMiniappMsgPushRequest) GetStoreId() int64 {
    return r.storeId
}
// TemplateId Setter
// 消息模板，miaoling_msg_isv_clue - 线索通知消息
func (r *AlibabaLsyMiniappMsgPushRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

// TemplateId Getter
func (r AlibabaLsyMiniappMsgPushRequest) GetTemplateId() string {
    return r.templateId
}
// Params Setter
// 消息参数
func (r *AlibabaLsyMiniappMsgPushRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

// Params Getter
func (r AlibabaLsyMiniappMsgPushRequest) GetParams() string {
    return r.params
}
