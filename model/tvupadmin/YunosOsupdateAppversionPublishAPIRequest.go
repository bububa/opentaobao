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
func (r YunosOsupdateAppversionPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PublishJson Setter
// 发布应用升级入参json
func (r *YunosOsupdateAppversionPublishAPIRequest) SetPublishJson(_publishJson string) error {
	r._publishJson = _publishJson
	r.Set("publish_json", _publishJson)
	return nil
}

// Get PublishJson Getter
func (r YunosOsupdateAppversionPublishAPIRequest) GetPublishJson() string {
	return r._publishJson
}
