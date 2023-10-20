package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscservicecenterservicestorequeryAPIRequest 根据天猫id查询门店信息 API请求
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
type AlibabasscservicecenterservicestorequeryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
}

// NewAlibabasscservicecenterservicestorequeryRequest 初始化AlibabasscservicecenterservicestorequeryAPIRequest对象
func NewAlibabasscservicecenterservicestorequeryRequest() *AlibabasscservicecenterservicestorequeryAPIRequest {
	return &AlibabasscservicecenterservicestorequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscservicecenterservicestorequeryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.servicecenter.servicestore.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscservicecenterservicestorequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscservicecenterservicestorequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 天猫id
func (r *AlibabasscservicecenterservicestorequeryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabasscservicecenterservicestorequeryAPIRequest) GetId() int64 {
	return r._id
}
