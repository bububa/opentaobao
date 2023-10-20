package wdkitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMorebarcodeOpsAPIRequest 商品一品多码维护操作 API请求
// alibaba.wdk.item.morebarcode.ops
//
// 商品一品多码维护操作
type AlibabaWdkItemMorebarcodeOpsAPIRequest struct {
	model.Params
	// bean
	_updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean
}

// NewAlibabaWdkItemMorebarcodeOpsRequest 初始化AlibabaWdkItemMorebarcodeOpsAPIRequest对象
func NewAlibabaWdkItemMorebarcodeOpsRequest() *AlibabaWdkItemMorebarcodeOpsAPIRequest {
	return &AlibabaWdkItemMorebarcodeOpsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkItemMorebarcodeOpsAPIRequest) Reset() {
	r._updateMoreBarCodeRequestBean = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.morebarcode.ops"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateMoreBarCodeRequestBean is UpdateMoreBarCodeRequestBean Setter
// bean
func (r *AlibabaWdkItemMorebarcodeOpsAPIRequest) SetUpdateMoreBarCodeRequestBean(_updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean) error {
	r._updateMoreBarCodeRequestBean = _updateMoreBarCodeRequestBean
	r.Set("update_more_bar_code_request_bean", _updateMoreBarCodeRequestBean)
	return nil
}

// GetUpdateMoreBarCodeRequestBean UpdateMoreBarCodeRequestBean Getter
func (r AlibabaWdkItemMorebarcodeOpsAPIRequest) GetUpdateMoreBarCodeRequestBean() *UpdateMoreBarCodeRequestBean {
	return r._updateMoreBarCodeRequestBean
}

var poolAlibabaWdkItemMorebarcodeOpsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkItemMorebarcodeOpsRequest()
	},
}

// GetAlibabaWdkItemMorebarcodeOpsRequest 从 sync.Pool 获取 AlibabaWdkItemMorebarcodeOpsAPIRequest
func GetAlibabaWdkItemMorebarcodeOpsAPIRequest() *AlibabaWdkItemMorebarcodeOpsAPIRequest {
	return poolAlibabaWdkItemMorebarcodeOpsAPIRequest.Get().(*AlibabaWdkItemMorebarcodeOpsAPIRequest)
}

// ReleaseAlibabaWdkItemMorebarcodeOpsAPIRequest 将 AlibabaWdkItemMorebarcodeOpsAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkItemMorebarcodeOpsAPIRequest(v *AlibabaWdkItemMorebarcodeOpsAPIRequest) {
	v.Reset()
	poolAlibabaWdkItemMorebarcodeOpsAPIRequest.Put(v)
}
