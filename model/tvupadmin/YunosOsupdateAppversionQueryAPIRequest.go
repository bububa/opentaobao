package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionQueryAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
