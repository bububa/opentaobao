package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKaolaScitemAddAPIRequest 考拉货品新增接口 API请求
// taobao.kaola.scitem.add
//
// 考拉货品新增接口
type TaobaoKaolaScitemAddAPIRequest struct {
	model.Params
	// 待新增的货品
	_cnsku *CnskuDto
	// 新增选项
	_option *AddCnskuOption
}

// NewTaobaoKaolaScitemAddRequest 初始化TaobaoKaolaScitemAddAPIRequest对象
func NewTaobaoKaolaScitemAddRequest() *TaobaoKaolaScitemAddAPIRequest {
	return &TaobaoKaolaScitemAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoKaolaScitemAddAPIRequest) Reset() {
	r._cnsku = nil
	r._option = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKaolaScitemAddAPIRequest) GetApiMethodName() string {
	return "taobao.kaola.scitem.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKaolaScitemAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoKaolaScitemAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnsku is Cnsku Setter
// 待新增的货品
func (r *TaobaoKaolaScitemAddAPIRequest) SetCnsku(_cnsku *CnskuDto) error {
	r._cnsku = _cnsku
	r.Set("cnsku", _cnsku)
	return nil
}

// GetCnsku Cnsku Getter
func (r TaobaoKaolaScitemAddAPIRequest) GetCnsku() *CnskuDto {
	return r._cnsku
}

// SetOption is Option Setter
// 新增选项
func (r *TaobaoKaolaScitemAddAPIRequest) SetOption(_option *AddCnskuOption) error {
	r._option = _option
	r.Set("option", _option)
	return nil
}

// GetOption Option Getter
func (r TaobaoKaolaScitemAddAPIRequest) GetOption() *AddCnskuOption {
	return r._option
}

var poolTaobaoKaolaScitemAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoKaolaScitemAddRequest()
	},
}

// GetTaobaoKaolaScitemAddRequest 从 sync.Pool 获取 TaobaoKaolaScitemAddAPIRequest
func GetTaobaoKaolaScitemAddAPIRequest() *TaobaoKaolaScitemAddAPIRequest {
	return poolTaobaoKaolaScitemAddAPIRequest.Get().(*TaobaoKaolaScitemAddAPIRequest)
}

// ReleaseTaobaoKaolaScitemAddAPIRequest 将 TaobaoKaolaScitemAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoKaolaScitemAddAPIRequest(v *TaobaoKaolaScitemAddAPIRequest) {
	v.Reset()
	poolTaobaoKaolaScitemAddAPIRequest.Put(v)
}
