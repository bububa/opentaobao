package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemAddAPIRequest 租赁商品发布 API请求
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
type AlibabaIdleRentItemAddAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// NewAlibabaIdleRentItemAddRequest 初始化AlibabaIdleRentItemAddAPIRequest对象
func NewAlibabaIdleRentItemAddRequest() *AlibabaIdleRentItemAddAPIRequest {
	return &AlibabaIdleRentItemAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentItemAddAPIRequest) Reset() {
	r._paramRentItemDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemAddAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentItemAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRentItemDTO is ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaIdleRentItemAddAPIRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDto) error {
	r._paramRentItemDTO = _paramRentItemDTO
	r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
	return nil
}

// GetParamRentItemDTO ParamRentItemDTO Getter
func (r AlibabaIdleRentItemAddAPIRequest) GetParamRentItemDTO() *RentItemDto {
	return r._paramRentItemDTO
}

var poolAlibabaIdleRentItemAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentItemAddRequest()
	},
}

// GetAlibabaIdleRentItemAddRequest 从 sync.Pool 获取 AlibabaIdleRentItemAddAPIRequest
func GetAlibabaIdleRentItemAddAPIRequest() *AlibabaIdleRentItemAddAPIRequest {
	return poolAlibabaIdleRentItemAddAPIRequest.Get().(*AlibabaIdleRentItemAddAPIRequest)
}

// ReleaseAlibabaIdleRentItemAddAPIRequest 将 AlibabaIdleRentItemAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentItemAddAPIRequest(v *AlibabaIdleRentItemAddAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentItemAddAPIRequest.Put(v)
}
