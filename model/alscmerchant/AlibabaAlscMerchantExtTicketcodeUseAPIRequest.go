package alscmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscMerchantExtTicketcodeUseAPIRequest 外部核销服务 API请求
// alibaba.alsc.merchant.ext.ticketcode.use
//
// 外部核销服务
type AlibabaAlscMerchantExtTicketcodeUseAPIRequest struct {
	model.Params
	// 外部券使用请求
	_useRequest *ExternalTicketUseRequest
}

// NewAlibabaAlscMerchantExtTicketcodeUseRequest 初始化AlibabaAlscMerchantExtTicketcodeUseAPIRequest对象
func NewAlibabaAlscMerchantExtTicketcodeUseRequest() *AlibabaAlscMerchantExtTicketcodeUseAPIRequest {
	return &AlibabaAlscMerchantExtTicketcodeUseAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscMerchantExtTicketcodeUseAPIRequest) Reset() {
	r._useRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.merchant.ext.ticketcode.use"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUseRequest is UseRequest Setter
// 外部券使用请求
func (r *AlibabaAlscMerchantExtTicketcodeUseAPIRequest) SetUseRequest(_useRequest *ExternalTicketUseRequest) error {
	r._useRequest = _useRequest
	r.Set("use_request", _useRequest)
	return nil
}

// GetUseRequest UseRequest Getter
func (r AlibabaAlscMerchantExtTicketcodeUseAPIRequest) GetUseRequest() *ExternalTicketUseRequest {
	return r._useRequest
}

var poolAlibabaAlscMerchantExtTicketcodeUseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscMerchantExtTicketcodeUseRequest()
	},
}

// GetAlibabaAlscMerchantExtTicketcodeUseRequest 从 sync.Pool 获取 AlibabaAlscMerchantExtTicketcodeUseAPIRequest
func GetAlibabaAlscMerchantExtTicketcodeUseAPIRequest() *AlibabaAlscMerchantExtTicketcodeUseAPIRequest {
	return poolAlibabaAlscMerchantExtTicketcodeUseAPIRequest.Get().(*AlibabaAlscMerchantExtTicketcodeUseAPIRequest)
}

// ReleaseAlibabaAlscMerchantExtTicketcodeUseAPIRequest 将 AlibabaAlscMerchantExtTicketcodeUseAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscMerchantExtTicketcodeUseAPIRequest(v *AlibabaAlscMerchantExtTicketcodeUseAPIRequest) {
	v.Reset()
	poolAlibabaAlscMerchantExtTicketcodeUseAPIRequest.Put(v)
}
