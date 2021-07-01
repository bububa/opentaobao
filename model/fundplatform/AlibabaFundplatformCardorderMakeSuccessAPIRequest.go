package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardorderMakeSuccessAPIRequest
通知制卡成功 API请求
alibaba.fundplatform.cardorder.make.success

当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方 */
type AlibabaFundplatformCardorderMakeSuccessAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabaFundplatformCardorderMakeSuccessStruct
}

// NewAlibabaFundplatformCardorderMakeSuccessRequest 初始化AlibabaFundplatformCardorderMakeSuccessAPIRequest对象
func NewAlibabaFundplatformCardorderMakeSuccessRequest() *AlibabaFundplatformCardorderMakeSuccessAPIRequest {
	return &AlibabaFundplatformCardorderMakeSuccessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.make.success"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
// 入参对象
func (r *AlibabaFundplatformCardorderMakeSuccessAPIRequest) SetRequest(_request *AlibabaFundplatformCardorderMakeSuccessStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r AlibabaFundplatformCardorderMakeSuccessAPIRequest) GetRequest() *AlibabaFundplatformCardorderMakeSuccessStruct {
	return r._request
}
