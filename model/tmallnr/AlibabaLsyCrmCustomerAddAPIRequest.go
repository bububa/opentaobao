package tmallnr

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.customer.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReqDto Setter
// 入参对象
func (r *AlibabaLsyCrmCustomerAddAPIRequest) SetReqDto(_reqDto *NrtCrmCreateCustomerReq) error {
	r._reqDto = _reqDto
	r.Set("req_dto", _reqDto)
	return nil
}

// Get ReqDto Getter
func (r AlibabaLsyCrmCustomerAddAPIRequest) GetReqDto() *NrtCrmCreateCustomerReq {
	return r._reqDto
}
