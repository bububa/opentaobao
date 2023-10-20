package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopensecurityuidgetAPIRequest 淘宝open security uid 获取接口 API请求
// taobao.opensecurity.uid.get
//
// 根据明文 taobao user id 换取 app的 open_uid
type TaobaoopensecurityuidgetAPIRequest struct {
	model.Params
	// 淘宝用户ID
	_tbUserId int64
}

// NewTaobaoopensecurityuidgetRequest 初始化TaobaoopensecurityuidgetAPIRequest对象
func NewTaobaoopensecurityuidgetRequest() *TaobaoopensecurityuidgetAPIRequest {
	return &TaobaoopensecurityuidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopensecurityuidgetAPIRequest) GetApiMethodName() string {
	return "taobao.opensecurity.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopensecurityuidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopensecurityuidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbUserId is TbUserId Setter
// 淘宝用户ID
func (r *TaobaoopensecurityuidgetAPIRequest) SetTbUserId(_tbUserId int64) error {
	r._tbUserId = _tbUserId
	r.Set("tb_user_id", _tbUserId)
	return nil
}

// GetTbUserId TbUserId Getter
func (r TaobaoopensecurityuidgetAPIRequest) GetTbUserId() int64 {
	return r._tbUserId
}
