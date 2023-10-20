package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest 补发单物流信息回传 API请求
// taobao.rdc.aligenius.warehouse.resend.logistics.msg.post
//
// 补发单erp物流信息回传平台
type TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest struct {
	model.Params
	// 参数
	_param0 *SendResendLogisticsMsgDto
}

// NewTaobaordcaligeniuswarehouseresendlogisticsmsgpostRequest 初始化TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest对象
func NewTaobaordcaligeniuswarehouseresendlogisticsmsgpostRequest() *TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest {
	return &TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.warehouse.resend.logistics.msg.post"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest) SetParam0(_param0 *SendResendLogisticsMsgDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaordcaligeniuswarehouseresendlogisticsmsgpostAPIRequest) GetParam0() *SendResendLogisticsMsgDto {
	return r._param0
}
