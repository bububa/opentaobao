package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemEditAPIRequest 租赁商品编辑 API请求
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
type AlibabaIdleRentItemEditAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// NewAlibabaIdleRentItemEditRequest 初始化AlibabaIdleRentItemEditAPIRequest对象
func NewAlibabaIdleRentItemEditRequest() *AlibabaIdleRentItemEditAPIRequest {
	return &AlibabaIdleRentItemEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentItemEditAPIRequest) Reset() {
	r._paramRentItemDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemEditAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentItemEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRentItemDTO is ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaIdleRentItemEditAPIRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDto) error {
	r._paramRentItemDTO = _paramRentItemDTO
	r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
	return nil
}

// GetParamRentItemDTO ParamRentItemDTO Getter
func (r AlibabaIdleRentItemEditAPIRequest) GetParamRentItemDTO() *RentItemDto {
	return r._paramRentItemDTO
}

var poolAlibabaIdleRentItemEditAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentItemEditRequest()
	},
}

// GetAlibabaIdleRentItemEditRequest 从 sync.Pool 获取 AlibabaIdleRentItemEditAPIRequest
func GetAlibabaIdleRentItemEditAPIRequest() *AlibabaIdleRentItemEditAPIRequest {
	return poolAlibabaIdleRentItemEditAPIRequest.Get().(*AlibabaIdleRentItemEditAPIRequest)
}

// ReleaseAlibabaIdleRentItemEditAPIRequest 将 AlibabaIdleRentItemEditAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentItemEditAPIRequest(v *AlibabaIdleRentItemEditAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentItemEditAPIRequest.Put(v)
}
