package eleenterpriseordernew

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest 订单配送信息跟踪 API请求
// alibaba.ele.enterprise.ordernew.gettrackinginfo
//
// 订单配送信息跟踪
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 用户手机号
	_phone string
}

// NewAlibabaEleEnterpriseOrdernewGettrackinginfoRequest 初始化AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewGettrackinginfoRequest() *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest {
	return &AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) Reset() {
	r._orderId = ""
	r._phone = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.gettrackinginfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPhone is Phone Setter
// 用户手机号
func (r *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) GetPhone() string {
	return r._phone
}

var poolAlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseOrdernewGettrackinginfoRequest()
	},
}

// GetAlibabaEleEnterpriseOrdernewGettrackinginfoRequest 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest
func GetAlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest() *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest {
	return poolAlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest.Get().(*AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest)
}

// ReleaseAlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest 将 AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest(v *AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest.Put(v)
}
