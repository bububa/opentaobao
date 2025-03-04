package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaUdUdsmartdataGetAPIRequest udsmart获取数据回传接口 API请求
// alibaba.ud.udsmartdata.get
//
// udsmart获取数据回传接口
type AlibabaUdUdsmartdataGetAPIRequest struct {
	model.Params
	// ud_smart_service_request
	_udSmartServiceRequest *UDSmartServiceRequest
}

// NewAlibabaUdUdsmartdataGetRequest 初始化AlibabaUdUdsmartdataGetAPIRequest对象
func NewAlibabaUdUdsmartdataGetRequest() *AlibabaUdUdsmartdataGetAPIRequest {
	return &AlibabaUdUdsmartdataGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaUdUdsmartdataGetAPIRequest) Reset() {
	r._udSmartServiceRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaUdUdsmartdataGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ud.udsmartdata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaUdUdsmartdataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaUdUdsmartdataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUdSmartServiceRequest is UdSmartServiceRequest Setter
// ud_smart_service_request
func (r *AlibabaUdUdsmartdataGetAPIRequest) SetUdSmartServiceRequest(_udSmartServiceRequest *UDSmartServiceRequest) error {
	r._udSmartServiceRequest = _udSmartServiceRequest
	r.Set("ud_smart_service_request", _udSmartServiceRequest)
	return nil
}

// GetUdSmartServiceRequest UdSmartServiceRequest Getter
func (r AlibabaUdUdsmartdataGetAPIRequest) GetUdSmartServiceRequest() *UDSmartServiceRequest {
	return r._udSmartServiceRequest
}

var poolAlibabaUdUdsmartdataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaUdUdsmartdataGetRequest()
	},
}

// GetAlibabaUdUdsmartdataGetRequest 从 sync.Pool 获取 AlibabaUdUdsmartdataGetAPIRequest
func GetAlibabaUdUdsmartdataGetAPIRequest() *AlibabaUdUdsmartdataGetAPIRequest {
	return poolAlibabaUdUdsmartdataGetAPIRequest.Get().(*AlibabaUdUdsmartdataGetAPIRequest)
}

// ReleaseAlibabaUdUdsmartdataGetAPIRequest 将 AlibabaUdUdsmartdataGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaUdUdsmartdataGetAPIRequest(v *AlibabaUdUdsmartdataGetAPIRequest) {
	v.Reset()
	poolAlibabaUdUdsmartdataGetAPIRequest.Put(v)
}
