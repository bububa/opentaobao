package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoelifelifecardconsumeAPIRequest 品牌惠卡券核销 API请求
// taobao.elife.lifecard.consume
//
// 用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
type TaobaoelifelifecardconsumeAPIRequest struct {
	model.Params
	// 交易请求参数
	_consumeRequest *ConsumeRequest
}

// NewTaobaoelifelifecardconsumeRequest 初始化TaobaoelifelifecardconsumeAPIRequest对象
func NewTaobaoelifelifecardconsumeRequest() *TaobaoelifelifecardconsumeAPIRequest {
	return &TaobaoelifelifecardconsumeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoelifelifecardconsumeAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoelifelifecardconsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoelifelifecardconsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsumeRequest is ConsumeRequest Setter
// 交易请求参数
func (r *TaobaoelifelifecardconsumeAPIRequest) SetConsumeRequest(_consumeRequest *ConsumeRequest) error {
	r._consumeRequest = _consumeRequest
	r.Set("consume_request", _consumeRequest)
	return nil
}

// GetConsumeRequest ConsumeRequest Getter
func (r TaobaoelifelifecardconsumeAPIRequest) GetConsumeRequest() *ConsumeRequest {
	return r._consumeRequest
}
