package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionPublishAPIRequest 发布应用升级 API请求
// yunos.osupdate.appversion.publish
//
// 发布应用升级任务
type YunosOsupdateAppversionPublishAPIRequest struct {
	model.Params
	// 发布应用升级入参json
	_publishJson string
}

// NewYunosOsupdateAppversionPublishRequest 初始化YunosOsupdateAppversionPublishAPIRequest对象
func NewYunosOsupdateAppversionPublishRequest() *YunosOsupdateAppversionPublishAPIRequest {
	return &YunosOsupdateAppversionPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionPublishAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateAppversionPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublishJson is PublishJson Setter
// 发布应用升级入参json
func (r *YunosOsupdateAppversionPublishAPIRequest) SetPublishJson(_publishJson string) error {
	r._publishJson = _publishJson
	r.Set("publish_json", _publishJson)
	return nil
}

// GetPublishJson PublishJson Getter
func (r YunosOsupdateAppversionPublishAPIRequest) GetPublishJson() string {
	return r._publishJson
}
