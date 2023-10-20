package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest 提交经纪人信息 API请求
// alibaba.alihouse.existinghome.broker.submit
//
// 提交经纪人信息
type AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest struct {
	model.Params
	// 经纪人
	_brokers []ProjectAdviserDto
}

// NewAlibabaAlihouseExistinghomeBrokerSubmitRequest 初始化AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest对象
func NewAlibabaAlihouseExistinghomeBrokerSubmitRequest() *AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest {
	return &AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) Reset() {
	r._brokers = r._brokers[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.broker.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrokers is Brokers Setter
// 经纪人
func (r *AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) SetBrokers(_brokers []ProjectAdviserDto) error {
	r._brokers = _brokers
	r.Set("brokers", _brokers)
	return nil
}

// GetBrokers Brokers Getter
func (r AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) GetBrokers() []ProjectAdviserDto {
	return r._brokers
}

var poolAlibabaAlihouseExistinghomeBrokerSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeBrokerSubmitRequest()
	},
}

// GetAlibabaAlihouseExistinghomeBrokerSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest
func GetAlibabaAlihouseExistinghomeBrokerSubmitAPIRequest() *AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest {
	return poolAlibabaAlihouseExistinghomeBrokerSubmitAPIRequest.Get().(*AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerSubmitAPIRequest 将 AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeBrokerSubmitAPIRequest(v *AlibabaAlihouseExistinghomeBrokerSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeBrokerSubmitAPIRequest.Put(v)
}
