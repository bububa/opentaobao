package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateappversionqueryAPIRequest 分页获取桌面升级任务 API请求
// yunos.osupdate.appversion.query
//
// 分页获取桌面升级任务
type YunososupdateappversionqueryAPIRequest struct {
	model.Params
	// 应用ID
	_appId int64
	// 页码值
	_page int64
	// 页大小
	_size int64
}

// NewYunososupdateappversionqueryRequest 初始化YunososupdateappversionqueryAPIRequest对象
func NewYunososupdateappversionqueryRequest() *YunososupdateappversionqueryAPIRequest {
	return &YunososupdateappversionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateappversionqueryAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateappversionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateappversionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppId is AppId Setter
// 应用ID
func (r *YunososupdateappversionqueryAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r YunososupdateappversionqueryAPIRequest) GetAppId() int64 {
	return r._appId
}

// SetPage is Page Setter
// 页码值
func (r *YunososupdateappversionqueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r YunososupdateappversionqueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetSize is Size Setter
// 页大小
func (r *YunososupdateappversionqueryAPIRequest) SetSize(_size int64) error {
	r._size = _size
	r.Set("size", _size)
	return nil
}

// GetSize Size Getter
func (r YunososupdateappversionqueryAPIRequest) GetSize() int64 {
	return r._size
}
