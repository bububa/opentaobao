package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMixnickWetoplayAPIRequest 微淘混淆nick转互动混淆nick API请求
// taobao.mixnick.wetoplay
//
// 微淘应用的混淆nick转为互动类型混淆nick
type TaobaoMixnickWetoplayAPIRequest struct {
	model.Params
	// 排查问题id
	_traceId string
	// 微淘混淆nick
	_weMixnick string
}

// NewTaobaoMixnickWetoplayRequest 初始化TaobaoMixnickWetoplayAPIRequest对象
func NewTaobaoMixnickWetoplayRequest() *TaobaoMixnickWetoplayAPIRequest {
	return &TaobaoMixnickWetoplayAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMixnickWetoplayAPIRequest) Reset() {
	r._traceId = ""
	r._weMixnick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMixnickWetoplayAPIRequest) GetApiMethodName() string {
	return "taobao.mixnick.wetoplay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMixnickWetoplayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMixnickWetoplayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 排查问题id
func (r *TaobaoMixnickWetoplayAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoMixnickWetoplayAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetWeMixnick is WeMixnick Setter
// 微淘混淆nick
func (r *TaobaoMixnickWetoplayAPIRequest) SetWeMixnick(_weMixnick string) error {
	r._weMixnick = _weMixnick
	r.Set("we_mixnick", _weMixnick)
	return nil
}

// GetWeMixnick WeMixnick Getter
func (r TaobaoMixnickWetoplayAPIRequest) GetWeMixnick() string {
	return r._weMixnick
}

var poolTaobaoMixnickWetoplayAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMixnickWetoplayRequest()
	},
}

// GetTaobaoMixnickWetoplayRequest 从 sync.Pool 获取 TaobaoMixnickWetoplayAPIRequest
func GetTaobaoMixnickWetoplayAPIRequest() *TaobaoMixnickWetoplayAPIRequest {
	return poolTaobaoMixnickWetoplayAPIRequest.Get().(*TaobaoMixnickWetoplayAPIRequest)
}

// ReleaseTaobaoMixnickWetoplayAPIRequest 将 TaobaoMixnickWetoplayAPIRequest 放入 sync.Pool
func ReleaseTaobaoMixnickWetoplayAPIRequest(v *TaobaoMixnickWetoplayAPIRequest) {
	v.Reset()
	poolTaobaoMixnickWetoplayAPIRequest.Put(v)
}
