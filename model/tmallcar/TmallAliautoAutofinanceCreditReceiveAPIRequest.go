package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoautofinancecreditreceiveAPIRequest 接收授信结果通知 API请求
// tmall.aliauto.autofinance.credit.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
type TmallaliautoautofinancecreditreceiveAPIRequest struct {
	model.Params
	// 授信通知描述
	_creditReceiveDto *CreditReceiveDto
}

// NewTmallaliautoautofinancecreditreceiveRequest 初始化TmallaliautoautofinancecreditreceiveAPIRequest对象
func NewTmallaliautoautofinancecreditreceiveRequest() *TmallaliautoautofinancecreditreceiveAPIRequest {
	return &TmallaliautoautofinancecreditreceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoautofinancecreditreceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.autofinance.credit.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoautofinancecreditreceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoautofinancecreditreceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreditReceiveDto is CreditReceiveDto Setter
// 授信通知描述
func (r *TmallaliautoautofinancecreditreceiveAPIRequest) SetCreditReceiveDto(_creditReceiveDto *CreditReceiveDto) error {
	r._creditReceiveDto = _creditReceiveDto
	r.Set("credit_receive_dto", _creditReceiveDto)
	return nil
}

// GetCreditReceiveDto CreditReceiveDto Getter
func (r TmallaliautoautofinancecreditreceiveAPIRequest) GetCreditReceiveDto() *CreditReceiveDto {
	return r._creditReceiveDto
}
