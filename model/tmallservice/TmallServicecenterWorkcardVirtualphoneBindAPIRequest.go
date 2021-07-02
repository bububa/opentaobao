package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardVirtualphoneBindAPIRequest 工单维度虚拟中间号绑定 API请求
// tmall.servicecenter.workcard.virtualphone.bind
//
// 服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
// 叫消费者的手机号虚拟号码给到客服。
type TmallServicecenterWorkcardVirtualphoneBindAPIRequest struct {
	model.Params
	// 绑定阿里通讯号入参
	_workcardRequest *WorkcardBaseRequest
}

// NewTmallServicecenterWorkcardVirtualphoneBindRequest 初始化TmallServicecenterWorkcardVirtualphoneBindAPIRequest对象
func NewTmallServicecenterWorkcardVirtualphoneBindRequest() *TmallServicecenterWorkcardVirtualphoneBindAPIRequest {
	return &TmallServicecenterWorkcardVirtualphoneBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.virtualphone.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkcardRequest Setter
// 绑定阿里通讯号入参
func (r *TmallServicecenterWorkcardVirtualphoneBindAPIRequest) SetWorkcardRequest(_workcardRequest *WorkcardBaseRequest) error {
	r._workcardRequest = _workcardRequest
	r.Set("workcard_request", _workcardRequest)
	return nil
}

// Get WorkcardRequest Getter
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetWorkcardRequest() *WorkcardBaseRequest {
	return r._workcardRequest
}
