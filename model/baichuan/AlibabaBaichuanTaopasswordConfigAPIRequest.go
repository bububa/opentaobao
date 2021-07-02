package baichuan

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordConfigAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordConfigAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
