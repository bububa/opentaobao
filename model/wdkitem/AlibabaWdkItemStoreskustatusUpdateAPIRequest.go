package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemStoreskustatusUpdateAPIRequest
修改门店商品状态 API请求
alibaba.wdk.item.storeskustatus.update

五道口商品 修改门店商品状态 */
type AlibabaWdkItemStoreskustatusUpdateAPIRequest struct {
	model.Params
	// bean
	_bean *UpdateStoreSkuLifeStatusRequestBean
}

// NewAlibabaWdkItemStoreskustatusUpdateRequest 初始化AlibabaWdkItemStoreskustatusUpdateAPIRequest对象
func NewAlibabaWdkItemStoreskustatusUpdateRequest() *AlibabaWdkItemStoreskustatusUpdateAPIRequest {
	return &AlibabaWdkItemStoreskustatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storeskustatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Bean Setter
// bean
func (r *AlibabaWdkItemStoreskustatusUpdateAPIRequest) SetBean(_bean *UpdateStoreSkuLifeStatusRequestBean) error {
	r._bean = _bean
	r.Set("bean", _bean)
	return nil
}

// Get Bean Getter
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetBean() *UpdateStoreSkuLifeStatusRequestBean {
	return r._bean
}
