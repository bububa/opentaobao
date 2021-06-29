package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店的商品列表(在售|已下架|全部) APIRequest
alibaba.nlife.store.items.get

利用该接口可以获取到零售+商品服务中符合条件的商品列表，包括在售的、已下架的或者是所有状态的商品。
*/
type AlibabaNlifeStoreItemsGetRequest struct {
    model.Params

    // 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
    storeIdType   string 

    // 门店ID/设备号
    storeId   string 

    // 商品类目ID
    cid   int64 

    // 品牌ID
    brandId   int64 

    // 商品状态: ON_SALE-在售 ; OFF_SALE-已下架 ; ALL-全部
    status   string 

    // 商品类型: STORE_GOODS-经销/现货 ; SUPPLIER_GOODS-代销/网直供 ; TAOKE-淘宝客 ; ALL-全部商品
    type   string 

    // 商品名称(支持模糊查询)
    title   string 

    // 查询开始时间
    startModified   string 

    // 查询结束时间
    endModified   string 

    // 分页的页码
    pageNo   int64 

    // 分页时每页的数量
    pageSize   int64 

    // 商品的来源：0-从零售加采购的商品；1-商户线下导入的商品
    itemType   int64 

}

func NewAlibabaNlifeStoreItemsGetRequest() *AlibabaNlifeStoreItemsGetRequest{
    return &AlibabaNlifeStoreItemsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeStoreItemsGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.items.get"
}

func (r AlibabaNlifeStoreItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeStoreItemsGetRequest) SetStoreIdType(storeIdType string) error {
    r.storeIdType = storeIdType
    r.Set("store_id_type", storeIdType)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetStoreIdType() string {
    return r.storeIdType
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetCid() int64 {
    return r.cid
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetBrandId(brandId int64) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetBrandId() int64 {
    return r.brandId
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetType() string {
    return r.type
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetTitle() string {
    return r.title
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetStartModified(startModified string) error {
    r.startModified = startModified
    r.Set("start_modified", startModified)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetStartModified() string {
    return r.startModified
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetEndModified(endModified string) error {
    r.endModified = endModified
    r.Set("end_modified", endModified)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetEndModified() string {
    return r.endModified
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaNlifeStoreItemsGetRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

func (r AlibabaNlifeStoreItemsGetRequest) GetItemType() int64 {
    return r.itemType
}

