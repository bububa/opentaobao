package alicom

import (
	"net/url"
	"sync"

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
	// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
	_queryDate int64
	// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
	_prodId int64
}

// NewAlibabaAliqinFcVoiceGetdetailRequest 初始化AlibabaAliqinFcVoiceGetdetailAPIRequest对象
func NewAlibabaAliqinFcVoiceGetdetailRequest() *AlibabaAliqinFcVoiceGetdetailAPIRequest {
	return &AlibabaAliqinFcVoiceGetdetailAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcVoiceGetdetailAPIRequest) Reset() {
	r._callId = ""
	r._queryDate = 0
	r._prodId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcVoiceGetdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAliqinFcVoiceGetdetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcVoiceGetdetailRequest()
	},
}

// GetAlibabaAliqinFcVoiceGetdetailRequest 从 sync.Pool 获取 AlibabaAliqinFcVoiceGetdetailAPIRequest
func GetAlibabaAliqinFcVoiceGetdetailAPIRequest() *AlibabaAliqinFcVoiceGetdetailAPIRequest {
	return poolAlibabaAliqinFcVoiceGetdetailAPIRequest.Get().(*AlibabaAliqinFcVoiceGetdetailAPIRequest)
}

// ReleaseAlibabaAliqinFcVoiceGetdetailAPIRequest 将 AlibabaAliqinFcVoiceGetdetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcVoiceGetdetailAPIRequest(v *AlibabaAliqinFcVoiceGetdetailAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcVoiceGetdetailAPIRequest.Put(v)
}
