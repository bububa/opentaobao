package refund

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRdcAligeniusSendgoodsCancelAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetApiMethodName() string {
	return "taobao.rdc.aligenius.sendgoods.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusSendgoodsCancelAPIRequest) SetParam(_param *CancelGoodsDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoRdcAligeniusSendgoodsCancelAPIRequest) GetParam() *CancelGoodsDto {
	return r._param
}

var poolTaobaoRdcAligeniusSendgoodsCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRdcAligeniusSendgoodsCancelRequest()
	},
}

// GetTaobaoRdcAligeniusSendgoodsCancelRequest 从 sync.Pool 获取 TaobaoRdcAligeniusSendgoodsCancelAPIRequest
func GetTaobaoRdcAligeniusSendgoodsCancelAPIRequest() *TaobaoRdcAligeniusSendgoodsCancelAPIRequest {
	return poolTaobaoRdcAligeniusSendgoodsCancelAPIRequest.Get().(*TaobaoRdcAligeniusSendgoodsCancelAPIRequest)
}

// ReleaseTaobaoRdcAligeniusSendgoodsCancelAPIRequest 将 TaobaoRdcAligeniusSendgoodsCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoRdcAligeniusSendgoodsCancelAPIRequest(v *TaobaoRdcAligeniusSendgoodsCancelAPIRequest) {
	v.Reset()
	poolTaobaoRdcAligeniusSendgoodsCancelAPIRequest.Put(v)
}
