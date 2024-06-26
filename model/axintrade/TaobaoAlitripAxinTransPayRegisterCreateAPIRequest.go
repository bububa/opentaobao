package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransPayRegisterCreateAPIRequest 提交支付服务开通 API请求
// taobao.alitrip.axin.trans.pay.register.create
//
// 阿信供销平台-提交支付服务开通接口
type TaobaoAlitripAxinTransPayRegisterCreateAPIRequest struct {
	model.Params
	// 提交支付服务开通接口入参
	_createDTO *AxinPayRegisterCreateDto
}

// NewTaobaoAlitripAxinTransPayRegisterCreateRequest 初始化TaobaoAlitripAxinTransPayRegisterCreateAPIRequest对象
func NewTaobaoAlitripAxinTransPayRegisterCreateRequest() *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest {
	return &TaobaoAlitripAxinTransPayRegisterCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) Reset() {
	r._createDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.pay.register.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateDTO is CreateDTO Setter
// 提交支付服务开通接口入参
func (r *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) SetCreateDTO(_createDTO *AxinPayRegisterCreateDto) error {
	r._createDTO = _createDTO
	r.Set("create_d_t_o", _createDTO)
	return nil
}

// GetCreateDTO CreateDTO Getter
func (r TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) GetCreateDTO() *AxinPayRegisterCreateDto {
	return r._createDTO
}

var poolTaobaoAlitripAxinTransPayRegisterCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripAxinTransPayRegisterCreateRequest()
	},
}

// GetTaobaoAlitripAxinTransPayRegisterCreateRequest 从 sync.Pool 获取 TaobaoAlitripAxinTransPayRegisterCreateAPIRequest
func GetTaobaoAlitripAxinTransPayRegisterCreateAPIRequest() *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest {
	return poolTaobaoAlitripAxinTransPayRegisterCreateAPIRequest.Get().(*TaobaoAlitripAxinTransPayRegisterCreateAPIRequest)
}

// ReleaseTaobaoAlitripAxinTransPayRegisterCreateAPIRequest 将 TaobaoAlitripAxinTransPayRegisterCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripAxinTransPayRegisterCreateAPIRequest(v *TaobaoAlitripAxinTransPayRegisterCreateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripAxinTransPayRegisterCreateAPIRequest.Put(v)
}
