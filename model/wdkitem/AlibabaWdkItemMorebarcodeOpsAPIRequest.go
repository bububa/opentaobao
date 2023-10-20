package wdkitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemmorebarcodeopsAPIRequest 商品一品多码维护操作 API请求
// alibaba.wdk.item.morebarcode.ops
//
// 商品一品多码维护操作
type AlibabawdkitemmorebarcodeopsAPIRequest struct {
	model.Params
	// bean
	_updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean
}

// NewAlibabawdkitemmorebarcodeopsRequest 初始化AlibabawdkitemmorebarcodeopsAPIRequest对象
func NewAlibabawdkitemmorebarcodeopsRequest() *AlibabawdkitemmorebarcodeopsAPIRequest {
	return &AlibabawdkitemmorebarcodeopsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemmorebarcodeopsAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.morebarcode.ops"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemmorebarcodeopsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemmorebarcodeopsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUpdateMoreBarCodeRequestBean is UpdateMoreBarCodeRequestBean Setter
// bean
func (r *AlibabawdkitemmorebarcodeopsAPIRequest) SetUpdateMoreBarCodeRequestBean(_updateMoreBarCodeRequestBean *UpdateMoreBarCodeRequestBean) error {
	r._updateMoreBarCodeRequestBean = _updateMoreBarCodeRequestBean
	r.Set("update_more_bar_code_request_bean", _updateMoreBarCodeRequestBean)
	return nil
}

// GetUpdateMoreBarCodeRequestBean UpdateMoreBarCodeRequestBean Getter
func (r AlibabawdkitemmorebarcodeopsAPIRequest) GetUpdateMoreBarCodeRequestBean() *UpdateMoreBarCodeRequestBean {
	return r._updateMoreBarCodeRequestBean
}
