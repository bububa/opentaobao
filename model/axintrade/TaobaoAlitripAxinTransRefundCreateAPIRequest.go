package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransRefundCreateAPIRequest 阿信创建退款单 API请求
// taobao.alitrip.axin.trans.refund.create
//
// 阿信供销平台-创建退款单服务
type TaobaoAlitripAxinTransRefundCreateAPIRequest struct {
	model.Params
	// 阿信创建退款单入参
	_axinRefundCreateDTO *AxinRefundCreateDto
}

// NewTaobaoAlitripAxinTransRefundCreateRequest 初始化TaobaoAlitripAxinTransRefundCreateAPIRequest对象
func NewTaobaoAlitripAxinTransRefundCreateRequest() *TaobaoAlitripAxinTransRefundCreateAPIRequest {
	return &TaobaoAlitripAxinTransRefundCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripAxinTransRefundCreateAPIRequest) Reset() {
	r._axinRefundCreateDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinRefundCreateDTO is AxinRefundCreateDTO Setter
// 阿信创建退款单入参
func (r *TaobaoAlitripAxinTransRefundCreateAPIRequest) SetAxinRefundCreateDTO(_axinRefundCreateDTO *AxinRefundCreateDto) error {
	r._axinRefundCreateDTO = _axinRefundCreateDTO
	r.Set("axin_refund_create_d_t_o", _axinRefundCreateDTO)
	return nil
}

// GetAxinRefundCreateDTO AxinRefundCreateDTO Getter
func (r TaobaoAlitripAxinTransRefundCreateAPIRequest) GetAxinRefundCreateDTO() *AxinRefundCreateDto {
	return r._axinRefundCreateDTO
}

var poolTaobaoAlitripAxinTransRefundCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripAxinTransRefundCreateRequest()
	},
}

// GetTaobaoAlitripAxinTransRefundCreateRequest 从 sync.Pool 获取 TaobaoAlitripAxinTransRefundCreateAPIRequest
func GetTaobaoAlitripAxinTransRefundCreateAPIRequest() *TaobaoAlitripAxinTransRefundCreateAPIRequest {
	return poolTaobaoAlitripAxinTransRefundCreateAPIRequest.Get().(*TaobaoAlitripAxinTransRefundCreateAPIRequest)
}

// ReleaseTaobaoAlitripAxinTransRefundCreateAPIRequest 将 TaobaoAlitripAxinTransRefundCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripAxinTransRefundCreateAPIRequest(v *TaobaoAlitripAxinTransRefundCreateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripAxinTransRefundCreateAPIRequest.Put(v)
}
