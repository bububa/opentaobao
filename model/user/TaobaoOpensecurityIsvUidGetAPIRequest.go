package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopensecurityisvuidgetAPIRequest 获取open security uid for isv API请求
// taobao.opensecurity.isv.uid.get
//
// 根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
type TaobaoopensecurityisvuidgetAPIRequest struct {
	model.Params
	// 基于Appkey生成的open security淘宝用户Id
	_openUid string
}

// NewTaobaoopensecurityisvuidgetRequest 初始化TaobaoopensecurityisvuidgetAPIRequest对象
func NewTaobaoopensecurityisvuidgetRequest() *TaobaoopensecurityisvuidgetAPIRequest {
	return &TaobaoopensecurityisvuidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopensecurityisvuidgetAPIRequest) GetApiMethodName() string {
	return "taobao.opensecurity.isv.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopensecurityisvuidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopensecurityisvuidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenUid is OpenUid Setter
// 基于Appkey生成的open security淘宝用户Id
func (r *TaobaoopensecurityisvuidgetAPIRequest) SetOpenUid(_openUid string) error {
	r._openUid = _openUid
	r.Set("open_uid", _openUid)
	return nil
}

// GetOpenUid OpenUid Getter
func (r TaobaoopensecurityisvuidgetAPIRequest) GetOpenUid() string {
	return r._openUid
}
