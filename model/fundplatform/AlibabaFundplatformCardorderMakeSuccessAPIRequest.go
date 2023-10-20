package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordermakesuccessAPIRequest 通知制卡成功 API请求
// alibaba.fundplatform.cardorder.make.success
//
// 当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
type AlibabafundplatformcardordermakesuccessAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabafundplatformcardordermakesuccessStruct
}

// NewAlibabafundplatformcardordermakesuccessRequest 初始化AlibabafundplatformcardordermakesuccessAPIRequest对象
func NewAlibabafundplatformcardordermakesuccessRequest() *AlibabafundplatformcardordermakesuccessAPIRequest {
	return &AlibabafundplatformcardordermakesuccessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardordermakesuccessAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.make.success"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardordermakesuccessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardordermakesuccessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 入参对象
func (r *AlibabafundplatformcardordermakesuccessAPIRequest) SetRequest(_request *AlibabafundplatformcardordermakesuccessStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r AlibabafundplatformcardordermakesuccessAPIRequest) GetRequest() *AlibabafundplatformcardordermakesuccessStruct {
	return r._request
}
