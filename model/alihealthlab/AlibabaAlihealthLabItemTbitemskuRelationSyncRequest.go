package alihealthlab

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 API请求
alibaba.alihealth.lab.item.tbitemsku.relation.sync

阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步
*/
type AlibabaAlihealthLabItemTbitemskuRelationSyncRequest struct {
    model.Params
    // EFFECTIVE 有效，INVALID 无效
    isvRelationStatus   string
    // 关联的淘宝商品SKU id，在商品没有sku的情况下传0
    tbSkuId   int64
    // 关联的淘宝商品 id
    tbItemId   int64
    // 检验检测项目isv侧code
    isvItemCode   string
}

// 初始化AlibabaAlihealthLabItemTbitemskuRelationSyncRequest对象
func NewAlibabaAlihealthLabItemTbitemskuRelationSyncRequest() *AlibabaAlihealthLabItemTbitemskuRelationSyncRequest{
    return &AlibabaAlihealthLabItemTbitemskuRelationSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.lab.item.tbitemsku.relation.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvRelationStatus Setter
// EFFECTIVE 有效，INVALID 无效
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) SetIsvRelationStatus(isvRelationStatus string) error {
    r.isvRelationStatus = isvRelationStatus
    r.Set("isv_relation_status", isvRelationStatus)
    return nil
}

// IsvRelationStatus Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) GetIsvRelationStatus() string {
    return r.isvRelationStatus
}
// TbSkuId Setter
// 关联的淘宝商品SKU id，在商品没有sku的情况下传0
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) SetTbSkuId(tbSkuId int64) error {
    r.tbSkuId = tbSkuId
    r.Set("tb_sku_id", tbSkuId)
    return nil
}

// TbSkuId Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) GetTbSkuId() int64 {
    return r.tbSkuId
}
// TbItemId Setter
// 关联的淘宝商品 id
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) SetTbItemId(tbItemId int64) error {
    r.tbItemId = tbItemId
    r.Set("tb_item_id", tbItemId)
    return nil
}

// TbItemId Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) GetTbItemId() int64 {
    return r.tbItemId
}
// IsvItemCode Setter
// 检验检测项目isv侧code
func (r *AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) SetIsvItemCode(isvItemCode string) error {
    r.isvItemCode = isvItemCode
    r.Set("isv_item_code", isvItemCode)
    return nil
}

// IsvItemCode Getter
func (r AlibabaAlihealthLabItemTbitemskuRelationSyncRequest) GetIsvItemCode() string {
    return r.isvItemCode
}
