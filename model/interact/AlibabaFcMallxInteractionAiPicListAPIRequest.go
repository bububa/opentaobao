package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafcmallxinteractionaipiclistAPIRequest 花园ai作画定制查询 API请求
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
type AlibabafcmallxinteractionaipiclistAPIRequest struct {
	model.Params
	// openid
	_openid string
	// openid
	_openId string
	// appKey
	_appId string
	// 商品信息
	_param *InteractiveTopItemParam
}

// NewAlibabafcmallxinteractionaipiclistRequest 初始化AlibabafcmallxinteractionaipiclistAPIRequest对象
func NewAlibabafcmallxinteractionaipiclistRequest() *AlibabafcmallxinteractionaipiclistAPIRequest {
	return &AlibabafcmallxinteractionaipiclistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetApiMethodName() string {
	return "alibaba.fc.mallx.interaction.ai.pic.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenid is Openid Setter
// openid
func (r *AlibabafcmallxinteractionaipiclistAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetOpenid() string {
	return r._openid
}

// SetOpenId is OpenId Setter
// openid
func (r *AlibabafcmallxinteractionaipiclistAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetOpenId() string {
	return r._openId
}

// SetAppId is AppId Setter
// appKey
func (r *AlibabafcmallxinteractionaipiclistAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetAppId() string {
	return r._appId
}

// SetParam is Param Setter
// 商品信息
func (r *AlibabafcmallxinteractionaipiclistAPIRequest) SetParam(_param *InteractiveTopItemParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabafcmallxinteractionaipiclistAPIRequest) GetParam() *InteractiveTopItemParam {
	return r._param
}
