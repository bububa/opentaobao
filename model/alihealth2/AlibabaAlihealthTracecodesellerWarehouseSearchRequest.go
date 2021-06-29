package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库api API请求
alibaba.alihealth.tracecodeseller.warehouse.search

查询仓库api
*/
type AlibabaAlihealthTracecodesellerWarehouseSearchRequest struct {
    model.Params
    // 身份认证
    _appkey   string
    // 商家id
    _entInfoId   int64
    // 第几页
    _page   int64
    // 每页多少条
    _pageSize   int64
}

// 初始化AlibabaAlihealthTracecodesellerWarehouseSearchRequest对象
func NewAlibabaAlihealthTracecodesellerWarehouseSearchRequest() *AlibabaAlihealthTracecodesellerWarehouseSearchRequest{
    return &AlibabaAlihealthTracecodesellerWarehouseSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.warehouse.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Appkey Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetAppkey(_appkey string) error {
    r._appkey = _appkey
    r.Set("appkey", _appkey)
    return nil
}

// Appkey Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetAppkey() string {
    return r._appkey
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
// Page Setter
// 第几页
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页多少条
func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetPageSize() int64 {
    return r._pageSize
}
