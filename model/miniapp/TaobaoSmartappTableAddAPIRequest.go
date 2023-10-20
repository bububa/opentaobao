package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableAddAPIRequest 智能应用服务登记工作表数据新增 API请求
// taobao.smartapp.table.add
//
// 智能应用服务登记工作表数据新增
type TaobaoSmartappTableAddAPIRequest struct {
	model.Params
	// 新增记录内容格式为JSON，新增仅支持单条新增，各组件新增输入格式详见文档
	_record string
	// 待插入表id
	_tableId string
}

// NewTaobaoSmartappTableAddRequest 初始化TaobaoSmartappTableAddAPIRequest对象
func NewTaobaoSmartappTableAddRequest() *TaobaoSmartappTableAddAPIRequest {
	return &TaobaoSmartappTableAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappTableAddAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappTableAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappTableAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecord is Record Setter
// 新增记录内容格式为JSON，新增仅支持单条新增，各组件新增输入格式详见文档
func (r *TaobaoSmartappTableAddAPIRequest) SetRecord(_record string) error {
	r._record = _record
	r.Set("record", _record)
	return nil
}

// GetRecord Record Getter
func (r TaobaoSmartappTableAddAPIRequest) GetRecord() string {
	return r._record
}

// SetTableId is TableId Setter
// 待插入表id
func (r *TaobaoSmartappTableAddAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaoSmartappTableAddAPIRequest) GetTableId() string {
	return r._tableId
}
