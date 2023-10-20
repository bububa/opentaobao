package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardrefuseAPIRequest 买家拒收 API请求
// tmall.servicecenter.workcard.refuse
//
// 买家拒收通知接口
type TmallservicecenterworkcardrefuseAPIRequest struct {
	model.Params
	// 买家拒收信息
	_buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest
}

// NewTmallservicecenterworkcardrefuseRequest 初始化TmallservicecenterworkcardrefuseAPIRequest对象
func NewTmallservicecenterworkcardrefuseRequest() *TmallservicecenterworkcardrefuseAPIRequest {
	return &TmallservicecenterworkcardrefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardrefuseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardrefuseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardrefuseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerRefuseAcceptRequest is BuyerRefuseAcceptRequest Setter
// 买家拒收信息
func (r *TmallservicecenterworkcardrefuseAPIRequest) SetBuyerRefuseAcceptRequest(_buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest) error {
	r._buyerRefuseAcceptRequest = _buyerRefuseAcceptRequest
	r.Set("buyer_refuse_accept_request", _buyerRefuseAcceptRequest)
	return nil
}

// GetBuyerRefuseAcceptRequest BuyerRefuseAcceptRequest Getter
func (r TmallservicecenterworkcardrefuseAPIRequest) GetBuyerRefuseAcceptRequest() *BuyerRefuseAcceptRequest {
	return r._buyerRefuseAcceptRequest
}
