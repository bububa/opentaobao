package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardRefuseAPIRequest 买家拒收 API请求
// tmall.servicecenter.workcard.refuse
//
// 买家拒收通知接口
type TmallServicecenterWorkcardRefuseAPIRequest struct {
	model.Params
	// 买家拒收信息
	_buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest
}

// NewTmallServicecenterWorkcardRefuseRequest 初始化TmallServicecenterWorkcardRefuseAPIRequest对象
func NewTmallServicecenterWorkcardRefuseRequest() *TmallServicecenterWorkcardRefuseAPIRequest {
	return &TmallServicecenterWorkcardRefuseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardRefuseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.refuse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardRefuseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BuyerRefuseAcceptRequest Setter
// 买家拒收信息
func (r *TmallServicecenterWorkcardRefuseAPIRequest) SetBuyerRefuseAcceptRequest(_buyerRefuseAcceptRequest *BuyerRefuseAcceptRequest) error {
	r._buyerRefuseAcceptRequest = _buyerRefuseAcceptRequest
	r.Set("buyer_refuse_accept_request", _buyerRefuseAcceptRequest)
	return nil
}

// Get BuyerRefuseAcceptRequest Getter
func (r TmallServicecenterWorkcardRefuseAPIRequest) GetBuyerRefuseAcceptRequest() *BuyerRefuseAcceptRequest {
	return r._buyerRefuseAcceptRequest
}
