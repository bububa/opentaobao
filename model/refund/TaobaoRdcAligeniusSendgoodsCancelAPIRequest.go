package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRdcAligeniusSendgoodsCancelAPIRequest 取消发货 API请求
// taobao.rdc.aligenius.sendgoods.cancel
//
// 提供商家在仅退款中发送取消发货状态
type TaobaoRdcAligeniusSendgoodsCancelAPIRequest struct {
	model.Params
	// 请求参数
	_param *CancelGoodsDto
}

// NewTaobaoRdcAligeniusSendgoodsCancelRequest 初始化TaobaoRdcAligeniusSendgoodsCancelAPIRequest对象
func NewTaobaoRdcAligeniusSendgoodsCancelRequest() *TaobaoRdcAligeniusSendgoodsCancelAPIRequest {
	return &TaobaoRdcAligeniusSendgoodsCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.sendgoods.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusSendgoodsCancelAPIRequest) SetParam(_param *CancelGoodsDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetParam() *CancelGoodsDto {
	return r._param
}
