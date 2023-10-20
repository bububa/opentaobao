package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemstoreskustatusupdateAPIRequest 修改门店商品状态 API请求
// alibaba.wdk.item.storeskustatus.update
//
// 五道口商品 修改门店商品状态
type AlibabawdkitemstoreskustatusupdateAPIRequest struct {
	model.Params
	// bean
	_bean *UpdateStoreSkuLifeStatusRequestBean
}

// NewAlibabawdkitemstoreskustatusupdateRequest 初始化AlibabawdkitemstoreskustatusupdateAPIRequest对象
func NewAlibabawdkitemstoreskustatusupdateRequest() *AlibabawdkitemstoreskustatusupdateAPIRequest {
	return &AlibabawdkitemstoreskustatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemstoreskustatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storeskustatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemstoreskustatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemstoreskustatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBean is Bean Setter
// bean
func (r *AlibabawdkitemstoreskustatusupdateAPIRequest) SetBean(_bean *UpdateStoreSkuLifeStatusRequestBean) error {
	r._bean = _bean
	r.Set("bean", _bean)
	return nil
}

// GetBean Bean Getter
func (r AlibabawdkitemstoreskustatusupdateAPIRequest) GetBean() *UpdateStoreSkuLifeStatusRequestBean {
	return r._bean
}
