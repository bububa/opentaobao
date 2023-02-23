package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasServiceGetServTimesAPIRequest 查询某个用户图片空间的使用情况 API请求
// taobao.vas.service.getServTimes
//
// 查询某个用户图片空间的使用情况
type TaobaoVasServiceGetServTimesAPIRequest struct {
	model.Params
	// 服务编码
	_servCode string
}

// NewTaobaoVasServiceGetServTimesRequest 初始化TaobaoVasServiceGetServTimesAPIRequest对象
func NewTaobaoVasServiceGetServTimesRequest() *TaobaoVasServiceGetServTimesAPIRequest {
	return &TaobaoVasServiceGetServTimesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVasServiceGetServTimesAPIRequest) GetApiMethodName() string {
	return "taobao.vas.service.getServTimes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVasServiceGetServTimesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVasServiceGetServTimesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServCode is ServCode Setter
// 服务编码
func (r *TaobaoVasServiceGetServTimesAPIRequest) SetServCode(_servCode string) error {
	r._servCode = _servCode
	r.Set("serv_code", _servCode)
	return nil
}

// GetServCode ServCode Getter
func (r TaobaoVasServiceGetServTimesAPIRequest) GetServCode() string {
	return r._servCode
}
