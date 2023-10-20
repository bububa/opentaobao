package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardvirtualphonebindAPIRequest 工单维度虚拟中间号绑定 API请求
// tmall.servicecenter.workcard.virtualphone.bind
//
// 服务供应链洗护服务ERP项目中，客服呼叫消费者的功能。
// 叫消费者的手机号虚拟号码给到客服。
type TmallservicecenterworkcardvirtualphonebindAPIRequest struct {
	model.Params
	// 绑定阿里通讯号入参
	_workcardRequest *WorkcardBaseRequest
}

// NewTmallservicecenterworkcardvirtualphonebindRequest 初始化TmallservicecenterworkcardvirtualphonebindAPIRequest对象
func NewTmallservicecenterworkcardvirtualphonebindRequest() *TmallservicecenterworkcardvirtualphonebindAPIRequest {
	return &TmallservicecenterworkcardvirtualphonebindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardvirtualphonebindAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.virtualphone.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardvirtualphonebindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardvirtualphonebindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardRequest is WorkcardRequest Setter
// 绑定阿里通讯号入参
func (r *TmallservicecenterworkcardvirtualphonebindAPIRequest) SetWorkcardRequest(_workcardRequest *WorkcardBaseRequest) error {
	r._workcardRequest = _workcardRequest
	r.Set("workcard_request", _workcardRequest)
	return nil
}

// GetWorkcardRequest WorkcardRequest Getter
func (r TmallservicecenterworkcardvirtualphonebindAPIRequest) GetWorkcardRequest() *WorkcardBaseRequest {
	return r._workcardRequest
}
