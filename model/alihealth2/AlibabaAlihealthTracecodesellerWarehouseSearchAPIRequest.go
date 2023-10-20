package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthtracecodesellerwarehousesearchAPIRequest 查询仓库api API请求
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
type AlibabaalihealthtracecodesellerwarehousesearchAPIRequest struct {
	model.Params
	// 身份认证
	_appkey string
	// 商家id
	_entInfoId int64
	// 第几页
	_page int64
	// 每页多少条
	_pageSize int64
}

// NewAlibabaalihealthtracecodesellerwarehousesearchRequest 初始化AlibabaalihealthtracecodesellerwarehousesearchAPIRequest对象
func NewAlibabaalihealthtracecodesellerwarehousesearchRequest() *AlibabaalihealthtracecodesellerwarehousesearchAPIRequest {
	return &AlibabaalihealthtracecodesellerwarehousesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.warehouse.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppkey is Appkey Setter
// 身份认证
func (r *AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) SetAppkey(_appkey string) error {
	r._appkey = _appkey
	r.Set("appkey", _appkey)
	return nil
}

// GetAppkey Appkey Getter
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetAppkey() string {
	return r._appkey
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetPage is Page Setter
// 第几页
func (r *AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页多少条
func (r *AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthtracecodesellerwarehousesearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
