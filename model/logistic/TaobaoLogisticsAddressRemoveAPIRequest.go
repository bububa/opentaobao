package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsaddressremoveAPIRequest 删除卖家地址库 API请求
// taobao.logistics.address.remove
//
// 用此接口删除卖家地址库
type TaobaologisticsaddressremoveAPIRequest struct {
	model.Params
	// 地址库ID
	_contactId int64
}

// NewTaobaologisticsaddressremoveRequest 初始化TaobaologisticsaddressremoveAPIRequest对象
func NewTaobaologisticsaddressremoveRequest() *TaobaologisticsaddressremoveAPIRequest {
	return &TaobaologisticsaddressremoveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsaddressremoveAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.remove"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsaddressremoveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsaddressremoveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContactId is ContactId Setter
// 地址库ID
func (r *TaobaologisticsaddressremoveAPIRequest) SetContactId(_contactId int64) error {
	r._contactId = _contactId
	r.Set("contact_id", _contactId)
	return nil
}

// GetContactId ContactId Getter
func (r TaobaologisticsaddressremoveAPIRequest) GetContactId() int64 {
	return r._contactId
}
