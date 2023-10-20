package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressRemoveAPIRequest 删除卖家地址库 API请求
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
type TaobaoLogisticsAddressRemoveAPIRequest struct {
	model.Params
	// 地址库ID
	_contactId int64
}

// NewTaobaoLogisticsAddressRemoveRequest 初始化TaobaoLogisticsAddressRemoveAPIRequest对象
func NewTaobaoLogisticsAddressRemoveRequest() *TaobaoLogisticsAddressRemoveAPIRequest {
	return &TaobaoLogisticsAddressRemoveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsAddressRemoveAPIRequest) Reset() {
	r._contactId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContactId is ContactId Setter
// 地址库ID
func (r *TaobaoLogisticsAddressRemoveAPIRequest) SetContactId(_contactId int64) error {
	r._contactId = _contactId
	r.Set("contact_id", _contactId)
	return nil
}

// GetContactId ContactId Getter
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetContactId() int64 {
	return r._contactId
}

var poolTaobaoLogisticsAddressRemoveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsAddressRemoveRequest()
	},
}

// GetTaobaoLogisticsAddressRemoveRequest 从 sync.Pool 获取 TaobaoLogisticsAddressRemoveAPIRequest
func GetTaobaoLogisticsAddressRemoveAPIRequest() *TaobaoLogisticsAddressRemoveAPIRequest {
	return poolTaobaoLogisticsAddressRemoveAPIRequest.Get().(*TaobaoLogisticsAddressRemoveAPIRequest)
}

// ReleaseTaobaoLogisticsAddressRemoveAPIRequest 将 TaobaoLogisticsAddressRemoveAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsAddressRemoveAPIRequest(v *TaobaoLogisticsAddressRemoveAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsAddressRemoveAPIRequest.Put(v)
}
