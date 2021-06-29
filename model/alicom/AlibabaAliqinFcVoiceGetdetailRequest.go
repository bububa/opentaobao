package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取呼叫详情 API请求
alibaba.aliqin.fc.voice.getdetail

通过呼叫id获取呼叫相关的数据
*/
type AlibabaAliqinFcVoiceGetdetailRequest struct {
    model.Params
    // 呼叫唯一ID
    callId   string
    // 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
    prodId   int64
    // Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
    queryDate   int64
}

// 初始化AlibabaAliqinFcVoiceGetdetailRequest对象
func NewAlibabaAliqinFcVoiceGetdetailRequest() *AlibabaAliqinFcVoiceGetdetailRequest{
    return &AlibabaAliqinFcVoiceGetdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceGetdetailRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.voice.getdetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceGetdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallId Setter
// 呼叫唯一ID
func (r *AlibabaAliqinFcVoiceGetdetailRequest) SetCallId(callId string) error {
    r.callId = callId
    r.Set("call_id", callId)
    return nil
}

// CallId Getter
func (r AlibabaAliqinFcVoiceGetdetailRequest) GetCallId() string {
    return r.callId
}
// ProdId Setter
// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
func (r *AlibabaAliqinFcVoiceGetdetailRequest) SetProdId(prodId int64) error {
    r.prodId = prodId
    r.Set("prod_id", prodId)
    return nil
}

// ProdId Getter
func (r AlibabaAliqinFcVoiceGetdetailRequest) GetProdId() int64 {
    return r.prodId
}
// QueryDate Setter
// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
func (r *AlibabaAliqinFcVoiceGetdetailRequest) SetQueryDate(queryDate int64) error {
    r.queryDate = queryDate
    r.Set("query_date", queryDate)
    return nil
}

// QueryDate Getter
func (r AlibabaAliqinFcVoiceGetdetailRequest) GetQueryDate() int64 {
    return r.queryDate
}
