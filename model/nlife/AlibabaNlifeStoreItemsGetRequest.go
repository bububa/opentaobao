package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店的商品列表(在售|已下架|全部) API请求
alibaba.nlife.store.items.get

利用该接口可以获取到零售+商品服务中符合条件的商品列表，包括在售的、已下架的或者是所有状态的商品。
*/
type AlibabaNlifeStoreItemsGetRequest struct {
    model.Params
    // 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
    _storeIdType   string
    // 门店ID/设备号
    _storeId   string
    // 商品类目ID
    _cid   int64
    // 品牌ID
    _brandId   int64
    // 商品状态: ON_SALE-在售 ; OFF_SALE-已下架 ; ALL-全部
    _status   string
    // 商品类型: STORE_GOODS-经销/现货 ; SUPPLIER_GOODS-代销/网直供 ; TAOKE-淘宝客 ; ALL-全部商品
    _type   string
    // 商品名称(支持模糊查询)
    _title   string
    // 查询开始时间
    _startModified   string
    // 查询结束时间
    _endModified   string
    // 分页的页码
    _pageNo   int64
    // 分页时每页的数量
    _pageSize   int64
    // 商品的来源：0-从零售加采购的商品；1-商户线下导入的商品
    _itemType   int64
}

// 初始化AlibabaNlifeStoreItemsGetRequest对象
func NewAlibabaNlifeStoreItemsGetRequest() *AlibabaNlifeStoreItemsGetRequest{
    return &AlibabaNlifeStoreItemsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreItemsGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.items.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreIdType Setter
// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
func (r *AlibabaNlifeStoreItemsGetRequest) SetStoreIdType(_storeIdType string) error {
    r._storeIdType = _storeIdType
    r.Set("store_id_type", _storeIdType)
    return nil
}

// StoreIdType Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetStoreIdType() string {
    return r._storeIdType
}
// StoreId Setter
// 门店ID/设备号
func (r *AlibabaNlifeStoreItemsGetRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetStoreId() string {
    return r._storeId
}
// Cid Setter
// 商品类目ID
func (r *AlibabaNlifeStoreItemsGetRequest) SetCid(_cid int64) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetCid() int64 {
    return r._cid
}
// BrandId Setter
// 品牌ID
func (r *AlibabaNlifeStoreItemsGetRequest) SetBrandId(_brandId int64) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetBrandId() int64 {
    return r._brandId
}
// Status Setter
// 商品状态: ON_SALE-在售 ; OFF_SALE-已下架 ; ALL-全部
func (r *AlibabaNlifeStoreItemsGetRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetStatus() string {
    return r._status
}
// Type Setter
// 商品类型: STORE_GOODS-经销/现货 ; SUPPLIER_GOODS-代销/网直供 ; TAOKE-淘宝客 ; ALL-全部商品
func (r *AlibabaNlifeStoreItemsGetRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetType() string {
    return r._type
}
// Title Setter
// 商品名称(支持模糊查询)
func (r *AlibabaNlifeStoreItemsGetRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetTitle() string {
    return r._title
}
// StartModified Setter
// 查询开始时间
func (r *AlibabaNlifeStoreItemsGetRequest) SetStartModified(_startModified string) error {
    r._startModified = _startModified
    r.Set("start_modified", _startModified)
    return nil
}

// StartModified Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetStartModified() string {
    return r._startModified
}
// EndModified Setter
// 查询结束时间
func (r *AlibabaNlifeStoreItemsGetRequest) SetEndModified(_endModified string) error {
    r._endModified = _endModified
    r.Set("end_modified", _endModified)
    return nil
}

// EndModified Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetEndModified() string {
    return r._endModified
}
// PageNo Setter
// 分页的页码
func (r *AlibabaNlifeStoreItemsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 分页时每页的数量
func (r *AlibabaNlifeStoreItemsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// ItemType Setter
// 商品的来源：0-从零售加采购的商品；1-商户线下导入的商品
func (r *AlibabaNlifeStoreItemsGetRequest) SetItemType(_itemType int64) error {
    r._itemType = _itemType
    r.Set("item_type", _itemType)
    return nil
}

// ItemType Getter
func (r AlibabaNlifeStoreItemsGetRequest) GetItemType() int64 {
    return r._itemType
}
