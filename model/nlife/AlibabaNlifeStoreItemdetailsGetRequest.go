package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取零售加商品详情信息 API请求
alibaba.nlife.store.itemdetails.get

批量获取零售加平台上的商品详情信息
*/
type AlibabaNlifeStoreItemdetailsGetRequest struct {
    model.Params
    // 门店ID/设备号
    storeId   string
    // 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ; 门店设备号-STORE_DEVICE
    storeIdType   string
    // 查询参数list
    itemQueryDOList   []ItemQueryDOList
    // 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
    itemType   *model.File
}

// 初始化AlibabaNlifeStoreItemdetailsGetRequest对象
func NewAlibabaNlifeStoreItemdetailsGetRequest() *AlibabaNlifeStoreItemdetailsGetRequest{
    return &AlibabaNlifeStoreItemdetailsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreItemdetailsGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.itemdetails.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreItemdetailsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID/设备号
func (r *AlibabaNlifeStoreItemdetailsGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreItemdetailsGetRequest) GetStoreId() string {
    return r.storeId
}
// StoreIdType Setter
// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ; 门店设备号-STORE_DEVICE
func (r *AlibabaNlifeStoreItemdetailsGetRequest) SetStoreIdType(storeIdType string) error {
    r.storeIdType = storeIdType
    r.Set("store_id_type", storeIdType)
    return nil
}

// StoreIdType Getter
func (r AlibabaNlifeStoreItemdetailsGetRequest) GetStoreIdType() string {
    return r.storeIdType
}
// ItemQueryDOList Setter
// 查询参数list
func (r *AlibabaNlifeStoreItemdetailsGetRequest) SetItemQueryDOList(itemQueryDOList []ItemQueryDOList) error {
    r.itemQueryDOList = itemQueryDOList
    r.Set("item_query_d_o_list", itemQueryDOList)
    return nil
}

// ItemQueryDOList Getter
func (r AlibabaNlifeStoreItemdetailsGetRequest) GetItemQueryDOList() []ItemQueryDOList {
    return r.itemQueryDOList
}
// ItemType Setter
// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
func (r *AlibabaNlifeStoreItemdetailsGetRequest) SetItemType(itemType *model.File) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r AlibabaNlifeStoreItemdetailsGetRequest) GetItemType() *model.File {
    return r.itemType
}
