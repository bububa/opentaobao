package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
检验检测业务，isv项目门店关系同步 API请求
alibaba.alihealth.lab.item.store.relation.sync

阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
*/
type AlibabaAlihealthLabItemStoreRelationSyncRequest struct {
    model.Params
    // EFFECTIVE 有效，INVALID 无效
    _isvRelationStatus   string
    // isv门店编码
    _isvStoreCodes   []string
    // 检验检测项目isv侧编码
    _isvItemCode   string
}

// 初始化AlibabaAlihealthLabItemStoreRelationSyncRequest对象
func NewAlibabaAlihealthLabItemStoreRelationSyncRequest() *AlibabaAlihealthLabItemStoreRelationSyncRequest{
    return &AlibabaAlihealthLabItemStoreRelationSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.item.store.relation.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvRelationStatus Setter
// EFFECTIVE 有效，INVALID 无效
func (r *AlibabaAlihealthLabItemStoreRelationSyncRequest) SetIsvRelationStatus(_isvRelationStatus string) error {
    r._isvRelationStatus = _isvRelationStatus
    r.Set("isv_relation_status", _isvRelationStatus)
    return nil
}

// IsvRelationStatus Getter
func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetIsvRelationStatus() string {
    return r._isvRelationStatus
}
// IsvStoreCodes Setter
// isv门店编码
func (r *AlibabaAlihealthLabItemStoreRelationSyncRequest) SetIsvStoreCodes(_isvStoreCodes []string) error {
    r._isvStoreCodes = _isvStoreCodes
    r.Set("isv_store_codes", _isvStoreCodes)
    return nil
}

// IsvStoreCodes Getter
func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetIsvStoreCodes() []string {
    return r._isvStoreCodes
}
// IsvItemCode Setter
// 检验检测项目isv侧编码
func (r *AlibabaAlihealthLabItemStoreRelationSyncRequest) SetIsvItemCode(_isvItemCode string) error {
    r._isvItemCode = _isvItemCode
    r.Set("isv_item_code", _isvItemCode)
    return nil
}

// IsvItemCode Getter
func (r AlibabaAlihealthLabItemStoreRelationSyncRequest) GetIsvItemCode() string {
    return r._isvItemCode
}
