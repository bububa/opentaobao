package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableUpdateAPIRequest 智能应用服务登记工作表数据更新 API请求
// taobao.smartapp.table.update
//
// 智能应用服务登记工作表数据更新
type TaobaoSmartappTableUpdateAPIRequest struct {
	model.Params
	// 待更新工作表id
	_tableId string
	// 待更新记录主键
	_id string
	// 提交更新数据，更新仅支持单条新增，格式同新增记录接口
	_record string
}

// NewTaobaoSmartappTableUpdateRequest 初始化TaobaoSmartappTableUpdateAPIRequest对象
func NewTaobaoSmartappTableUpdateRequest() *TaobaoSmartappTableUpdateAPIRequest {
	return &TaobaoSmartappTableUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappTableUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappTableUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappTableUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableId is TableId Setter
// 待更新工作表id
func (r *TaobaoSmartappTableUpdateAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaoSmartappTableUpdateAPIRequest) GetTableId() string {
	return r._tableId
}

// SetId is Id Setter
// 待更新记录主键
func (r *TaobaoSmartappTableUpdateAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoSmartappTableUpdateAPIRequest) GetId() string {
	return r._id
}

// SetRecord is Record Setter
// 提交更新数据，更新仅支持单条新增，格式同新增记录接口
func (r *TaobaoSmartappTableUpdateAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r TaobaoSmartappTableUpdateAPIRequest) GetRecord() string {
	return r._record
}
