package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcVoiceGetdetailAPIRequest 获取呼叫详情 API请求
// alibaba.aliqin.fc.voice.getdetail
//
// 通过呼叫id获取呼叫相关的数据
type AlibabaAliqinFcVoiceGetdetailAPIRequest struct {
	model.Params
	// 呼叫唯一ID
	_callId string
	// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
	_prodId int64
	// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
	_queryDate int64
}

// NewAlibabaAliqinFcVoiceGetdetailRequest 初始化AlibabaAliqinFcVoiceGetdetailAPIRequest对象
func NewAlibabaAliqinFcVoiceGetdetailRequest() *AlibabaAliqinFcVoiceGetdetailAPIRequest {
	return &AlibabaAliqinFcVoiceGetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCallId is CallId Setter
// 呼叫唯一ID
func (r *AlibabaAliqinFcVoiceGetdetailAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// GetCallId CallId Getter
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetCallId() string {
	return r._callId
}

// SetProdId is ProdId Setter
// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
func (r *AlibabaAliqinFcVoiceGetdetailAPIRequest) SetProdId(_prodId int64) error {
	r._prodId = _prodId
	r.Set("prod_id", _prodId)
	return nil
}

// GetProdId ProdId Getter
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetProdId() int64 {
	return r._prodId
}

// SetQueryDate is QueryDate Setter
// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
func (r *AlibabaAliqinFcVoiceGetdetailAPIRequest) SetQueryDate(_queryDate int64) error {
	r._queryDate = _queryDate
	r.Set("query_date", _queryDate)
	return nil
}

// GetQueryDate QueryDate Getter
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetQueryDate() int64 {
	return r._queryDate
}
