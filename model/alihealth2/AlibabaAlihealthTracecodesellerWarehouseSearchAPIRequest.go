package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest 查询仓库api API请求
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) Reset() {
	r._appkey = ""
	r._entInfoId = 0
	r._page = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.warehouse.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppkey is Appkey Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetAppkey(_appkey string) error {
	r._appkey = _appkey
	r.Set("appkey", _appkey)
	return nil
}

// GetAppkey Appkey Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetAppkey() string {
	return r._appkey
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetPage is Page Setter
// 第几页
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页多少条
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesellerWarehouseSearchRequest()
	},
}

// GetAlibabaAlihealthTracecodesellerWarehouseSearchRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest
func GetAlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest() *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest {
	return poolAlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest.Get().(*AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest 将 AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest(v *AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest.Put(v)
}
