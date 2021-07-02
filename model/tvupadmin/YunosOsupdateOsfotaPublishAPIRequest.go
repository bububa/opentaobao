package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateOsfotaPublishAPIRequest 系统升级发布 API请求
// yunos.osupdate.osfota.publish
//
// 发布osupdate系统升级任务
type YunosOsupdateOsfotaPublishAPIRequest struct {
	model.Params
	// 入参json格式
	_publishJson string
}

// NewYunosOsupdateOsfotaPublishRequest 初始化YunosOsupdateOsfotaPublishAPIRequest对象
func NewYunosOsupdateOsfotaPublishRequest() *YunosOsupdateOsfotaPublishAPIRequest {
	return &YunosOsupdateOsfotaPublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaPublishAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaPublishAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPublishJson is PublishJson Setter
// 入参json格式
func (r *YunosOsupdateOsfotaPublishAPIRequest) SetPublishJson(_publishJson string) error {
	r._publishJson = _publishJson
	r.Set("publish_json", _publishJson)
	return nil
}

// GetPublishJson PublishJson Getter
func (r YunosOsupdateOsfotaPublishAPIRequest) GetPublishJson() string {
	return r._publishJson
}
