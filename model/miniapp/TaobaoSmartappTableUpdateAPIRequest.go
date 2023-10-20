package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartapptableupdateAPIRequest 智能应用服务登记工作表数据更新 API请求
// taobao.smartapp.table.update
//
// 智能应用服务登记工作表数据更新
type TaobaosmartapptableupdateAPIRequest struct {
	model.Params
	// 待更新工作表id
	_tableId string
	// 待更新记录主键
	_id string
	// 提交更新数据，更新仅支持单条新增，格式同新增记录接口
	_record string
}

// NewTaobaosmartapptableupdateRequest 初始化TaobaosmartapptableupdateAPIRequest对象
func NewTaobaosmartapptableupdateRequest() *TaobaosmartapptableupdateAPIRequest {
	return &TaobaosmartapptableupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosmartapptableupdateAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosmartapptableupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosmartapptableupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableId is TableId Setter
// 待更新工作表id
func (r *TaobaosmartapptableupdateAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaosmartapptableupdateAPIRequest) GetTableId() string {
	return r._tableId
}

// SetId is Id Setter
// 待更新记录主键
func (r *TaobaosmartapptableupdateAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaosmartapptableupdateAPIRequest) GetId() string {
	return r._id
}

// SetRecord is Record Setter
// 提交更新数据，更新仅支持单条新增，格式同新增记录接口
func (r *TaobaosmartapptableupdateAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r TaobaosmartapptableupdateAPIRequest) GetRecord() string {
	return r._record
}
