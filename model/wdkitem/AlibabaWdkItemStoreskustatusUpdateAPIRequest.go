package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskustatusUpdateAPIRequest 修改门店商品状态 API请求
// alibaba.wdk.item.storeskustatus.update
//
// 五道口商品 修改门店商品状态
type AlibabaWdkItemStoreskustatusUpdateAPIRequest struct {
	model.Params
	// bean
	_bean *UpdateStoreSkuLifeStatusRequestBean
}

// NewAlibabaWdkItemStoreskustatusUpdateRequest 初始化AlibabaWdkItemStoreskustatusUpdateAPIRequest对象
func NewAlibabaWdkItemStoreskustatusUpdateRequest() *AlibabaWdkItemStoreskustatusUpdateAPIRequest {
	return &AlibabaWdkItemStoreskustatusUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemStoreskustatusUpdateAPIRequest) Reset() {
	r._bean = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storeskustatus.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBean is Bean Setter
// bean
func (r *AlibabaWdkItemStoreskustatusUpdateAPIRequest) SetBean(_bean *UpdateStoreSkuLifeStatusRequestBean) error {
	r._bean = _bean
	r.Set("bean", _bean)
	return nil
}

// GetBean Bean Getter
func (r AlibabaWdkItemStoreskustatusUpdateAPIRequest) GetBean() *UpdateStoreSkuLifeStatusRequestBean {
	return r._bean
}

var poolAlibabaWdkItemStoreskustatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemStoreskustatusUpdateRequest()
	},
}

// GetAlibabaWdkItemStoreskustatusUpdateRequest 从 sync.Pool 获取 AlibabaWdkItemStoreskustatusUpdateAPIRequest
func GetAlibabaWdkItemStoreskustatusUpdateAPIRequest() *AlibabaWdkItemStoreskustatusUpdateAPIRequest {
	return poolAlibabaWdkItemStoreskustatusUpdateAPIRequest.Get().(*AlibabaWdkItemStoreskustatusUpdateAPIRequest)
}

// ReleaseAlibabaWdkItemStoreskustatusUpdateAPIRequest 将 AlibabaWdkItemStoreskustatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemStoreskustatusUpdateAPIRequest(v *AlibabaWdkItemStoreskustatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemStoreskustatusUpdateAPIRequest.Put(v)
}
