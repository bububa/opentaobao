package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscServicecenterServicestoreQueryAPIRequest 根据天猫id查询门店信息 API请求
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
type AlibabaSscServicecenterServicestoreQueryAPIRequest struct {
	model.Params
	// 天猫id
	_id int64
}

// NewAlibabaSscServicecenterServicestoreQueryRequest 初始化AlibabaSscServicecenterServicestoreQueryAPIRequest对象
func NewAlibabaSscServicecenterServicestoreQueryRequest() *AlibabaSscServicecenterServicestoreQueryAPIRequest {
	return &AlibabaSscServicecenterServicestoreQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.servicecenter.servicestore.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 天猫id
func (r *AlibabaSscServicecenterServicestoreQueryAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaSscServicecenterServicestoreQueryAPIRequest) GetId() int64 {
	return r._id
}
