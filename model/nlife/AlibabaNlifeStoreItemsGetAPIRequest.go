package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeStoreItemsGetAPIRequest 获取门店的商品列表(在售|已下架|全部) API请求
// alibaba.nlife.store.items.get
//
// 利用该接口可以获取到零售+商品服务中符合条件的商品列表，包括在售的、已下架的或者是所有状态的商品。
type AlibabaNlifeStoreItemsGetAPIRequest struct {
	model.Params
	// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
	_storeIdType string
	// 门店ID/设备号
	_storeId string
	// 商品类目ID
	_cid int64
	// 品牌ID
	_brandId int64
	// 商品状态: ON_SALE-在售 ; OFF_SALE-已下架 ; ALL-全部
	_status string
	// 商品类型: STORE_GOODS-经销/现货 ; SUPPLIER_GOODS-代销/网直供 ; TAOKE-淘宝客 ; ALL-全部商品
	_type string
	// 商品名称(支持模糊查询)
	_title string
	// 查询开始时间
	_startModified string
	// 查询结束时间
	_endModified string
	// 分页的页码
	_pageNo int64
	// 分页时每页的数量
	_pageSize int64
	// 商品的来源：0-从零售加采购的商品；1-商户线下导入的商品
	_itemType int64
}

// NewAlibabaNlifeStoreItemsGetRequest 初始化AlibabaNlifeStoreItemsGetAPIRequest对象
func NewAlibabaNlifeStoreItemsGetRequest() *AlibabaNlifeStoreItemsGetAPIRequest {
	return &AlibabaNlifeStoreItemsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.nlife.store.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreIdType Setter
// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetStoreIdType(_storeIdType string) error {
	r._storeIdType = _storeIdType
	r.Set("store_id_type", _storeIdType)
	return nil
}

// Get StoreIdType Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetStoreIdType() string {
	return r._storeIdType
}

// Set is StoreId Setter
// 门店ID/设备号
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is Cid Setter
// 商品类目ID
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// Get Cid Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetCid() int64 {
	return r._cid
}

// Set is BrandId Setter
// 品牌ID
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// Get BrandId Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// Set is Status Setter
// 商品状态: ON_SALE-在售 ; OFF_SALE-已下架 ; ALL-全部
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetStatus() string {
	return r._status
}

// Set is Type Setter
// 商品类型: STORE_GOODS-经销/现货 ; SUPPLIER_GOODS-代销/网直供 ; TAOKE-淘宝客 ; ALL-全部商品
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetType() string {
	return r._type
}

// Set is Title Setter
// 商品名称(支持模糊查询)
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetTitle() string {
	return r._title
}

// Set is StartModified Setter
// 查询开始时间
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetStartModified(_startModified string) error {
	r._startModified = _startModified
	r.Set("start_modified", _startModified)
	return nil
}

// Get StartModified Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetStartModified() string {
	return r._startModified
}

// Set is EndModified Setter
// 查询结束时间
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetEndModified(_endModified string) error {
	r._endModified = _endModified
	r.Set("end_modified", _endModified)
	return nil
}

// Get EndModified Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetEndModified() string {
	return r._endModified
}

// Set is PageNo Setter
// 分页的页码
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 分页时每页的数量
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is ItemType Setter
// 商品的来源：0-从零售加采购的商品；1-商户线下导入的商品
func (r *AlibabaNlifeStoreItemsGetAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// Get ItemType Getter
func (r AlibabaNlifeStoreItemsGetAPIRequest) GetItemType() int64 {
	return r._itemType
}
