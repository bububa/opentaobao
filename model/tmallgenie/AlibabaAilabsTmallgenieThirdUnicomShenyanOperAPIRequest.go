package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest 联通神眼注册操作 API请求
// alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper
//
// 联通神眼注册操作
type AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest struct {
	model.Params
	// cuei
	_cuei string
	// 联通用户id
	_uid string
	// traceId
	_traceId string
	// 操作类型：1.打开神眼 2.关闭神眼
	_type string
}

// NewAlibabaailabstmallgeniethirdunicomshenyanoperRequest 初始化AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest对象
func NewAlibabaailabstmallgeniethirdunicomshenyanoperRequest() *AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest {
	return &AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCuei is Cuei Setter
// cuei
func (r *AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) SetCuei(_cuei string) error {
	r._cuei = _cuei
	r.Set("cuei", _cuei)
	return nil
}

// GetCuei Cuei Getter
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetCuei() string {
	return r._cuei
}

// SetUid is Uid Setter
// 联通用户id
func (r *AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetUid() string {
	return r._uid
}

// SetTraceId is TraceId Setter
// traceId
func (r *AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetType is Type Setter
// 操作类型：1.打开神眼 2.关闭神眼
func (r *AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaailabstmallgeniethirdunicomshenyanoperAPIRequest) GetType() string {
	return r._type
}
