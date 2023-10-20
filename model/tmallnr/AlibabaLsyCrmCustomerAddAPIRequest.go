package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCustomerAddAPIRequest 私域导购添加活动留资入口 API请求
// alibaba.lsy.crm.customer.add
//
// 私域导购添加活动留资入口
type AlibabaLsyCrmCustomerAddAPIRequest struct {
	model.Params
	// 入参对象
	_reqDto *NrtCrmCreateCustomerReq
}

// NewAlibabaLsyCrmCustomerAddRequest 初始化AlibabaLsyCrmCustomerAddAPIRequest对象
func NewAlibabaLsyCrmCustomerAddRequest() *AlibabaLsyCrmCustomerAddAPIRequest {
	return &AlibabaLsyCrmCustomerAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLsyCrmCustomerAddAPIRequest) Reset() {
	r._reqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqDto is ReqDto Setter
// 入参对象
func (r *AlibabaLsyCrmCustomerAddAPIRequest) SetReqDto(_reqDto *NrtCrmCreateCustomerReq) error {
	r._reqDto = _reqDto
	r.Set("req_dto", _reqDto)
	return nil
}

// GetReqDto ReqDto Getter
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetReqDto() *NrtCrmCreateCustomerReq {
	return r._reqDto
}

var poolAlibabaLsyCrmCustomerAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLsyCrmCustomerAddRequest()
	},
}

// GetAlibabaLsyCrmCustomerAddRequest 从 sync.Pool 获取 AlibabaLsyCrmCustomerAddAPIRequest
func GetAlibabaLsyCrmCustomerAddAPIRequest() *AlibabaLsyCrmCustomerAddAPIRequest {
	return poolAlibabaLsyCrmCustomerAddAPIRequest.Get().(*AlibabaLsyCrmCustomerAddAPIRequest)
}

// ReleaseAlibabaLsyCrmCustomerAddAPIRequest 将 AlibabaLsyCrmCustomerAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaLsyCrmCustomerAddAPIRequest(v *AlibabaLsyCrmCustomerAddAPIRequest) {
	v.Reset()
	poolAlibabaLsyCrmCustomerAddAPIRequest.Put(v)
}
