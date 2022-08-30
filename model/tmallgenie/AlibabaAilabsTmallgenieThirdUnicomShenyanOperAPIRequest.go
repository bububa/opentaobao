package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest 联通神眼注册操作 API请求
// alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper
//
// 联通神眼注册操作
type AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest struct {
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

// NewAlibabaAilabsTmallgenieThirdUnicomShenyanOperRequest 初始化AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest对象
func NewAlibabaAilabsTmallgenieThirdUnicomShenyanOperRequest() *AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest {
	return &AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCuei is Cuei Setter
// cuei
func (r *AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) SetCuei(_cuei string) error {
	r._cuei = _cuei
	r.Set("cuei", _cuei)
	return nil
}

// GetCuei Cuei Getter
func (r AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) GetCuei() string {
	return r._cuei
}

// SetUid is Uid Setter
// 联通用户id
func (r *AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) GetUid() string {
	return r._uid
}

// SetTraceId is TraceId Setter
// traceId
func (r *AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetType is Type Setter
// 操作类型：1.打开神眼 2.关闭神眼
func (r *AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIRequest) GetType() string {
	return r._type
}
