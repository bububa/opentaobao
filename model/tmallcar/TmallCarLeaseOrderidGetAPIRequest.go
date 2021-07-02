package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseOrderidGetAPIRequest 天猫开新车查询订单id API请求
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
type TmallCarLeaseOrderidGetAPIRequest struct {
	model.Params
	// openid
	_openId string
}

// NewTmallCarLeaseOrderidGetRequest 初始化TmallCarLeaseOrderidGetAPIRequest对象
func NewTmallCarLeaseOrderidGetRequest() *TmallCarLeaseOrderidGetAPIRequest {
	return &TmallCarLeaseOrderidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseOrderidGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.orderid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseOrderidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOpenId is OpenId Setter
// openid
func (r *TmallCarLeaseOrderidGetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TmallCarLeaseOrderidGetAPIRequest) GetOpenId() string {
	return r._openId
}
