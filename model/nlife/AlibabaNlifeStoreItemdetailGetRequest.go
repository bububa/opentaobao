package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的详情信息 API请求
alibaba.nlife.store.itemdetail.get

查询零售加平台上单个商品的详情信息
*/
type AlibabaNlifeStoreItemdetailGetRequest struct {
    model.Params
    // 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
    storeIdType   string
    // 门店ID
    storeId   string
    // 商品在外部商家的编码(与item_id不能同时为空)
    outerId   string
    // 商品Item的ID(与outer_id不能同时为空)
    itemId   int64
    // skuId列表-可查询指定的sku
    skuIdList   []int64
    // 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
    itemType   *model.File
    // 商家对商品的自用编码
    code   string
}

// 初始化AlibabaNlifeStoreItemdetailGetRequest对象
func NewAlibabaNlifeStoreItemdetailGetRequest() *AlibabaNlifeStoreItemdetailGetRequest{
    return &AlibabaNlifeStoreItemdetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreItemdetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.itemdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreItemdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreIdType Setter
// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetStoreIdType(storeIdType string) error {
    r.storeIdType = storeIdType
    r.Set("store_id_type", storeIdType)
    return nil
}

// StoreIdType Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetStoreIdType() string {
    return r.storeIdType
}
// StoreId Setter
// 门店ID
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetStoreId() string {
    return r.storeId
}
// OuterId Setter
// 商品在外部商家的编码(与item_id不能同时为空)
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetOuterId() string {
    return r.outerId
}
// ItemId Setter
// 商品Item的ID(与outer_id不能同时为空)
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetItemId() int64 {
    return r.itemId
}
// SkuIdList Setter
// skuId列表-可查询指定的sku
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetSkuIdList(skuIdList []int64) error {
    r.skuIdList = skuIdList
    r.Set("sku_id_list", skuIdList)
    return nil
}

// SkuIdList Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetSkuIdList() []int64 {
    return r.skuIdList
}
// ItemType Setter
// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetItemType(itemType *model.File) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetItemType() *model.File {
    return r.itemType
}
// Code Setter
// 商家对商品的自用编码
func (r *AlibabaNlifeStoreItemdetailGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaNlifeStoreItemdetailGetRequest) GetCode() string {
    return r.code
}
