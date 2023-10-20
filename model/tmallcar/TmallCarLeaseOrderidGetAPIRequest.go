package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseorderidgetAPIRequest 天猫开新车查询订单id API请求
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
type TmallcarleaseorderidgetAPIRequest struct {
	model.Params
	// openid
	_openId string
}

// NewTmallcarleaseorderidgetRequest 初始化TmallcarleaseorderidgetAPIRequest对象
func NewTmallcarleaseorderidgetRequest() *TmallcarleaseorderidgetAPIRequest {
	return &TmallcarleaseorderidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseorderidgetAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.orderid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseorderidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseorderidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenId is OpenId Setter
// openid
func (r *TmallcarleaseorderidgetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TmallcarleaseorderidgetAPIRequest) GetOpenId() string {
	return r._openId
}
