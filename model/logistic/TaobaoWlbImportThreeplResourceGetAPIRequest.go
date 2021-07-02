package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportThreeplResourceGetAPIRequest 3PL直邮获取资源列表 API请求
// taobao.wlb.import.threepl.resource.get
//
// 获取3pl直邮的发货可用资源
type TaobaoWlbImportThreeplResourceGetAPIRequest struct {
	model.Params
	// ONLINE或者OFFLINE
	_type string
	// 发货地区域id
	_fromId int64
	// 收件人地址
	_toAddress *AddressDto
}

// NewTaobaoWlbImportThreeplResourceGetRequest 初始化TaobaoWlbImportThreeplResourceGetAPIRequest对象
func NewTaobaoWlbImportThreeplResourceGetRequest() *TaobaoWlbImportThreeplResourceGetAPIRequest {
	return &TaobaoWlbImportThreeplResourceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.import.threepl.resource.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Type Setter
// ONLINE或者OFFLINE
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetType() string {
	return r._type
}

// Set is FromId Setter
// 发货地区域id
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// Get FromId Getter
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetFromId() int64 {
	return r._fromId
}

// Set is ToAddress Setter
// 收件人地址
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) SetToAddress(_toAddress *AddressDto) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// Get ToAddress Getter
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetToAddress() *AddressDto {
	return r._toAddress
}
