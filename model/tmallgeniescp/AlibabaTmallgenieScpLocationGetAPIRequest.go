package tmallgeniescp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpLocationGetAPIRequest 2-IBP查询CDC和RDC数据接口 API请求
// alibaba.tmallgenie.scp.location.get
//
// 天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口
type AlibabaTmallgenieScpLocationGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpLocationGetRequest 初始化AlibabaTmallgenieScpLocationGetAPIRequest对象
func NewAlibabaTmallgenieScpLocationGetRequest() *AlibabaTmallgenieScpLocationGetAPIRequest {
	return &AlibabaTmallgenieScpLocationGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTmallgenieScpLocationGetAPIRequest) Reset() {
	r._requestExtendJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpLocationGetAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.location.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpLocationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTmallgenieScpLocationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpLocationGetAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabaTmallgenieScpLocationGetAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}

var poolAlibabaTmallgenieScpLocationGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTmallgenieScpLocationGetRequest()
	},
}

// GetAlibabaTmallgenieScpLocationGetRequest 从 sync.Pool 获取 AlibabaTmallgenieScpLocationGetAPIRequest
func GetAlibabaTmallgenieScpLocationGetAPIRequest() *AlibabaTmallgenieScpLocationGetAPIRequest {
	return poolAlibabaTmallgenieScpLocationGetAPIRequest.Get().(*AlibabaTmallgenieScpLocationGetAPIRequest)
}

// ReleaseAlibabaTmallgenieScpLocationGetAPIRequest 将 AlibabaTmallgenieScpLocationGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaTmallgenieScpLocationGetAPIRequest(v *AlibabaTmallgenieScpLocationGetAPIRequest) {
	v.Reset()
	poolAlibabaTmallgenieScpLocationGetAPIRequest.Put(v)
}
