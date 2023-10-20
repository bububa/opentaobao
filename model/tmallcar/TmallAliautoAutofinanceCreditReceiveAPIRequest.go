package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoAutofinanceCreditReceiveAPIRequest 接收授信结果通知 API请求
// tmall.aliauto.autofinance.credit.receive
//
// 天猫汽车的金融业务场景中，需要接收外部ISV对用户授信申请的通知结果.
type TmallAliautoAutofinanceCreditReceiveAPIRequest struct {
	model.Params
	// 授信通知描述
	_creditReceiveDto *CreditReceiveDto
}

// NewTmallAliautoAutofinanceCreditReceiveRequest 初始化TmallAliautoAutofinanceCreditReceiveAPIRequest对象
func NewTmallAliautoAutofinanceCreditReceiveRequest() *TmallAliautoAutofinanceCreditReceiveAPIRequest {
	return &TmallAliautoAutofinanceCreditReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoAutofinanceCreditReceiveAPIRequest) Reset() {
	r._creditReceiveDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.autofinance.credit.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreditReceiveDto is CreditReceiveDto Setter
// 授信通知描述
func (r *TmallAliautoAutofinanceCreditReceiveAPIRequest) SetCreditReceiveDto(_creditReceiveDto *CreditReceiveDto) error {
	r._creditReceiveDto = _creditReceiveDto
	r.Set("credit_receive_dto", _creditReceiveDto)
	return nil
}

// GetCreditReceiveDto CreditReceiveDto Getter
func (r TmallAliautoAutofinanceCreditReceiveAPIRequest) GetCreditReceiveDto() *CreditReceiveDto {
	return r._creditReceiveDto
}

var poolTmallAliautoAutofinanceCreditReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoAutofinanceCreditReceiveRequest()
	},
}

// GetTmallAliautoAutofinanceCreditReceiveRequest 从 sync.Pool 获取 TmallAliautoAutofinanceCreditReceiveAPIRequest
func GetTmallAliautoAutofinanceCreditReceiveAPIRequest() *TmallAliautoAutofinanceCreditReceiveAPIRequest {
	return poolTmallAliautoAutofinanceCreditReceiveAPIRequest.Get().(*TmallAliautoAutofinanceCreditReceiveAPIRequest)
}

// ReleaseTmallAliautoAutofinanceCreditReceiveAPIRequest 将 TmallAliautoAutofinanceCreditReceiveAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoAutofinanceCreditReceiveAPIRequest(v *TmallAliautoAutofinanceCreditReceiveAPIRequest) {
	v.Reset()
	poolTmallAliautoAutofinanceCreditReceiveAPIRequest.Put(v)
}
