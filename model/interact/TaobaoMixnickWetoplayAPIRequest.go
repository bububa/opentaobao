package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomixnickwetoplayAPIRequest 微淘混淆nick转互动混淆nick API请求
// taobao.mixnick.wetoplay
//
// 微淘应用的混淆nick转为互动类型混淆nick
type TaobaomixnickwetoplayAPIRequest struct {
	model.Params
	// 排查问题id
	_traceId string
	// 微淘混淆nick
	_weMixnick string
}

// NewTaobaomixnickwetoplayRequest 初始化TaobaomixnickwetoplayAPIRequest对象
func NewTaobaomixnickwetoplayRequest() *TaobaomixnickwetoplayAPIRequest {
	return &TaobaomixnickwetoplayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomixnickwetoplayAPIRequest) GetApiMethodName() string {
	return "taobao.mixnick.wetoplay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomixnickwetoplayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomixnickwetoplayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 排查问题id
func (r *TaobaomixnickwetoplayAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaomixnickwetoplayAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetWeMixnick is WeMixnick Setter
// 微淘混淆nick
func (r *TaobaomixnickwetoplayAPIRequest) SetWeMixnick(_weMixnick string) error {
	r._weMixnick = _weMixnick
	r.Set("we_mixnick", _weMixnick)
	return nil
}

// GetWeMixnick WeMixnick Getter
func (r TaobaomixnickwetoplayAPIRequest) GetWeMixnick() string {
	return r._weMixnick
}
