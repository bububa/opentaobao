package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascppurchasepricecreateAPIRequest ascp采购价写入接口 API请求
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
type AlibabaascppurchasepricecreateAPIRequest struct {
	model.Params
	// 采购价创建/更新请求
	_createRequest *AlibabaascppurchasepricecreateRequest
}

// NewAlibabaascppurchasepricecreateRequest 初始化AlibabaascppurchasepricecreateAPIRequest对象
func NewAlibabaascppurchasepricecreateRequest() *AlibabaascppurchasepricecreateAPIRequest {
	return &AlibabaascppurchasepricecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascppurchasepricecreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.purchase.price.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascppurchasepricecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascppurchasepricecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 采购价创建/更新请求
func (r *AlibabaascppurchasepricecreateAPIRequest) SetCreateRequest(_createRequest *AlibabaascppurchasepricecreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r AlibabaascppurchasepricecreateAPIRequest) GetCreateRequest() *AlibabaascppurchasepricecreateRequest {
	return r._createRequest
}
