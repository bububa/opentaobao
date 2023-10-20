package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateModelSearchAPIRequest 机型检索 API请求
// yunos.osupdate.model.search
//
// 机型检索
type YunosOsupdateModelSearchAPIRequest struct {
	model.Params
	// 关键词
	_name string
	// 应用ID
	_appId int64
}

// NewYunosOsupdateModelSearchRequest 初始化YunosOsupdateModelSearchAPIRequest对象
func NewYunosOsupdateModelSearchRequest() *YunosOsupdateModelSearchAPIRequest {
	return &YunosOsupdateModelSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateModelSearchAPIRequest) Reset() {
	r._name = ""
	r._appId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateModelSearchAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.model.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateModelSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateModelSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 关键词
func (r *YunosOsupdateModelSearchAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r YunosOsupdateModelSearchAPIRequest) GetName() string {
	return r._name
}

// SetAppId is AppId Setter
// 应用ID
func (r *YunosOsupdateModelSearchAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r YunosOsupdateModelSearchAPIRequest) GetAppId() int64 {
	return r._appId
}

var poolYunosOsupdateModelSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateModelSearchRequest()
	},
}

// GetYunosOsupdateModelSearchRequest 从 sync.Pool 获取 YunosOsupdateModelSearchAPIRequest
func GetYunosOsupdateModelSearchAPIRequest() *YunosOsupdateModelSearchAPIRequest {
	return poolYunosOsupdateModelSearchAPIRequest.Get().(*YunosOsupdateModelSearchAPIRequest)
}

// ReleaseYunosOsupdateModelSearchAPIRequest 将 YunosOsupdateModelSearchAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateModelSearchAPIRequest(v *YunosOsupdateModelSearchAPIRequest) {
	v.Reset()
	poolYunosOsupdateModelSearchAPIRequest.Put(v)
}
