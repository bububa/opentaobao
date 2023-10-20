package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeleteitemAPIRequest 大麦换验平台-第三方对外开放-票品接口deleteItem API请求
// alibaba.damai.mev.open.deleteitem
//
// deleteItem
type AlibabadamaimevopendeleteitemAPIRequest struct {
	model.Params
	// 入参deleteItemParam
	_deleteItemParam *TicketItemIdOpenParam
}

// NewAlibabadamaimevopendeleteitemRequest 初始化AlibabadamaimevopendeleteitemAPIRequest对象
func NewAlibabadamaimevopendeleteitemRequest() *AlibabadamaimevopendeleteitemAPIRequest {
	return &AlibabadamaimevopendeleteitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopendeleteitemAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopendeleteitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopendeleteitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteItemParam is DeleteItemParam Setter
// 入参deleteItemParam
func (r *AlibabadamaimevopendeleteitemAPIRequest) SetDeleteItemParam(_deleteItemParam *TicketItemIdOpenParam) error {
	r._deleteItemParam = _deleteItemParam
	r.Set("delete_item_param", _deleteItemParam)
	return nil
}

// GetDeleteItemParam DeleteItemParam Getter
func (r AlibabadamaimevopendeleteitemAPIRequest) GetDeleteItemParam() *TicketItemIdOpenParam {
	return r._deleteItemParam
}
