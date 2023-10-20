package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPaySignCheckAPIRequest 阿信支付宝验签服务 API请求
// taobao.alitrip.axin.trans.pay.sign.check
//
// 阿信支付宝验签服务
type TaobaoAlitripAxinTransPaySignCheckAPIRequest struct {
	model.Params
	// 验签对象
	_axinPayCheckSignDto *AxinPayCheckSignDto
}

// NewTaobaoAlitripAxinTransPaySignCheckRequest 初始化TaobaoAlitripAxinTransPaySignCheckAPIRequest对象
func NewTaobaoAlitripAxinTransPaySignCheckRequest() *TaobaoAlitripAxinTransPaySignCheckAPIRequest {
	return &TaobaoAlitripAxinTransPaySignCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripAxinTransPaySignCheckAPIRequest) Reset() {
	r._axinPayCheckSignDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.sign.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinPayCheckSignDto is AxinPayCheckSignDto Setter
// 验签对象
func (r *TaobaoAlitripAxinTransPaySignCheckAPIRequest) SetAxinPayCheckSignDto(_axinPayCheckSignDto *AxinPayCheckSignDto) error {
	r._axinPayCheckSignDto = _axinPayCheckSignDto
	r.Set("axin_pay_check_sign_dto", _axinPayCheckSignDto)
	return nil
}

// GetAxinPayCheckSignDto AxinPayCheckSignDto Getter
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetAxinPayCheckSignDto() *AxinPayCheckSignDto {
	return r._axinPayCheckSignDto
}

var poolTaobaoAlitripAxinTransPaySignCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripAxinTransPaySignCheckRequest()
	},
}

// GetTaobaoAlitripAxinTransPaySignCheckRequest 从 sync.Pool 获取 TaobaoAlitripAxinTransPaySignCheckAPIRequest
func GetTaobaoAlitripAxinTransPaySignCheckAPIRequest() *TaobaoAlitripAxinTransPaySignCheckAPIRequest {
	return poolTaobaoAlitripAxinTransPaySignCheckAPIRequest.Get().(*TaobaoAlitripAxinTransPaySignCheckAPIRequest)
}

// ReleaseTaobaoAlitripAxinTransPaySignCheckAPIRequest 将 TaobaoAlitripAxinTransPaySignCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripAxinTransPaySignCheckAPIRequest(v *TaobaoAlitripAxinTransPaySignCheckAPIRequest) {
	v.Reset()
	poolTaobaoAlitripAxinTransPaySignCheckAPIRequest.Put(v)
}
