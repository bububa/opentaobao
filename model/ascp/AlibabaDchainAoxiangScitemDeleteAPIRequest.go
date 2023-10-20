package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitemdeleteAPIRequest 货品删除 API请求
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
type AlibabadchainaoxiangscitemdeleteAPIRequest struct {
	model.Params
	// 入参结构体
	_deleteScItemRequest *DeleteScItemRequest
}

// NewAlibabadchainaoxiangscitemdeleteRequest 初始化AlibabadchainaoxiangscitemdeleteAPIRequest对象
func NewAlibabadchainaoxiangscitemdeleteRequest() *AlibabadchainaoxiangscitemdeleteAPIRequest {
	return &AlibabadchainaoxiangscitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangscitemdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.scitem.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangscitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangscitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteScItemRequest is DeleteScItemRequest Setter
// 入参结构体
func (r *AlibabadchainaoxiangscitemdeleteAPIRequest) SetDeleteScItemRequest(_deleteScItemRequest *DeleteScItemRequest) error {
	r._deleteScItemRequest = _deleteScItemRequest
	r.Set("delete_sc_item_request", _deleteScItemRequest)
	return nil
}

// GetDeleteScItemRequest DeleteScItemRequest Getter
func (r AlibabadchainaoxiangscitemdeleteAPIRequest) GetDeleteScItemRequest() *DeleteScItemRequest {
	return r._deleteScItemRequest
}
