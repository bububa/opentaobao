package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpensecurityIsvUidGetAPIRequest 获取open security uid for isv API请求
// taobao.opensecurity.isv.uid.get
//
// 根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
type TaobaoOpensecurityIsvUidGetAPIRequest struct {
	model.Params
	// 基于Appkey生成的open security淘宝用户Id
	_openUid string
}

// NewTaobaoOpensecurityIsvUidGetRequest 初始化TaobaoOpensecurityIsvUidGetAPIRequest对象
func NewTaobaoOpensecurityIsvUidGetRequest() *TaobaoOpensecurityIsvUidGetAPIRequest {
	return &TaobaoOpensecurityIsvUidGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpensecurityIsvUidGetAPIRequest) Reset() {
	r._openUid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpensecurityIsvUidGetAPIRequest) GetApiMethodName() string {
	return "taobao.opensecurity.isv.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpensecurityIsvUidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpensecurityIsvUidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUid is OpenUid Setter
// 基于Appkey生成的open security淘宝用户Id
func (r *TaobaoOpensecurityIsvUidGetAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r TaobaoOpensecurityIsvUidGetAPIRequest) GetOpenUid() string {
	return r._openUid
}

var poolTaobaoOpensecurityIsvUidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpensecurityIsvUidGetRequest()
	},
}

// GetTaobaoOpensecurityIsvUidGetRequest 从 sync.Pool 获取 TaobaoOpensecurityIsvUidGetAPIRequest
func GetTaobaoOpensecurityIsvUidGetAPIRequest() *TaobaoOpensecurityIsvUidGetAPIRequest {
	return poolTaobaoOpensecurityIsvUidGetAPIRequest.Get().(*TaobaoOpensecurityIsvUidGetAPIRequest)
}

// ReleaseTaobaoOpensecurityIsvUidGetAPIRequest 将 TaobaoOpensecurityIsvUidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpensecurityIsvUidGetAPIRequest(v *TaobaoOpensecurityIsvUidGetAPIRequest) {
	v.Reset()
	poolTaobaoOpensecurityIsvUidGetAPIRequest.Put(v)
}
