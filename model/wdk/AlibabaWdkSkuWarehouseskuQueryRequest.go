package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品查询接口(指定商品编码) APIRequest
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询
*/
type AlibabaWdkSkuWarehouseskuQueryRequest struct {
    model.Params

    // 商品编码
    skuCodeList   []string 

    // 仓编码
    warehouseCode   string 

}

func NewAlibabaWdkSkuWarehouseskuQueryRequest() *AlibabaWdkSkuWarehouseskuQueryRequest{
    return &AlibabaWdkSkuWarehouseskuQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkSkuWarehouseskuQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.sku.warehousesku.query"
}

func (r AlibabaWdkSkuWarehouseskuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkSkuWarehouseskuQueryRequest) SetSkuCodeList(skuCodeList []string) error {
    r.skuCodeList = skuCodeList
    r.Set("sku_code_list", skuCodeList)
    return nil
}

func (r AlibabaWdkSkuWarehouseskuQueryRequest) GetSkuCodeList() []string {
    return r.skuCodeList
}

func (r *AlibabaWdkSkuWarehouseskuQueryRequest) SetWarehouseCode(warehouseCode string) error {
    r.warehouseCode = warehouseCode
    r.Set("warehouse_code", warehouseCode)
    return nil
}

func (r AlibabaWdkSkuWarehouseskuQueryRequest) GetWarehouseCode() string {
    return r.warehouseCode
}

