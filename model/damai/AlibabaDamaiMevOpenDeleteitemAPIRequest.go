package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteitemAPIRequest
大麦换验平台-第三方对外开放-票品接口deleteItem API请求
alibaba.damai.mev.open.deleteitem

deleteItem */
type AlibabaDamaiMevOpenDeleteitemAPIRequest struct {
	model.Params
	// 入参deleteItemParam
	_deleteItemParam *TicketItemIdOpenParam
}

// NewAlibabaDamaiMevOpenDeleteitemRequest 初始化AlibabaDamaiMevOpenDeleteitemAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteitemRequest() *AlibabaDamaiMevOpenDeleteitemAPIRequest {
	return &AlibabaDamaiMevOpenDeleteitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteitemAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deleteitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeleteItemParam Setter
// 入参deleteItemParam
func (r *AlibabaDamaiMevOpenDeleteitemAPIRequest) SetDeleteItemParam(_deleteItemParam *TicketItemIdOpenParam) error {
	r._deleteItemParam = _deleteItemParam
	r.Set("delete_item_param", _deleteItemParam)
	return nil
}

// Get DeleteItemParam Getter
func (r AlibabaDamaiMevOpenDeleteitemAPIRequest) GetDeleteItemParam() *TicketItemIdOpenParam {
	return r._deleteItemParam
}
