package baichuanctg

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgToutiaoContentAPIRequest 微博输出头条数据 API请求
// alibaba.baichuan.ctg.toutiao.content
//
// 百川头条内容获取
type AlibabaBaichuanCtgToutiaoContentAPIRequest struct {
	model.Params
	// param0
	_param0 *CtgRequest
}

// NewAlibabaBaichuanCtgToutiaoContentRequest 初始化AlibabaBaichuanCtgToutiaoContentAPIRequest对象
func NewAlibabaBaichuanCtgToutiaoContentRequest() *AlibabaBaichuanCtgToutiaoContentAPIRequest {
	return &AlibabaBaichuanCtgToutiaoContentAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBaichuanCtgToutiaoContentAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.toutiao.content"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// param0
func (r *AlibabaBaichuanCtgToutiaoContentAPIRequest) SetParam0(_param0 *CtgRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetParam0() *CtgRequest {
	return r._param0
}

var poolAlibabaBaichuanCtgToutiaoContentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBaichuanCtgToutiaoContentRequest()
	},
}

// GetAlibabaBaichuanCtgToutiaoContentRequest 从 sync.Pool 获取 AlibabaBaichuanCtgToutiaoContentAPIRequest
func GetAlibabaBaichuanCtgToutiaoContentAPIRequest() *AlibabaBaichuanCtgToutiaoContentAPIRequest {
	return poolAlibabaBaichuanCtgToutiaoContentAPIRequest.Get().(*AlibabaBaichuanCtgToutiaoContentAPIRequest)
}

// ReleaseAlibabaBaichuanCtgToutiaoContentAPIRequest 将 AlibabaBaichuanCtgToutiaoContentAPIRequest 放入 sync.Pool
func ReleaseAlibabaBaichuanCtgToutiaoContentAPIRequest(v *AlibabaBaichuanCtgToutiaoContentAPIRequest) {
	v.Reset()
	poolAlibabaBaichuanCtgToutiaoContentAPIRequest.Put(v)
}
