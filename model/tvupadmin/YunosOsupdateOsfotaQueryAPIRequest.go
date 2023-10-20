package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateosfotaqueryAPIRequest 系统升级分页查询 API请求
// yunos.osupdate.osfota.query
//
// 分页查询osoupdate系统升级列表
type YunososupdateosfotaqueryAPIRequest struct {
	model.Params
	// 设备型号ID
	_modleId int64
	// 页码
	_page int64
	// 每页数量
	_pageSize int64
}

// NewYunososupdateosfotaqueryRequest 初始化YunososupdateosfotaqueryAPIRequest对象
func NewYunososupdateosfotaqueryRequest() *YunososupdateosfotaqueryAPIRequest {
	return &YunososupdateosfotaqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunososupdateosfotaqueryAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunososupdateosfotaqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunososupdateosfotaqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModleId is ModleId Setter
// 设备型号ID
func (r *YunososupdateosfotaqueryAPIRequest) SetModleId(_modleId int64) error {
	r._modleId = _modleId
	r.Set("modle_id", _modleId)
	return nil
}

// GetModleId ModleId Getter
func (r YunososupdateosfotaqueryAPIRequest) GetModleId() int64 {
	return r._modleId
}

// SetPage is Page Setter
// 页码
func (r *YunososupdateosfotaqueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r YunososupdateosfotaqueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页数量
func (r *YunososupdateosfotaqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunososupdateosfotaqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
