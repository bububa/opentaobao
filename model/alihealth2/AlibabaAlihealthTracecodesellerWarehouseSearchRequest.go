package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库api APIRequest
alibaba.alihealth.tracecodeseller.warehouse.search

查询仓库api
*/
type AlibabaAlihealthTracecodesellerWarehouseSearchRequest struct {
    model.Params

    // 身份认证
    appkey   string 

    // 商家id
    entInfoId   int64 

    // 第几页
    page   int64 

    // 每页多少条
    pageSize   int64 

}

func NewAlibabaAlihealthTracecodesellerWarehouseSearchRequest() *AlibabaAlihealthTracecodesellerWarehouseSearchRequest{
    return &AlibabaAlihealthTracecodesellerWarehouseSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.warehouse.search"
}

func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetAppkey(appkey string) error {
    r.appkey = appkey
    r.Set("appkey", appkey)
    return nil
}

func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetAppkey() string {
    return r.appkey
}

func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetEntInfoId() int64 {
    return r.entInfoId
}

func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthTracecodesellerWarehouseSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthTracecodesellerWarehouseSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

