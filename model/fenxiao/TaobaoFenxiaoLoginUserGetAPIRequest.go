package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoLoginUserGetAPIRequest 获取分销用户登录信息 API请求
// taobao.fenxiao.login.user.get
//
// 获取用户登录信息
type TaobaoFenxiaoLoginUserGetAPIRequest struct {
	model.Params
}

// NewTaobaoFenxiaoLoginUserGetRequest 初始化TaobaoFenxiaoLoginUserGetAPIRequest对象
func NewTaobaoFenxiaoLoginUserGetRequest() *TaobaoFenxiaoLoginUserGetAPIRequest {
	return &TaobaoFenxiaoLoginUserGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoLoginUserGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoLoginUserGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.login.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoLoginUserGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoLoginUserGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoFenxiaoLoginUserGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoLoginUserGetRequest()
	},
}

// GetTaobaoFenxiaoLoginUserGetRequest 从 sync.Pool 获取 TaobaoFenxiaoLoginUserGetAPIRequest
func GetTaobaoFenxiaoLoginUserGetAPIRequest() *TaobaoFenxiaoLoginUserGetAPIRequest {
	return poolTaobaoFenxiaoLoginUserGetAPIRequest.Get().(*TaobaoFenxiaoLoginUserGetAPIRequest)
}

// ReleaseTaobaoFenxiaoLoginUserGetAPIRequest 将 TaobaoFenxiaoLoginUserGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoLoginUserGetAPIRequest(v *TaobaoFenxiaoLoginUserGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoLoginUserGetAPIRequest.Put(v)
}
