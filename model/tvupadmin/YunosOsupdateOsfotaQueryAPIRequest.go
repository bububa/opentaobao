package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateOsfotaQueryAPIRequest 系统升级分页查询 API请求
// yunos.osupdate.osfota.query
//
// 分页查询osoupdate系统升级列表
type YunosOsupdateOsfotaQueryAPIRequest struct {
	model.Params
	// 设备型号ID
	_modleId int64
	// 页码
	_page int64
	// 每页数量
	_pageSize int64
}

// NewYunosOsupdateOsfotaQueryRequest 初始化YunosOsupdateOsfotaQueryAPIRequest对象
func NewYunosOsupdateOsfotaQueryRequest() *YunosOsupdateOsfotaQueryAPIRequest {
	return &YunosOsupdateOsfotaQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateOsfotaQueryAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.osfota.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateOsfotaQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetModleId is ModleId Setter
// 设备型号ID
func (r *YunosOsupdateOsfotaQueryAPIRequest) SetModleId(_modleId int64) error {
	r._modleId = _modleId
	r.Set("modle_id", _modleId)
	return nil
}

// GetModleId ModleId Getter
func (r YunosOsupdateOsfotaQueryAPIRequest) GetModleId() int64 {
	return r._modleId
}

// SetPage is Page Setter
// 页码
func (r *YunosOsupdateOsfotaQueryAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r YunosOsupdateOsfotaQueryAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页数量
func (r *YunosOsupdateOsfotaQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosOsupdateOsfotaQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
