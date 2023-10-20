package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupstoresgetAPIRequest 首公里揽收-集货仓列表查询 API请求
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
type CainiaoglobalimpickupstoresgetAPIRequest struct {
	model.Params
	// 请求体
	_transferstoreQueryRequest *TransferstoreQueryRequest
}

// NewCainiaoglobalimpickupstoresgetRequest 初始化CainiaoglobalimpickupstoresgetAPIRequest对象
func NewCainiaoglobalimpickupstoresgetRequest() *CainiaoglobalimpickupstoresgetAPIRequest {
	return &CainiaoglobalimpickupstoresgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupstoresgetAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.stores.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupstoresgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupstoresgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTransferstoreQueryRequest is TransferstoreQueryRequest Setter
// 请求体
func (r *CainiaoglobalimpickupstoresgetAPIRequest) SetTransferstoreQueryRequest(_transferstoreQueryRequest *TransferstoreQueryRequest) error {
	r._transferstoreQueryRequest = _transferstoreQueryRequest
	r.Set("transferstore_query_request", _transferstoreQueryRequest)
	return nil
}

// GetTransferstoreQueryRequest TransferstoreQueryRequest Getter
func (r CainiaoglobalimpickupstoresgetAPIRequest) GetTransferstoreQueryRequest() *TransferstoreQueryRequest {
	return r._transferstoreQueryRequest
}
