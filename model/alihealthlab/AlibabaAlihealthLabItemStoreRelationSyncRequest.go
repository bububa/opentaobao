package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
检验检测业务，isv项目门店关系同步 APIRequest
alibaba.alihealth.lab.item.store.relation.sync

阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
*/
type AlibabaAlihealthLabItemStoreRelationSyncRequest struct {
    model.Params

    // EFFECTIVE 有效，INVALID 无效
    isvRelationStatus   string 

    // isv门店编码
    isvStoreCodes   []string 

    // 检验检测项目isv侧编码
    isvItemCode   string 

}

func NewAlibabaAlihealthLabItemStoreRelationSyncRequest() *AlibabaAlihealthLabItemStoreRelationSyncRequest{
    return &AlibabaAlihealthLabItemStoreRelationSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.item.store.relation.sync"
}

func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthLabItemStoreRelationSyncRequest) SetIsvRelationStatus(isvRelationStatus string) error {
    r.isvRelationStatus = isvRelationStatus
    r.Set("isv_relation_status", isvRelationStatus)
    return nil
}

func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetIsvRelationStatus() string {
    return r.isvRelationStatus
}

func (r *AlibabaAlihealthLabItemStoreRelationSyncRequest) SetIsvStoreCodes(isvStoreCodes []string) error {
    r.isvStoreCodes = isvStoreCodes
    r.Set("isv_store_codes", isvStoreCodes)
    return nil
}

func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetIsvStoreCodes() []string {
    return r.isvStoreCodes
}

func (r *AlibabaAlihealthLabItemStoreRelationSyncRequest) SetIsvItemCode(isvItemCode string) error {
    r.isvItemCode = isvItemCode
    r.Set("isv_item_code", isvItemCode)
    return nil
}

func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetIsvItemCode() string {
    return r.isvItemCode
}

