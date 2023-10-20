package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoeticketstatusAPIRequest 查询电子凭证状态 API请求
// tmall.aliauto.eticket.status
//
// 查询天猫汽车二轮车行业门店电子凭证状态
type TmallaliautoeticketstatusAPIRequest struct {
	model.Params
	// 核销码
	_consumeCode string
}

// NewTmallaliautoeticketstatusRequest 初始化TmallaliautoeticketstatusAPIRequest对象
func NewTmallaliautoeticketstatusRequest() *TmallaliautoeticketstatusAPIRequest {
	return &TmallaliautoeticketstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoeticketstatusAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.eticket.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoeticketstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoeticketstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsumeCode is ConsumeCode Setter
// 核销码
func (r *TmallaliautoeticketstatusAPIRequest) SetConsumeCode(_consumeCode string) error {
	r._consumeCode = _consumeCode
	r.Set("consume_code", _consumeCode)
	return nil
}

// GetConsumeCode ConsumeCode Getter
func (r TmallaliautoeticketstatusAPIRequest) GetConsumeCode() string {
	return r._consumeCode
}
