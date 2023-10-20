package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeItemTagSubmitAPIRequest ETC上翻商品打标接口 API请求
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
type AlibabaAlihouseNewhomeItemTagSubmitAPIRequest struct {
	model.Params
	//  打标结构体
	_itemTagRequestDto *ItemTagRequestDto
}

// NewAlibabaAlihouseNewhomeItemTagSubmitRequest 初始化AlibabaAlihouseNewhomeItemTagSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeItemTagSubmitRequest() *AlibabaAlihouseNewhomeItemTagSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeItemTagSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) Reset() {
	r._itemTagRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.item.tag.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemTagRequestDto is ItemTagRequestDto Setter
//
//	打标结构体
func (r *AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) SetItemTagRequestDto(_itemTagRequestDto *ItemTagRequestDto) error {
	r._itemTagRequestDto = _itemTagRequestDto
	r.Set("item_tag_request_dto", _itemTagRequestDto)
	return nil
}

// GetItemTagRequestDto ItemTagRequestDto Getter
func (r AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) GetItemTagRequestDto() *ItemTagRequestDto {
	return r._itemTagRequestDto
}

var poolAlibabaAlihouseNewhomeItemTagSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeItemTagSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeItemTagSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeItemTagSubmitAPIRequest
func GetAlibabaAlihouseNewhomeItemTagSubmitAPIRequest() *AlibabaAlihouseNewhomeItemTagSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeItemTagSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeItemTagSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeItemTagSubmitAPIRequest 将 AlibabaAlihouseNewhomeItemTagSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeItemTagSubmitAPIRequest(v *AlibabaAlihouseNewhomeItemTagSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeItemTagSubmitAPIRequest.Put(v)
}
