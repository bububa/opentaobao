package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序消息推送 APIRequest
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

func NewAlibabaLsyMiniappMsgPushRequest() *AlibabaLsyMiniappMsgPushRequest{
    return &AlibabaLsyMiniappMsgPushRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLsyMiniappMsgPushRequest) GetApiMethodName() string {
    return "alibaba.lsy.miniapp.msg.push"
}

func (r AlibabaLsyMiniappMsgPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLsyMiniappMsgPushRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

func (r AlibabaLsyMiniappMsgPushRequest) GetAppId() string {
    return r.appId
}

func (r *AlibabaLsyMiniappMsgPushRequest) SetMsgId(msgId int64) error {
    r.msgId = msgId
    r.Set("msg_id", msgId)
    return nil
}

func (r AlibabaLsyMiniappMsgPushRequest) GetMsgId() int64 {
    return r.msgId
}

func (r *AlibabaLsyMiniappMsgPushRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaLsyMiniappMsgPushRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *AlibabaLsyMiniappMsgPushRequest) SetTemplateId(templateId string) error {
    r.templateId = templateId
    r.Set("template_id", templateId)
    return nil
}

func (r AlibabaLsyMiniappMsgPushRequest) GetTemplateId() string {
    return r.templateId
}

func (r *AlibabaLsyMiniappMsgPushRequest) SetParams(params string) error {
    r.params = params
    r.Set("params", params)
    return nil
}

func (r AlibabaLsyMiniappMsgPushRequest) GetParams() string {
    return r.params
}

