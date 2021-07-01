package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsAddressRemoveAPIRequest
删除卖家地址库 API请求
taobao.logistics.address.remove

用此接口删除卖家地址库 */
type TaobaoLogisticsAddressRemoveAPIRequest struct {
	model.Params
	// 地址库ID
	_contactId int64
}

// NewTaobaoLogisticsAddressRemoveRequest 初始化TaobaoLogisticsAddressRemoveAPIRequest对象
func NewTaobaoLogisticsAddressRemoveRequest() *TaobaoLogisticsAddressRemoveAPIRequest {
	return &TaobaoLogisticsAddressRemoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ContactId Setter
// 地址库ID
func (r *TaobaoLogisticsAddressRemoveAPIRequest) SetContactId(_contactId int64) error {
	r._contactId = _contactId
	r.Set("contact_id", _contactId)
	return nil
}

// Get ContactId Getter
func (r TaobaoLogisticsAddressRemoveAPIRequest) GetContactId() int64 {
	return r._contactId
}
