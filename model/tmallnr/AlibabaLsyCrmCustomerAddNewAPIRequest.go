package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCustomerAddNewAPIRequest 导购域留资接口 API请求
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
type AlibabaLsyCrmCustomerAddNewAPIRequest struct {
	model.Params
	// 入参DTO
	_reqDTO *NrtCrmCreateCustomerReq
}

// NewAlibabaLsyCrmCustomerAddNewRequest 初始化AlibabaLsyCrmCustomerAddNewAPIRequest对象
func NewAlibabaLsyCrmCustomerAddNewRequest() *AlibabaLsyCrmCustomerAddNewAPIRequest {
	return &AlibabaLsyCrmCustomerAddNewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLsyCrmCustomerAddNewAPIRequest) Reset() {
	r._reqDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCustomerAddNewAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.add.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCustomerAddNewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmCustomerAddNewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqDTO is ReqDTO Setter
// 入参DTO
func (r *AlibabaLsyCrmCustomerAddNewAPIRequest) SetReqDTO(_reqDTO *NrtCrmCreateCustomerReq) error {
	r._reqDTO = _reqDTO
	r.Set("req_d_t_o", _reqDTO)
	return nil
}

// GetReqDTO ReqDTO Getter
func (r AlibabaLsyCrmCustomerAddNewAPIRequest) GetReqDTO() *NrtCrmCreateCustomerReq {
	return r._reqDTO
}

var poolAlibabaLsyCrmCustomerAddNewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLsyCrmCustomerAddNewRequest()
	},
}

// GetAlibabaLsyCrmCustomerAddNewRequest 从 sync.Pool 获取 AlibabaLsyCrmCustomerAddNewAPIRequest
func GetAlibabaLsyCrmCustomerAddNewAPIRequest() *AlibabaLsyCrmCustomerAddNewAPIRequest {
	return poolAlibabaLsyCrmCustomerAddNewAPIRequest.Get().(*AlibabaLsyCrmCustomerAddNewAPIRequest)
}

// ReleaseAlibabaLsyCrmCustomerAddNewAPIRequest 将 AlibabaLsyCrmCustomerAddNewAPIRequest 放入 sync.Pool
func ReleaseAlibabaLsyCrmCustomerAddNewAPIRequest(v *AlibabaLsyCrmCustomerAddNewAPIRequest) {
	v.Reset()
	poolAlibabaLsyCrmCustomerAddNewAPIRequest.Put(v)
}
