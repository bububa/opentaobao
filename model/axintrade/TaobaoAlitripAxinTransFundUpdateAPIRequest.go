package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransFundUpdateAPIRequest 修改资金单接口 API请求
// taobao.alitrip.axin.trans.fund.update
//
// 阿信供销平台-修改资金单接口
type TaobaoAlitripAxinTransFundUpdateAPIRequest struct {
	model.Params
	// 更新资金单接口入参
	_axinFundUpdateDTO *AxinFundUpdateDto
}

// NewTaobaoAlitripAxinTransFundUpdateRequest 初始化TaobaoAlitripAxinTransFundUpdateAPIRequest对象
func NewTaobaoAlitripAxinTransFundUpdateRequest() *TaobaoAlitripAxinTransFundUpdateAPIRequest {
	return &TaobaoAlitripAxinTransFundUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripAxinTransFundUpdateAPIRequest) Reset() {
	r._axinFundUpdateDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripAxinTransFundUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAxinFundUpdateDTO is AxinFundUpdateDTO Setter
// 更新资金单接口入参
func (r *TaobaoAlitripAxinTransFundUpdateAPIRequest) SetAxinFundUpdateDTO(_axinFundUpdateDTO *AxinFundUpdateDto) error {
	r._axinFundUpdateDTO = _axinFundUpdateDTO
	r.Set("axin_fund_update_d_t_o", _axinFundUpdateDTO)
	return nil
}

// GetAxinFundUpdateDTO AxinFundUpdateDTO Getter
func (r TaobaoAlitripAxinTransFundUpdateAPIRequest) GetAxinFundUpdateDTO() *AxinFundUpdateDto {
	return r._axinFundUpdateDTO
}

var poolTaobaoAlitripAxinTransFundUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripAxinTransFundUpdateRequest()
	},
}

// GetTaobaoAlitripAxinTransFundUpdateRequest 从 sync.Pool 获取 TaobaoAlitripAxinTransFundUpdateAPIRequest
func GetTaobaoAlitripAxinTransFundUpdateAPIRequest() *TaobaoAlitripAxinTransFundUpdateAPIRequest {
	return poolTaobaoAlitripAxinTransFundUpdateAPIRequest.Get().(*TaobaoAlitripAxinTransFundUpdateAPIRequest)
}

// ReleaseTaobaoAlitripAxinTransFundUpdateAPIRequest 将 TaobaoAlitripAxinTransFundUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripAxinTransFundUpdateAPIRequest(v *TaobaoAlitripAxinTransFundUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAlitripAxinTransFundUpdateAPIRequest.Put(v)
}
