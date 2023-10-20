package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeitemtagsubmitAPIRequest ETC上翻商品打标接口 API请求
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
type AlibabaalihousenewhomeitemtagsubmitAPIRequest struct {
	model.Params
	//  打标结构体
	_itemTagRequestDto *ItemTagRequestDto
}

// NewAlibabaalihousenewhomeitemtagsubmitRequest 初始化AlibabaalihousenewhomeitemtagsubmitAPIRequest对象
func NewAlibabaalihousenewhomeitemtagsubmitRequest() *AlibabaalihousenewhomeitemtagsubmitAPIRequest {
	return &AlibabaalihousenewhomeitemtagsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeitemtagsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.item.tag.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeitemtagsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeitemtagsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemTagRequestDto is ItemTagRequestDto Setter
//
//	打标结构体
func (r *AlibabaalihousenewhomeitemtagsubmitAPIRequest) SetItemTagRequestDto(_itemTagRequestDto *ItemTagRequestDto) error {
	r._itemTagRequestDto = _itemTagRequestDto
	r.Set("item_tag_request_dto", _itemTagRequestDto)
	return nil
}

// GetItemTagRequestDto ItemTagRequestDto Getter
func (r AlibabaalihousenewhomeitemtagsubmitAPIRequest) GetItemTagRequestDto() *ItemTagRequestDto {
	return r._itemTagRequestDto
}
