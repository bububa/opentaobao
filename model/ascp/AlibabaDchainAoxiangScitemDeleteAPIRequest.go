package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangScitemDeleteAPIRequest 货品删除 API请求
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
type AlibabaDchainAoxiangScitemDeleteAPIRequest struct {
	model.Params
	// 入参结构体
	_deleteScItemRequest *DeleteScItemRequest
}

// NewAlibabaDchainAoxiangScitemDeleteRequest 初始化AlibabaDchainAoxiangScitemDeleteAPIRequest对象
func NewAlibabaDchainAoxiangScitemDeleteRequest() *AlibabaDchainAoxiangScitemDeleteAPIRequest {
	return &AlibabaDchainAoxiangScitemDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeleteScItemRequest is DeleteScItemRequest Setter
// 入参结构体
func (r *AlibabaDchainAoxiangScitemDeleteAPIRequest) SetDeleteScItemRequest(_deleteScItemRequest *DeleteScItemRequest) error {
	r._deleteScItemRequest = _deleteScItemRequest
	r.Set("delete_sc_item_request", _deleteScItemRequest)
	return nil
}

// GetDeleteScItemRequest DeleteScItemRequest Getter
func (r AlibabaDchainAoxiangScitemDeleteAPIRequest) GetDeleteScItemRequest() *DeleteScItemRequest {
	return r._deleteScItemRequest
}
