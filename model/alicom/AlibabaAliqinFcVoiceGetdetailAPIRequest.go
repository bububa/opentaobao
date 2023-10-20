package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcvoicegetdetailAPIRequest 获取呼叫详情 API请求
// alibaba.aliqin.fc.voice.getdetail
//
// 通过呼叫id获取呼叫相关的数据
type AlibabaaliqinfcvoicegetdetailAPIRequest struct {
	model.Params
	// 呼叫唯一ID
	_callId string
	// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
	_queryDate int64
	// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
	_prodId int64
}

// NewAlibabaaliqinfcvoicegetdetailRequest 初始化AlibabaaliqinfcvoicegetdetailAPIRequest对象
func NewAlibabaaliqinfcvoicegetdetailRequest() *AlibabaaliqinfcvoicegetdetailAPIRequest {
	return &AlibabaaliqinfcvoicegetdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfcvoicegetdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.voice.getdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfcvoicegetdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfcvoicegetdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallId is CallId Setter
// 呼叫唯一ID
func (r *AlibabaaliqinfcvoicegetdetailAPIRequest) SetCallId(_callId string) error {
	r._callId = _callId
	r.Set("call_id", _callId)
	return nil
}

// GetCallId CallId Getter
func (r AlibabaaliqinfcvoicegetdetailAPIRequest) GetCallId() string {
	return r._callId
}

// SetQueryDate is QueryDate Setter
// Unix时间戳，会查询这个时间点对应那一天的记录（单位毫秒）
func (r *AlibabaaliqinfcvoicegetdetailAPIRequest) SetQueryDate(_queryDate int64) error {
	r._queryDate = _queryDate
	r.Set("query_date", _queryDate)
	return nil
}

// GetQueryDate QueryDate Getter
func (r AlibabaaliqinfcvoicegetdetailAPIRequest) GetQueryDate() int64 {
	return r._queryDate
}

// SetProdId is ProdId Setter
// 语音通知为:11000000300006, 语音验证码为:11010000138001, IVR为:11000000300005, 点击拨号为:11000000300004, SIP为:11000000300009
func (r *AlibabaaliqinfcvoicegetdetailAPIRequest) SetProdId(_prodId int64) error {
	r._prodId = _prodId
	r.Set("prod_id", _prodId)
	return nil
}

// GetProdId ProdId Getter
func (r AlibabaaliqinfcvoicegetdetailAPIRequest) GetProdId() int64 {
	return r._prodId
}
