package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateosfotapublishAPIRequest 系统升级发布 API请求
// yunos.osupdate.osfota.publish
//
// 发布osupdate系统升级任务
type YunososupdateosfotapublishAPIRequest struct {
	model.Params
	// 入参json格式
	_publishJson string
}

// NewYunososupdateosfotapublishRequest 初始化YunososupdateosfotapublishAPIRequest对象
func NewYunososupdateosfotapublishRequest() *YunososupdateosfotapublishAPIRequest {
	return &YunososupdateosfotapublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateosfotapublishAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateosfotapublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateosfotapublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublishJson is PublishJson Setter
// 入参json格式
func (r *YunososupdateosfotapublishAPIRequest) SetPublishJson(_publishJson string) error {
	r._publishJson = _publishJson
	r.Set("publish_json", _publishJson)
	return nil
}

// GetPublishJson PublishJson Getter
func (r YunososupdateosfotapublishAPIRequest) GetPublishJson() string {
	return r._publishJson
}
