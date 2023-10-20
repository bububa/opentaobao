package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressModifyAppointAPIRequest 快递改约api API请求
// taobao.logistics.express.modify.appoint
//
// 商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
type TaobaoLogisticsExpressModifyAppointAPIRequest struct {
	model.Params
	// 改约请求对象
	_expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto
}

// NewTaobaoLogisticsExpressModifyAppointRequest 初始化TaobaoLogisticsExpressModifyAppointAPIRequest对象
func NewTaobaoLogisticsExpressModifyAppointRequest() *TaobaoLogisticsExpressModifyAppointAPIRequest {
	return &TaobaoLogisticsExpressModifyAppointAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressModifyAppointAPIRequest) Reset() {
	r._expressModifyAppointTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressModifyAppointAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.modify.appoint"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressModifyAppointAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressModifyAppointAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExpressModifyAppointTopRequest is ExpressModifyAppointTopRequest Setter
// 改约请求对象
func (r *TaobaoLogisticsExpressModifyAppointAPIRequest) SetExpressModifyAppointTopRequest(_expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto) error {
	r._expressModifyAppointTopRequest = _expressModifyAppointTopRequest
	r.Set("express_modify_appoint_top_request", _expressModifyAppointTopRequest)
	return nil
}

// GetExpressModifyAppointTopRequest ExpressModifyAppointTopRequest Getter
func (r TaobaoLogisticsExpressModifyAppointAPIRequest) GetExpressModifyAppointTopRequest() *ExpressModifyAppointTopRequestDto {
	return r._expressModifyAppointTopRequest
}

var poolTaobaoLogisticsExpressModifyAppointAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressModifyAppointRequest()
	},
}

// GetTaobaoLogisticsExpressModifyAppointRequest 从 sync.Pool 获取 TaobaoLogisticsExpressModifyAppointAPIRequest
func GetTaobaoLogisticsExpressModifyAppointAPIRequest() *TaobaoLogisticsExpressModifyAppointAPIRequest {
	return poolTaobaoLogisticsExpressModifyAppointAPIRequest.Get().(*TaobaoLogisticsExpressModifyAppointAPIRequest)
}

// ReleaseTaobaoLogisticsExpressModifyAppointAPIRequest 将 TaobaoLogisticsExpressModifyAppointAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressModifyAppointAPIRequest(v *TaobaoLogisticsExpressModifyAppointAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressModifyAppointAPIRequest.Put(v)
}
