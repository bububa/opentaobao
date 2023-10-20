package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuantaopasswordconfigAPIRequest 淘口令配置数据 API请求
// alibaba.baichuan.taopassword.config
//
// 百川淘口令规则配置接口
type AlibababaichuantaopasswordconfigAPIRequest struct {
	model.Params
}

// NewAlibababaichuantaopasswordconfigRequest 初始化AlibababaichuantaopasswordconfigAPIRequest对象
func NewAlibababaichuantaopasswordconfigRequest() *AlibababaichuantaopasswordconfigAPIRequest {
	return &AlibababaichuantaopasswordconfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuantaopasswordconfigAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuantaopasswordconfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuantaopasswordconfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}
