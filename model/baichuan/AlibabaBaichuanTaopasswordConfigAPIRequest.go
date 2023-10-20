package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordConfigAPIRequest 淘口令配置数据 API请求
// alibaba.baichuan.taopassword.config
//
// 百川淘口令规则配置接口
type AlibabaBaichuanTaopasswordConfigAPIRequest struct {
	model.Params
}

// NewAlibabaBaichuanTaopasswordConfigRequest 初始化AlibabaBaichuanTaopasswordConfigAPIRequest对象
func NewAlibabaBaichuanTaopasswordConfigRequest() *AlibabaBaichuanTaopasswordConfigAPIRequest {
	return &AlibabaBaichuanTaopasswordConfigAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBaichuanTaopasswordConfigAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordConfigAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordConfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaichuanTaopasswordConfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaBaichuanTaopasswordConfigAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBaichuanTaopasswordConfigRequest()
	},
}

// GetAlibabaBaichuanTaopasswordConfigRequest 从 sync.Pool 获取 AlibabaBaichuanTaopasswordConfigAPIRequest
func GetAlibabaBaichuanTaopasswordConfigAPIRequest() *AlibabaBaichuanTaopasswordConfigAPIRequest {
	return poolAlibabaBaichuanTaopasswordConfigAPIRequest.Get().(*AlibabaBaichuanTaopasswordConfigAPIRequest)
}

// ReleaseAlibabaBaichuanTaopasswordConfigAPIRequest 将 AlibabaBaichuanTaopasswordConfigAPIRequest 放入 sync.Pool
func ReleaseAlibabaBaichuanTaopasswordConfigAPIRequest(v *AlibabaBaichuanTaopasswordConfigAPIRequest) {
	v.Reset()
	poolAlibabaBaichuanTaopasswordConfigAPIRequest.Put(v)
}
