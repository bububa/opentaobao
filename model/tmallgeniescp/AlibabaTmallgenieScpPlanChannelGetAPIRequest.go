package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanChannelGetAPIRequest 5-IBP同步渠道接口 API请求
// alibaba.tmallgenie.scp.plan.channel.get
//
// IBP同步渠道接口
type AlibabaTmallgenieScpPlanChannelGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanChannelGetRequest 初始化AlibabaTmallgenieScpPlanChannelGetAPIRequest对象
func NewAlibabaTmallgenieScpPlanChannelGetRequest() *AlibabaTmallgenieScpPlanChannelGetAPIRequest {
	return &AlibabaTmallgenieScpPlanChannelGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpPlanChannelGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.channel.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanChannelGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanChannelGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpPlanChannelGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpPlanChannelGetRequest()
	},
}

// GetAlibabaTmallgenieScpPlanChannelGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpPlanChannelGetAPIRequest
func GetAlibabaTmallgenieScpPlanChannelGetAPIRequest() *AlibabaTmallgenieScpPlanChannelGetAPIRequest {
	return poolAlibabaTmallgenieScpPlanChannelGetAPIRequest.Get().(*AlibabaTmallgenieScpPlanChannelGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpPlanChannelGetAPIRequest 将 AlibabaTmallgenieScpPlanChannelGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanChannelGetAPIRequest(v *AlibabaTmallgenieScpPlanChannelGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanChannelGetAPIRequest.Put(v)
}
