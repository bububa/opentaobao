package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的详情信息 APIRequest
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

func NewAlibabaNlifeStoreItemdetailGetRequest() *AlibabaNlifeStoreItemdetailGetRequest{
    return &AlibabaNlifeStoreItemdetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.itemdetail.get"
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeStoreItemdetailGetRequest) SetStoreIdType(storeIdType string) error {
    r.storeIdType = storeIdType
    r.Set("store_id_type", storeIdType)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetStoreIdType() string {
    return r.storeIdType
}

func (r *AlibabaNlifeStoreItemdetailGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaNlifeStoreItemdetailGetRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetOuterId() string {
    return r.outerId
}

func (r *AlibabaNlifeStoreItemdetailGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaNlifeStoreItemdetailGetRequest) SetSkuIdList(skuIdList []int64) error {
    r.skuIdList = skuIdList
    r.Set("sku_id_list", skuIdList)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetSkuIdList() []int64 {
    return r.skuIdList
}

func (r *AlibabaNlifeStoreItemdetailGetRequest) SetItemType(itemType *model.File) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetItemType() *model.File {
    return r.itemType
}

func (r *AlibabaNlifeStoreItemdetailGetRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaNlifeStoreItemdetailGetRequest) GetCode() string {
    return r.code
}

