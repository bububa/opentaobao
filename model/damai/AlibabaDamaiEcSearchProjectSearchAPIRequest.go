package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiEcSearchProjectSearchAPIRequest 大麦电商对外搜索服务 API请求
// alibaba.damai.ec.search.project.search
//
// 大麦电商对外搜索服务
type AlibabaDamaiEcSearchProjectSearchAPIRequest struct {
	model.Params
	// 入参对象
	_param *TopSearchProjectParam
}

// NewAlibabaDamaiEcSearchProjectSearchRequest 初始化AlibabaDamaiEcSearchProjectSearchAPIRequest对象
func NewAlibabaDamaiEcSearchProjectSearchRequest() *AlibabaDamaiEcSearchProjectSearchAPIRequest {
	return &AlibabaDamaiEcSearchProjectSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiEcSearchProjectSearchAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.ec.search.project.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlibabaDamaiEcSearchProjectSearchAPIRequest) SetParam(_param *TopSearchProjectParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetParam() *TopSearchProjectParam {
	return r._param
}

var poolAlibabaDamaiEcSearchProjectSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiEcSearchProjectSearchRequest()
	},
}

// GetAlibabaDamaiEcSearchProjectSearchRequest 从 sync.Pool 获取 AlibabaDamaiEcSearchProjectSearchAPIRequest
func GetAlibabaDamaiEcSearchProjectSearchAPIRequest() *AlibabaDamaiEcSearchProjectSearchAPIRequest {
	return poolAlibabaDamaiEcSearchProjectSearchAPIRequest.Get().(*AlibabaDamaiEcSearchProjectSearchAPIRequest)
}

// ReleaseAlibabaDamaiEcSearchProjectSearchAPIRequest 将 AlibabaDamaiEcSearchProjectSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiEcSearchProjectSearchAPIRequest(v *AlibabaDamaiEcSearchProjectSearchAPIRequest) {
	v.Reset()
	poolAlibabaDamaiEcSearchProjectSearchAPIRequest.Put(v)
}
