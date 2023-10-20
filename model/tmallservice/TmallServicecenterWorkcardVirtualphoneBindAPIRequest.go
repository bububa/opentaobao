package tmallservice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardVirtualphoneBindAPIRequest) Reset() {
	r._workcardRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.virtualphone.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardRequest is WorkcardRequest Setter
// 绑定阿里通讯号入参
func (r *TmallServicecenterWorkcardVirtualphoneBindAPIRequest) SetWorkcardRequest(_workcardRequest *WorkcardBaseRequest) error {
	r._workcardRequest = _workcardRequest
	r.Set("workcard_request", _workcardRequest)
	return nil
}

// GetWorkcardRequest WorkcardRequest Getter
func (r TmallServicecenterWorkcardVirtualphoneBindAPIRequest) GetWorkcardRequest() *WorkcardBaseRequest {
	return r._workcardRequest
}

var poolTmallServicecenterWorkcardVirtualphoneBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardVirtualphoneBindRequest()
	},
}

// GetTmallServicecenterWorkcardVirtualphoneBindRequest 从 sync.Pool 获取 TmallServicecenterWorkcardVirtualphoneBindAPIRequest
func GetTmallServicecenterWorkcardVirtualphoneBindAPIRequest() *TmallServicecenterWorkcardVirtualphoneBindAPIRequest {
	return poolTmallServicecenterWorkcardVirtualphoneBindAPIRequest.Get().(*TmallServicecenterWorkcardVirtualphoneBindAPIRequest)
}

// ReleaseTmallServicecenterWorkcardVirtualphoneBindAPIRequest 将 TmallServicecenterWorkcardVirtualphoneBindAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardVirtualphoneBindAPIRequest(v *TmallServicecenterWorkcardVirtualphoneBindAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardVirtualphoneBindAPIRequest.Put(v)
}
