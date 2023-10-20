package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityStoreGetstorelistAPIRequest ISV查询门店 API请求
// alibaba.lsy.crm.activity.store.getstorelist
//
// ISV查询门店
type AlibabaLsyCrmActivityStoreGetstorelistAPIRequest struct {
	model.Params
	// 系统自动生成
	_queryStoreReq *NrtQueryStoreReq
}

// NewAlibabaLsyCrmActivityStoreGetstorelistRequest 初始化AlibabaLsyCrmActivityStoreGetstorelistAPIRequest对象
func NewAlibabaLsyCrmActivityStoreGetstorelistRequest() *AlibabaLsyCrmActivityStoreGetstorelistAPIRequest {
	return &AlibabaLsyCrmActivityStoreGetstorelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityStoreGetstorelistAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.store.getstorelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityStoreGetstorelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmActivityStoreGetstorelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryStoreReq is QueryStoreReq Setter
// 系统自动生成
func (r *AlibabaLsyCrmActivityStoreGetstorelistAPIRequest) SetQueryStoreReq(_queryStoreReq *NrtQueryStoreReq) error {
	r._queryStoreReq = _queryStoreReq
	r.Set("query_store_req", _queryStoreReq)
	return nil
}

// GetQueryStoreReq QueryStoreReq Getter
func (r AlibabaLsyCrmActivityStoreGetstorelistAPIRequest) GetQueryStoreReq() *NrtQueryStoreReq {
	return r._queryStoreReq
}
