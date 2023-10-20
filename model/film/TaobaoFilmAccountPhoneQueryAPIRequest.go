package film

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFilmAccountPhoneQueryAPIRequest 根据手机查询匹配账号列表 API请求
// taobao.film.account.phone.query
//
// 根据手机号查询匹配的账号列表
type TaobaoFilmAccountPhoneQueryAPIRequest struct {
	model.Params
	// 手机号
	_phone string
}

// NewTaobaoFilmAccountPhoneQueryRequest 初始化TaobaoFilmAccountPhoneQueryAPIRequest对象
func NewTaobaoFilmAccountPhoneQueryRequest() *TaobaoFilmAccountPhoneQueryAPIRequest {
	return &TaobaoFilmAccountPhoneQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFilmAccountPhoneQueryAPIRequest) Reset() {
	r._phone = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFilmAccountPhoneQueryAPIRequest) GetApiMethodName() string {
	return "taobao.film.account.phone.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFilmAccountPhoneQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFilmAccountPhoneQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *TaobaoFilmAccountPhoneQueryAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaoFilmAccountPhoneQueryAPIRequest) GetPhone() string {
	return r._phone
}

var poolTaobaoFilmAccountPhoneQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFilmAccountPhoneQueryRequest()
	},
}

// GetTaobaoFilmAccountPhoneQueryRequest 从 sync.Pool 获取 TaobaoFilmAccountPhoneQueryAPIRequest
func GetTaobaoFilmAccountPhoneQueryAPIRequest() *TaobaoFilmAccountPhoneQueryAPIRequest {
	return poolTaobaoFilmAccountPhoneQueryAPIRequest.Get().(*TaobaoFilmAccountPhoneQueryAPIRequest)
}

// ReleaseTaobaoFilmAccountPhoneQueryAPIRequest 将 TaobaoFilmAccountPhoneQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoFilmAccountPhoneQueryAPIRequest(v *TaobaoFilmAccountPhoneQueryAPIRequest) {
	v.Reset()
	poolTaobaoFilmAccountPhoneQueryAPIRequest.Put(v)
}
