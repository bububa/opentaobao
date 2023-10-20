package yunosappstore

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstorePadHpApplistAPIRequest 查询HpPad appList API请求
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
type YunosAppstorePadHpApplistAPIRequest struct {
	model.Params
	// 获取的应用群code
	_code string
}

// NewYunosAppstorePadHpApplistRequest 初始化YunosAppstorePadHpApplistAPIRequest对象
func NewYunosAppstorePadHpApplistRequest() *YunosAppstorePadHpApplistAPIRequest {
	return &YunosAppstorePadHpApplistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAppstorePadHpApplistAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAppstorePadHpApplistAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.pad.hp.applist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAppstorePadHpApplistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAppstorePadHpApplistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 获取的应用群code
func (r *YunosAppstorePadHpApplistAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r YunosAppstorePadHpApplistAPIRequest) GetCode() string {
	return r._code
}

var poolYunosAppstorePadHpApplistAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAppstorePadHpApplistRequest()
	},
}

// GetYunosAppstorePadHpApplistRequest 从 sync.Pool 获取 YunosAppstorePadHpApplistAPIRequest
func GetYunosAppstorePadHpApplistAPIRequest() *YunosAppstorePadHpApplistAPIRequest {
	return poolYunosAppstorePadHpApplistAPIRequest.Get().(*YunosAppstorePadHpApplistAPIRequest)
}

// ReleaseYunosAppstorePadHpApplistAPIRequest 将 YunosAppstorePadHpApplistAPIRequest 放入 sync.Pool
func ReleaseYunosAppstorePadHpApplistAPIRequest(v *YunosAppstorePadHpApplistAPIRequest) {
	v.Reset()
	poolYunosAppstorePadHpApplistAPIRequest.Put(v)
}
