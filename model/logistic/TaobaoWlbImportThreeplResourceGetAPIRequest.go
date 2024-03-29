package logistic

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) Reset() {
	r._type = ""
	r._fromId = 0
	r._toAddress = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.import.threepl.resource.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// ONLINE或者OFFLINE
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetType() string {
	return r._type
}

// SetFromId is FromId Setter
// 发货地区域id
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) SetFromId(_fromId int64) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetFromId() int64 {
	return r._fromId
}

// SetToAddress is ToAddress Setter
// 收件人地址
func (r *TaobaoWlbImportThreeplResourceGetAPIRequest) SetToAddress(_toAddress *AddressDto) error {
	r._toAddress = _toAddress
	r.Set("to_address", _toAddress)
	return nil
}

// GetToAddress ToAddress Getter
func (r TaobaoWlbImportThreeplResourceGetAPIRequest) GetToAddress() *AddressDto {
	return r._toAddress
}

var poolTaobaoWlbImportThreeplResourceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbImportThreeplResourceGetRequest()
	},
}

// GetTaobaoWlbImportThreeplResourceGetRequest 从 sync.Pool 获取 TaobaoWlbImportThreeplResourceGetAPIRequest
func GetTaobaoWlbImportThreeplResourceGetAPIRequest() *TaobaoWlbImportThreeplResourceGetAPIRequest {
	return poolTaobaoWlbImportThreeplResourceGetAPIRequest.Get().(*TaobaoWlbImportThreeplResourceGetAPIRequest)
}

// ReleaseTaobaoWlbImportThreeplResourceGetAPIRequest 将 TaobaoWlbImportThreeplResourceGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbImportThreeplResourceGetAPIRequest(v *TaobaoWlbImportThreeplResourceGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbImportThreeplResourceGetAPIRequest.Put(v)
}
