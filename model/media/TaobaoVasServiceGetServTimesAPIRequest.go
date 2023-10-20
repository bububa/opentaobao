package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovasservicegetServTimesAPIRequest 查询某个用户图片空间的使用情况 API请求
// taobao.vas.service.getServTimes
//
// 查询某个用户图片空间的使用情况
type TaobaovasservicegetServTimesAPIRequest struct {
	model.Params
	// 服务编码
	_servCode string
}

// NewTaobaovasservicegetServTimesRequest 初始化TaobaovasservicegetServTimesAPIRequest对象
func NewTaobaovasservicegetServTimesRequest() *TaobaovasservicegetServTimesAPIRequest {
	return &TaobaovasservicegetServTimesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovasservicegetServTimesAPIRequest) GetApiMethodName() string {
	return "taobao.vas.service.getServTimes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovasservicegetServTimesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovasservicegetServTimesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServCode is ServCode Setter
// 服务编码
func (r *TaobaovasservicegetServTimesAPIRequest) SetServCode(_servCode string) error {
	r._servCode = _servCode
	r.Set("serv_code", _servCode)
	return nil
}

// GetServCode ServCode Getter
func (r TaobaovasservicegetServTimesAPIRequest) GetServCode() string {
	return r._servCode
}
