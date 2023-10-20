package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosappstorepadhpapplistAPIRequest 查询HpPad appList API请求
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
type YunosappstorepadhpapplistAPIRequest struct {
	model.Params
	// 获取的应用群code
	_code string
}

// NewYunosappstorepadhpapplistRequest 初始化YunosappstorepadhpapplistAPIRequest对象
func NewYunosappstorepadhpapplistRequest() *YunosappstorepadhpapplistAPIRequest {
	return &YunosappstorepadhpapplistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosappstorepadhpapplistAPIRequest) GetApiMethodName() string {
	return "yunos.appstore.pad.hp.applist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosappstorepadhpapplistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosappstorepadhpapplistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 获取的应用群code
func (r *YunosappstorepadhpapplistAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r YunosappstorepadhpapplistAPIRequest) GetCode() string {
	return r._code
}
