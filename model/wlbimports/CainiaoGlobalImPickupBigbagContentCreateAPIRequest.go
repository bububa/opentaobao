package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupbigbagcontentcreateAPIRequest 大包创建 API请求
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
type CainiaoglobalimpickupbigbagcontentcreateAPIRequest struct {
	model.Params
	// 大包创建入参
	_bigbagCreateRequest *BigbagCreateRequest
}

// NewCainiaoglobalimpickupbigbagcontentcreateRequest 初始化CainiaoglobalimpickupbigbagcontentcreateAPIRequest对象
func NewCainiaoglobalimpickupbigbagcontentcreateRequest() *CainiaoglobalimpickupbigbagcontentcreateAPIRequest {
	return &CainiaoglobalimpickupbigbagcontentcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoglobalimpickupbigbagcontentcreateAPIRequest) GetApiMethodName() string {
	return "cainiao.global.im.pickup.bigbag.content.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoglobalimpickupbigbagcontentcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoglobalimpickupbigbagcontentcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBigbagCreateRequest is BigbagCreateRequest Setter
// 大包创建入参
func (r *CainiaoglobalimpickupbigbagcontentcreateAPIRequest) SetBigbagCreateRequest(_bigbagCreateRequest *BigbagCreateRequest) error {
	r._bigbagCreateRequest = _bigbagCreateRequest
	r.Set("bigbag_create_request", _bigbagCreateRequest)
	return nil
}

// GetBigbagCreateRequest BigbagCreateRequest Getter
func (r CainiaoglobalimpickupbigbagcontentcreateAPIRequest) GetBigbagCreateRequest() *BigbagCreateRequest {
	return r._bigbagCreateRequest
}
