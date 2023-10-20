package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateappversionpublishAPIRequest 发布应用升级 API请求
// yunos.osupdate.appversion.publish
//
// 发布应用升级任务
type YunososupdateappversionpublishAPIRequest struct {
	model.Params
	// 发布应用升级入参json
	_publishJson string
}

// NewYunososupdateappversionpublishRequest 初始化YunososupdateappversionpublishAPIRequest对象
func NewYunososupdateappversionpublishRequest() *YunososupdateappversionpublishAPIRequest {
	return &YunososupdateappversionpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateappversionpublishAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateappversionpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateappversionpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPublishJson is PublishJson Setter
// 发布应用升级入参json
func (r *YunososupdateappversionpublishAPIRequest) SetPublishJson(_publishJson string) error {
	r._publishJson = _publishJson
	r.Set("publish_json", _publishJson)
	return nil
}

// GetPublishJson PublishJson Getter
func (r YunososupdateappversionpublishAPIRequest) GetPublishJson() string {
	return r._publishJson
}
