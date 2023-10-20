package cityretail

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCityretailWmflConvertWarehouseAPIRequest 同城零售完美履约转仓 API请求
// taobao.cityretail.wmfl.convert.warehouse
//
// 同城零售完美履约转仓
type TaobaoCityretailWmflConvertWarehouseAPIRequest struct {
	model.Params
	// 淘宝交易单id
	_tbOrderId string
}

// NewTaobaoCityretailWmflConvertWarehouseRequest 初始化TaobaoCityretailWmflConvertWarehouseAPIRequest对象
func NewTaobaoCityretailWmflConvertWarehouseRequest() *TaobaoCityretailWmflConvertWarehouseAPIRequest {
	return &TaobaoCityretailWmflConvertWarehouseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCityretailWmflConvertWarehouseAPIRequest) Reset() {
	r._tbOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCityretailWmflConvertWarehouseAPIRequest) GetApiMethodName() string {
	return "taobao.cityretail.wmfl.convert.warehouse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCityretailWmflConvertWarehouseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCityretailWmflConvertWarehouseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbOrderId is TbOrderId Setter
// 淘宝交易单id
func (r *TaobaoCityretailWmflConvertWarehouseAPIRequest) SetTbOrderId(_tbOrderId string) error {
	r._tbOrderId = _tbOrderId
	r.Set("tb_order_id", _tbOrderId)
	return nil
}

// GetTbOrderId TbOrderId Getter
func (r TaobaoCityretailWmflConvertWarehouseAPIRequest) GetTbOrderId() string {
	return r._tbOrderId
}

var poolTaobaoCityretailWmflConvertWarehouseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCityretailWmflConvertWarehouseRequest()
	},
}

// GetTaobaoCityretailWmflConvertWarehouseRequest 从 sync.Pool 获取 TaobaoCityretailWmflConvertWarehouseAPIRequest
func GetTaobaoCityretailWmflConvertWarehouseAPIRequest() *TaobaoCityretailWmflConvertWarehouseAPIRequest {
	return poolTaobaoCityretailWmflConvertWarehouseAPIRequest.Get().(*TaobaoCityretailWmflConvertWarehouseAPIRequest)
}

// ReleaseTaobaoCityretailWmflConvertWarehouseAPIRequest 将 TaobaoCityretailWmflConvertWarehouseAPIRequest 放入 sync.Pool
func ReleaseTaobaoCityretailWmflConvertWarehouseAPIRequest(v *TaobaoCityretailWmflConvertWarehouseAPIRequest) {
	v.Reset()
	poolTaobaoCityretailWmflConvertWarehouseAPIRequest.Put(v)
}
