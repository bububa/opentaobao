package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiIdGetAPIRequest 获取请求id API请求
// alibaba.happytrip.taxi.id.get
//
// 获取订单号
type AlibabaHappytripTaxiIdGetAPIRequest struct {
	model.Params
	// 用户唯一标识
	_uid string
}

// NewAlibabaHappytripTaxiIdGetRequest 初始化AlibabaHappytripTaxiIdGetAPIRequest对象
func NewAlibabaHappytripTaxiIdGetRequest() *AlibabaHappytripTaxiIdGetAPIRequest {
	return &AlibabaHappytripTaxiIdGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiIdGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.id.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiIdGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiIdGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUid is Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiIdGetAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabaHappytripTaxiIdGetAPIRequest) GetUid() string {
	return r._uid
}
