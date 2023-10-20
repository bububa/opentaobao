package tvupadmin

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateAppversionPublishAPIRequest) Reset() {
	r._publishJson = ""
	r.Params.ToZero()
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

var poolYunosOsupdateAppversionPublishAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateAppversionPublishRequest()
	},
}

// GetYunosOsupdateAppversionPublishRequest 从 sync.Pool 获取 YunosOsupdateAppversionPublishAPIRequest
func GetYunosOsupdateAppversionPublishAPIRequest() *YunosOsupdateAppversionPublishAPIRequest {
	return poolYunosOsupdateAppversionPublishAPIRequest.Get().(*YunosOsupdateAppversionPublishAPIRequest)
}

// ReleaseYunosOsupdateAppversionPublishAPIRequest 将 YunosOsupdateAppversionPublishAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateAppversionPublishAPIRequest(v *YunosOsupdateAppversionPublishAPIRequest) {
	v.Reset()
	poolYunosOsupdateAppversionPublishAPIRequest.Put(v)
}
