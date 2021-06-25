package cloudpush

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川云推送发送消息给ios APIRequest
taobao.cloudpush.message.ios

百川云推送发送消息给iOS设备.
*/
type TaobaoCloudpushMessageIosRequest struct {
    model.Params

    // 发送的消息内容.
    body   string 

    // 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
    target   string 

    // 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
    targetValue   string 

}

func NewTaobaoCloudpushMessageIosRequest() *TaobaoCloudpushMessageIosRequest{
    return &TaobaoCloudpushMessageIosRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCloudpushMessageIosRequest) GetApiMethodName() string {
    return "taobao.cloudpush.message.ios"
}

func (r TaobaoCloudpushMessageIosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCloudpushMessageIosRequest) SetBody(body string) error {
    r.body = body
    r.Set("body", body)
    return nil
}

func (r TaobaoCloudpushMessageIosRequest) GetBody() string {
    return r.body
}

func (r *TaobaoCloudpushMessageIosRequest) SetTarget(target string) error {
    r.target = target
    r.Set("target", target)
    return nil
}

func (r TaobaoCloudpushMessageIosRequest) GetTarget() string {
    return r.target
}

func (r *TaobaoCloudpushMessageIosRequest) SetTargetValue(targetValue string) error {
    r.targetValue = targetValue
    r.Set("target_value", targetValue)
    return nil
}

func (r TaobaoCloudpushMessageIosRequest) GetTargetValue() string {
    return r.targetValue
}

