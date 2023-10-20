package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionQueryAPIRequest 分页获取桌面升级任务 API请求
// yunos.osupdate.appversion.query
//
// 分页获取桌面升级任务
type YunosOsupdateAppversionQueryAPIRequest struct {
	model.Params
	// 应用ID
	_appId int64
	// 页码值
	_page int64
	// 页大小
	_size int64
}

// NewYunosOsupdateAppversionQueryRequest 初始化YunosOsupdateAppversionQueryAPIRequest对象
func NewYunosOsupdateAppversionQueryRequest() *YunosOsupdateAppversionQueryAPIRequest {
	return &YunosOsupdateAppversionQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateAppversionQueryAPIRequest) Reset() {
	r._appId = 0
	r._page = 0
	r._size = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionQueryAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateAppversionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 应用ID
func (r *YunosOsupdateAppversionQueryAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r YunosOsupdateAppversionQueryAPIRequest) GetAppId() int64 {
	return r._appId
}

// SetPage is Page Setter
// 页码值
func (r *YunosOsupdateAppversionQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r YunosOsupdateAppversionQueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetSize is Size Setter
// 页大小
func (r *YunosOsupdateAppversionQueryAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r YunosOsupdateAppversionQueryAPIRequest) GetSize() int64 {
	return r._size
}

var poolYunosOsupdateAppversionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateAppversionQueryRequest()
	},
}

// GetYunosOsupdateAppversionQueryRequest 从 sync.Pool 获取 YunosOsupdateAppversionQueryAPIRequest
func GetYunosOsupdateAppversionQueryAPIRequest() *YunosOsupdateAppversionQueryAPIRequest {
	return poolYunosOsupdateAppversionQueryAPIRequest.Get().(*YunosOsupdateAppversionQueryAPIRequest)
}

// ReleaseYunosOsupdateAppversionQueryAPIRequest 将 YunosOsupdateAppversionQueryAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateAppversionQueryAPIRequest(v *YunosOsupdateAppversionQueryAPIRequest) {
	v.Reset()
	poolYunosOsupdateAppversionQueryAPIRequest.Put(v)
}
