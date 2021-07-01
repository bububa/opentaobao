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
type AlibabaNlifeStoreItemdetailsGetAPIRequest struct {
    model.Params
    // 门店ID/设备号
    _storeId   string
    // 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ; 门店设备号-STORE_DEVICE
    _storeIdType   string
    // 查询参数list
    _itemQueryDOList   []ItemQueryDOList
    // 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
    _itemType   *model.File
}

// 初始化AlibabaNlifeStoreItemdetailsGetAPIRequest对象
func NewAlibabaNlifeStoreItemdetailsGetRequest() *AlibabaNlifeStoreItemdetailsGetAPIRequest{
    return &AlibabaNlifeStoreItemdetailsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreItemdetailsGetAPIRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.itemdetails.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreItemdetailsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID/设备号
func (r *AlibabaNlifeStoreItemdetailsGetAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreItemdetailsGetAPIRequest) GetStoreId() string {
    return r._storeId
}
// StoreIdType Setter
// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ; 门店设备号-STORE_DEVICE
func (r *AlibabaNlifeStoreItemdetailsGetAPIRequest) SetStoreIdType(_storeIdType string) error {
    r._storeIdType = _storeIdType
    r.Set("store_id_type", _storeIdType)
    return nil
}

// StoreIdType Getter
func (r AlibabaNlifeStoreItemdetailsGetAPIRequest) GetStoreIdType() string {
    return r._storeIdType
}
// ItemQueryDOList Setter
// 查询参数list
func (r *AlibabaNlifeStoreItemdetailsGetAPIRequest) SetItemQueryDOList(_itemQueryDOList []ItemQueryDOList) error {
    r._itemQueryDOList = _itemQueryDOList
    r.Set("item_query_d_o_list", _itemQueryDOList)
    return nil
}

// ItemQueryDOList Getter
func (r AlibabaNlifeStoreItemdetailsGetAPIRequest) GetItemQueryDOList() []ItemQueryDOList {
    return r._itemQueryDOList
}
// ItemType Setter
// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
func (r *AlibabaNlifeStoreItemdetailsGetAPIRequest) SetItemType(_itemType *model.File) error {
    r._itemType = _itemType
    r.Set("item_type", _itemType)
    return nil
}

// ItemType Getter
func (r AlibabaNlifeStoreItemdetailsGetAPIRequest) GetItemType() *model.File {
    return r._itemType
}
