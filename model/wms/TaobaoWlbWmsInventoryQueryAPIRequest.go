package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsinventoryqueryAPIRequest 菜鸟商品库存查询 API请求
// taobao.wlb.wms.inventory.query
//
// 支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存
type TaobaowlbwmsinventoryqueryAPIRequest struct {
	model.Params
	// 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
	_channelCode string
	// 失效日期，type=2时字段传值有效。
	_dueDate string
	// 生产日期，type=2时字段传值有效。
	_produceDate string
	// 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
	_batchCode string
	// 仓库编码
	_storeCode string
	// 菜鸟商品ID
	_itemId string
	// 每页多少条，最大50条
	_pageSize int64
	// 分页的第几页
	_pageNo int64
	// 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
	_type int64
	// 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
	_inventoryType int64
}

// NewTaobaowlbwmsinventoryqueryRequest 初始化TaobaowlbwmsinventoryqueryAPIRequest对象
func NewTaobaowlbwmsinventoryqueryRequest() *TaobaowlbwmsinventoryqueryAPIRequest {
	return &TaobaowlbwmsinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCode is ChannelCode Setter
// 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetChannelCode() string {
	return r._channelCode
}

// SetDueDate is DueDate Setter
// 失效日期，type=2时字段传值有效。
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetDueDate(_dueDate string) error {
	r._dueDate = _dueDate
	r.Set("due_date", _dueDate)
	return nil
}

// GetDueDate DueDate Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetDueDate() string {
	return r._dueDate
}

// SetProduceDate is ProduceDate Setter
// 生产日期，type=2时字段传值有效。
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetProduceDate(_produceDate string) error {
	r._produceDate = _produceDate
	r.Set("produce_date", _produceDate)
	return nil
}

// GetProduceDate ProduceDate Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetProduceDate() string {
	return r._produceDate
}

// SetBatchCode is BatchCode Setter
// 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetBatchCode(_batchCode string) error {
	r._batchCode = _batchCode
	r.Set("batch_code", _batchCode)
	return nil
}

// GetBatchCode BatchCode Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetBatchCode() string {
	return r._batchCode
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetItemId is ItemId Setter
// 菜鸟商品ID
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetItemId() string {
	return r._itemId
}

// SetPageSize is PageSize Setter
// 每页多少条，最大50条
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 分页的第几页
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetType is Type Setter
// 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetType() int64 {
	return r._type
}

// SetInventoryType is InventoryType Setter
// 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
func (r *TaobaowlbwmsinventoryqueryAPIRequest) SetInventoryType(_inventoryType int64) error {
	r._inventoryType = _inventoryType
	r.Set("inventory_type", _inventoryType)
	return nil
}

// GetInventoryType InventoryType Getter
func (r TaobaowlbwmsinventoryqueryAPIRequest) GetInventoryType() int64 {
	return r._inventoryType
}
