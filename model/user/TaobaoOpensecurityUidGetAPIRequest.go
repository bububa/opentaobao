package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpensecurityUidGetAPIRequest 淘宝open security uid 获取接口 API请求
// taobao.opensecurity.uid.get
//
// 根据明文 taobao user id 换取 app的 open_uid
type TaobaoOpensecurityUidGetAPIRequest struct {
	model.Params
	// 淘宝用户ID
	_tbUserId int64
}

// NewTaobaoOpensecurityUidGetRequest 初始化TaobaoOpensecurityUidGetAPIRequest对象
func NewTaobaoOpensecurityUidGetRequest() *TaobaoOpensecurityUidGetAPIRequest {
	return &TaobaoOpensecurityUidGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpensecurityUidGetAPIRequest) Reset() {
	r._tbUserId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpensecurityUidGetAPIRequest) GetApiMethodName() string {
	return "taobao.opensecurity.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpensecurityUidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpensecurityUidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbUserId is TbUserId Setter
// 淘宝用户ID
func (r *TaobaoOpensecurityUidGetAPIRequest) SetTbUserId(_tbUserId int64) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// GetTbUserId TbUserId Getter
func (r TaobaoOpensecurityUidGetAPIRequest) GetTbUserId() int64 {
	return r._tbUserId
}

var poolTaobaoOpensecurityUidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpensecurityUidGetRequest()
	},
}

// GetTaobaoOpensecurityUidGetRequest 从 sync.Pool 获取 TaobaoOpensecurityUidGetAPIRequest
func GetTaobaoOpensecurityUidGetAPIRequest() *TaobaoOpensecurityUidGetAPIRequest {
	return poolTaobaoOpensecurityUidGetAPIRequest.Get().(*TaobaoOpensecurityUidGetAPIRequest)
}

// ReleaseTaobaoOpensecurityUidGetAPIRequest 将 TaobaoOpensecurityUidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpensecurityUidGetAPIRequest(v *TaobaoOpensecurityUidGetAPIRequest) {
	v.Reset()
	poolTaobaoOpensecurityUidGetAPIRequest.Put(v)
}
