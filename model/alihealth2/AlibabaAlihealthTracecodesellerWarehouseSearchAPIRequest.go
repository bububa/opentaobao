package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest
查询仓库api API请求
alibaba.alihealth.tracecodeseller.warehouse.search

查询仓库api */
type AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest struct {
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

// NewAlibabaAlihealthTracecodesellerWarehouseSearchRequest 初始化AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerWarehouseSearchRequest() *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest {
	return &AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.warehouse.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Appkey Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetAppkey(_appkey string) error {
	r._appkey = _appkey
	r.Set("appkey", _appkey)
	return nil
}

// Get Appkey Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetAppkey() string {
	return r._appkey
}

// Set is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// Get EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// Set is Page Setter
// 第几页
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetPage() int64 {
	return r._page
}

// Set is PageSize Setter
// 每页多少条
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
