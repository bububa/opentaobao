package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest 补发单物流信息回传 API请求
// taobao.rdc.aligenius.warehouse.resend.logistics.msg.post
//
// 补发单erp物流信息回传平台
type TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest struct {
	model.Params
	// 参数
	_param0 *SendResendLogisticsMsgDto
}

// NewTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest 初始化TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest对象
func NewTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest() *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest {
	return &TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.resend.logistics.msg.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) SetParam0(_param0 *SendResendLogisticsMsgDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) GetParam0() *SendResendLogisticsMsgDto {
	return r._param0
}

var poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest()
	},
}

// GetTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostRequest 从 sync.Pool 获取 TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest
func GetTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest() *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest {
	return poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest.Get().(*TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest)
}

// ReleaseTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest 将 TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest(v *TaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusWarehouseResendLogisticsMsgPostAPIRequest.Put(v)
}
