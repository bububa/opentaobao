package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInvTurnoverQueryAPIRequest 业务库存流水查询 API请求
// taobao.inv.turnover.query
//
// 业务库存流水
type TaobaoInvTurnoverQueryAPIRequest struct {
	model.Params
	// 库存类型
	_inventoryType []string
	// 货品ID
	_scItemId string
	// 仓编码
	_storeCode string
	// 单据类型
	_activityType string
	// 开始时间
	_sdate string
	// 结束时间
	_edate string
	// 查询页
	_pageIndex string
	// 单页记录数
	_pageSize string
}

// NewTaobaoInvTurnoverQueryRequest 初始化TaobaoInvTurnoverQueryAPIRequest对象
func NewTaobaoInvTurnoverQueryRequest() *TaobaoInvTurnoverQueryAPIRequest {
	return &TaobaoInvTurnoverQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInvTurnoverQueryAPIRequest) GetApiMethodName() string {
	return "taobao.inv.turnover.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInvTurnoverQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInventoryType is InventoryType Setter
// 库存类型
func (r *TaobaoInvTurnoverQueryAPIRequest) SetInventoryType(_inventoryType []string) error {
	r._inventoryType = _inventoryType
	r.Set("inventory_type", _inventoryType)
	return nil
}

// GetInventoryType InventoryType Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetInventoryType() []string {
	return r._inventoryType
}

// SetScItemId is ScItemId Setter
// 货品ID
func (r *TaobaoInvTurnoverQueryAPIRequest) SetScItemId(_scItemId string) error {
	r._scItemId = _scItemId
	r.Set("sc_item_id", _scItemId)
	return nil
}

// GetScItemId ScItemId Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetScItemId() string {
	return r._scItemId
}

// SetStoreCode is StoreCode Setter
// 仓编码
func (r *TaobaoInvTurnoverQueryAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetActivityType is ActivityType Setter
// 单据类型
func (r *TaobaoInvTurnoverQueryAPIRequest) SetActivityType(_activityType string) error {
	r._activityType = _activityType
	r.Set("activity_type", _activityType)
	return nil
}

// GetActivityType ActivityType Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetActivityType() string {
	return r._activityType
}

// SetSdate is Sdate Setter
// 开始时间
func (r *TaobaoInvTurnoverQueryAPIRequest) SetSdate(_sdate string) error {
	r._sdate = _sdate
	r.Set("sdate", _sdate)
	return nil
}

// GetSdate Sdate Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetSdate() string {
	return r._sdate
}

// SetEdate is Edate Setter
// 结束时间
func (r *TaobaoInvTurnoverQueryAPIRequest) SetEdate(_edate string) error {
	r._edate = _edate
	r.Set("edate", _edate)
	return nil
}

// GetEdate Edate Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetEdate() string {
	return r._edate
}

// SetPageIndex is PageIndex Setter
// 查询页
func (r *TaobaoInvTurnoverQueryAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 单页记录数
func (r *TaobaoInvTurnoverQueryAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoInvTurnoverQueryAPIRequest) GetPageSize() string {
	return r._pageSize
}
