package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbinventorylogqueryAPIRequest 根据商品ID查询所有库存变更记录 API请求
// taobao.wlb.inventorylog.query
//
// 通过商品ID等几个条件来分页查询库存变更记录
type TaobaowlbinventorylogqueryAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
	// 单号
	_orderCode string
	// 起始修改时间,大于等于该时间
	_gmtStart string
	// 结束修改时间,小于等于该时间
	_gmtEnd string
	// 库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
	_opType string
	// 商品ID
	_itemId int64
	// 当前页
	_pageNo int64
	// 分页记录个数
	_pageSize int64
	// 可指定授权的用户来查询
	_opUserId int64
}

// NewTaobaowlbinventorylogqueryRequest 初始化TaobaowlbinventorylogqueryAPIRequest对象
func NewTaobaowlbinventorylogqueryRequest() *TaobaowlbinventorylogqueryAPIRequest {
	return &TaobaowlbinventorylogqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbinventorylogqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.inventorylog.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbinventorylogqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbinventorylogqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaowlbinventorylogqueryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetOrderCode is OrderCode Setter
// 单号
func (r *TaobaowlbinventorylogqueryAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetGmtStart is GmtStart Setter
// 起始修改时间,大于等于该时间
func (r *TaobaowlbinventorylogqueryAPIRequest) SetGmtStart(_gmtStart string) error {
	r._gmtStart = _gmtStart
	r.Set("gmt_start", _gmtStart)
	return nil
}

// GetGmtStart GmtStart Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetGmtStart() string {
	return r._gmtStart
}

// SetGmtEnd is GmtEnd Setter
// 结束修改时间,小于等于该时间
func (r *TaobaowlbinventorylogqueryAPIRequest) SetGmtEnd(_gmtEnd string) error {
	r._gmtEnd = _gmtEnd
	r.Set("gmt_end", _gmtEnd)
	return nil
}

// GetGmtEnd GmtEnd Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetGmtEnd() string {
	return r._gmtEnd
}

// SetOpType is OpType Setter
// 库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
func (r *TaobaowlbinventorylogqueryAPIRequest) SetOpType(_opType string) error {
	r._opType = _opType
	r.Set("op_type", _opType)
	return nil
}

// GetOpType OpType Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetOpType() string {
	return r._opType
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaowlbinventorylogqueryAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetPageNo is PageNo Setter
// 当前页
func (r *TaobaowlbinventorylogqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页记录个数
func (r *TaobaowlbinventorylogqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOpUserId is OpUserId Setter
// 可指定授权的用户来查询
func (r *TaobaowlbinventorylogqueryAPIRequest) SetOpUserId(_opUserId int64) error {
	r._opUserId = _opUserId
	r.Set("op_user_id", _opUserId)
	return nil
}

// GetOpUserId OpUserId Getter
func (r TaobaowlbinventorylogqueryAPIRequest) GetOpUserId() int64 {
	return r._opUserId
}
