package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitystoregetstorelistAPIRequest ISV查询门店 API请求
// alibaba.lsy.crm.activity.store.getstorelist
//
// ISV查询门店
type AlibabalsycrmactivitystoregetstorelistAPIRequest struct {
	model.Params
	// 系统自动生成
	_queryStoreReq *NrtQueryStoreReq
}

// NewAlibabalsycrmactivitystoregetstorelistRequest 初始化AlibabalsycrmactivitystoregetstorelistAPIRequest对象
func NewAlibabalsycrmactivitystoregetstorelistRequest() *AlibabalsycrmactivitystoregetstorelistAPIRequest {
	return &AlibabalsycrmactivitystoregetstorelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivitystoregetstorelistAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.store.getstorelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivitystoregetstorelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivitystoregetstorelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryStoreReq is QueryStoreReq Setter
// 系统自动生成
func (r *AlibabalsycrmactivitystoregetstorelistAPIRequest) SetQueryStoreReq(_queryStoreReq *NrtQueryStoreReq) error {
	r._queryStoreReq = _queryStoreReq
	r.Set("query_store_req", _queryStoreReq)
	return nil
}

// GetQueryStoreReq QueryStoreReq Getter
func (r AlibabalsycrmactivitystoregetstorelistAPIRequest) GetQueryStoreReq() *NrtQueryStoreReq {
	return r._queryStoreReq
}
