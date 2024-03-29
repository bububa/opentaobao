package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest 创建服务履约物流单 API请求
// tmall.servicecenter.workcard.logisticsinfo.create
//
// 创建服务履约物流单
type TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest struct {
	model.Params
	// 创建服务履约物流单信息
	_createLogisticsOrderRequest *CreateLogisticsOrderRequest
}

// NewTmallServicecenterWorkcardLogisticsinfoCreateRequest 初始化TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest对象
func NewTmallServicecenterWorkcardLogisticsinfoCreateRequest() *TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest {
	return &TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) Reset() {
	r._createLogisticsOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsinfo.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateLogisticsOrderRequest is CreateLogisticsOrderRequest Setter
// 创建服务履约物流单信息
func (r *TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) SetCreateLogisticsOrderRequest(_createLogisticsOrderRequest *CreateLogisticsOrderRequest) error {
	r._createLogisticsOrderRequest = _createLogisticsOrderRequest
	r.Set("create_logistics_order_request", _createLogisticsOrderRequest)
	return nil
}

// GetCreateLogisticsOrderRequest CreateLogisticsOrderRequest Getter
func (r TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) GetCreateLogisticsOrderRequest() *CreateLogisticsOrderRequest {
	return r._createLogisticsOrderRequest
}

var poolTmallServicecenterWorkcardLogisticsinfoCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardLogisticsinfoCreateRequest()
	},
}

// GetTmallServicecenterWorkcardLogisticsinfoCreateRequest 从 sync.Pool 获取 TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest
func GetTmallServicecenterWorkcardLogisticsinfoCreateAPIRequest() *TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest {
	return poolTmallServicecenterWorkcardLogisticsinfoCreateAPIRequest.Get().(*TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest)
}

// ReleaseTmallServicecenterWorkcardLogisticsinfoCreateAPIRequest 将 TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardLogisticsinfoCreateAPIRequest(v *TmallServicecenterWorkcardLogisticsinfoCreateAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardLogisticsinfoCreateAPIRequest.Put(v)
}
