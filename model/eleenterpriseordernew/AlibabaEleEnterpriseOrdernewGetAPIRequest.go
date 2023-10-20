package eleenterpriseordernew

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGetAPIRequest 查询订单详情 API请求
// alibaba.ele.enterprise.ordernew.get
//
// 查询订单详情
type AlibabaEleEnterpriseOrdernewGetAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 电话号码
	_phone string
}

// NewAlibabaEleEnterpriseOrdernewGetRequest 初始化AlibabaEleEnterpriseOrdernewGetAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewGetRequest() *AlibabaEleEnterpriseOrdernewGetAPIRequest {
	return &AlibabaEleEnterpriseOrdernewGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseOrdernewGetAPIRequest) Reset() {
	r._orderId = ""
	r._phone = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseOrdernewGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 饿了么订单ID
func (r *AlibabaEleEnterpriseOrdernewGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaEleEnterpriseOrdernewGetAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetPhone is Phone Setter
// 电话号码
func (r *AlibabaEleEnterpriseOrdernewGetAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaEleEnterpriseOrdernewGetAPIRequest) GetPhone() string {
	return r._phone
}

var poolAlibabaEleEnterpriseOrdernewGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseOrdernewGetRequest()
	},
}

// GetAlibabaEleEnterpriseOrdernewGetRequest 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewGetAPIRequest
func GetAlibabaEleEnterpriseOrdernewGetAPIRequest() *AlibabaEleEnterpriseOrdernewGetAPIRequest {
	return poolAlibabaEleEnterpriseOrdernewGetAPIRequest.Get().(*AlibabaEleEnterpriseOrdernewGetAPIRequest)
}

// ReleaseAlibabaEleEnterpriseOrdernewGetAPIRequest 将 AlibabaEleEnterpriseOrdernewGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewGetAPIRequest(v *AlibabaEleEnterpriseOrdernewGetAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewGetAPIRequest.Put(v)
}
