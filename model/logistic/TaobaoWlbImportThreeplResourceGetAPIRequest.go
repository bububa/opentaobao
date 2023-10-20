package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportthreeplresourcegetAPIRequest 3PL直邮获取资源列表 API请求
// taobao.wlb.import.threepl.resource.get
//
// 获取3pl直邮的发货可用资源
type TaobaowlbimportthreeplresourcegetAPIRequest struct {
	model.Params
	// ONLINE或者OFFLINE
	_type string
	// 发货地区域id
	_fromId int64
	// 收件人地址
	_toAddress *AddressDto
}

// NewTaobaowlbimportthreeplresourcegetRequest 初始化TaobaowlbimportthreeplresourcegetAPIRequest对象
func NewTaobaowlbimportthreeplresourcegetRequest() *TaobaowlbimportthreeplresourcegetAPIRequest {
	return &TaobaowlbimportthreeplresourcegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportthreeplresourcegetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.import.threepl.resource.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportthreeplresourcegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportthreeplresourcegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// ONLINE或者OFFLINE
func (r *TaobaowlbimportthreeplresourcegetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaowlbimportthreeplresourcegetAPIRequest) GetType() string {
	return r._type
}

// SetFromId is FromId Setter
// 发货地区域id
func (r *TaobaowlbimportthreeplresourcegetAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r TaobaowlbimportthreeplresourcegetAPIRequest) GetFromId() int64 {
	return r._fromId
}

// SetToAddress is ToAddress Setter
// 收件人地址
func (r *TaobaowlbimportthreeplresourcegetAPIRequest) SetToAddress(_toAddress *AddressDto) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r TaobaowlbimportthreeplresourcegetAPIRequest) GetToAddress() *AddressDto {
	return r._toAddress
}
