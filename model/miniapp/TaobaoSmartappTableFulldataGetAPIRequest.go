package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartapptablefulldatagetAPIRequest 智能应用工作表地址查询 API请求
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
type TaobaosmartapptablefulldatagetAPIRequest struct {
	model.Params
	// 数据空间表名
	_tableId string
	// 数据主键ID
	_id string
}

// NewTaobaosmartapptablefulldatagetRequest 初始化TaobaosmartapptablefulldatagetAPIRequest对象
func NewTaobaosmartapptablefulldatagetRequest() *TaobaosmartapptablefulldatagetAPIRequest {
	return &TaobaosmartapptablefulldatagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosmartapptablefulldatagetAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.fulldata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosmartapptablefulldatagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosmartapptablefulldatagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableId is TableId Setter
// 数据空间表名
func (r *TaobaosmartapptablefulldatagetAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaosmartapptablefulldatagetAPIRequest) GetTableId() string {
	return r._tableId
}

// SetId is Id Setter
// 数据主键ID
func (r *TaobaosmartapptablefulldatagetAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaosmartapptablefulldatagetAPIRequest) GetId() string {
	return r._id
}
