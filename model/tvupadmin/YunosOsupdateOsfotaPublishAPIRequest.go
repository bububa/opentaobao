package tvupadmin

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateOsfotaPublishAPIRequest) Reset() {
	r._publishJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaPublishAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaPublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateOsfotaPublishAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolYunosOsupdateOsfotaPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateOsfotaPublishRequest()
	},
}

// GetYunosOsupdateOsfotaPublishRequest 从 sync.Pool 获取 YunosOsupdateOsfotaPublishAPIRequest
func GetYunosOsupdateOsfotaPublishAPIRequest() *YunosOsupdateOsfotaPublishAPIRequest {
	return poolYunosOsupdateOsfotaPublishAPIRequest.Get().(*YunosOsupdateOsfotaPublishAPIRequest)
}

// ReleaseYunosOsupdateOsfotaPublishAPIRequest 将 YunosOsupdateOsfotaPublishAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateOsfotaPublishAPIRequest(v *YunosOsupdateOsfotaPublishAPIRequest) {
	v.Reset()
	poolYunosOsupdateOsfotaPublishAPIRequest.Put(v)
}
