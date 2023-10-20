package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerQueryAPIRequest 根据外部经纪人ID查询 API请求
// alibaba.alihouse.existinghome.broker.query
//
// 根据外部经纪人ID查询
type AlibabaAlihouseExistinghomeBrokerQueryAPIRequest struct {
	model.Params
	// 外部经纪人ID
	_outerBrokerId string
}

// NewAlibabaAlihouseExistinghomeBrokerQueryRequest 初始化AlibabaAlihouseExistinghomeBrokerQueryAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrokerQueryRequest() *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest {
	return &AlibabaAlihouseExistinghomeBrokerQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) Reset() {
	r._outerBrokerId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterBrokerId is OuterBrokerId Setter
// 外部经纪人ID
func (r *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) SetOuterBrokerId(_outerBrokerId string) error {
	r._outerBrokerId = _outerBrokerId
	r.Set("outer_broker_id", _outerBrokerId)
	return nil
}

// GetOuterBrokerId OuterBrokerId Getter
func (r AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) GetOuterBrokerId() string {
	return r._outerBrokerId
}

var poolAlibabaAlihouseExistinghomeBrokerQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeBrokerQueryRequest()
	},
}

// GetAlibabaAlihouseExistinghomeBrokerQueryRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrokerQueryAPIRequest
func GetAlibabaAlihouseExistinghomeBrokerQueryAPIRequest() *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest {
	return poolAlibabaAlihouseExistinghomeBrokerQueryAPIRequest.Get().(*AlibabaAlihouseExistinghomeBrokerQueryAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerQueryAPIRequest 将 AlibabaAlihouseExistinghomeBrokerQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrokerQueryAPIRequest(v *AlibabaAlihouseExistinghomeBrokerQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrokerQueryAPIRequest.Put(v)
}
