package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableMetaGetAPIRequest 智能应用服务登记工作表元数据查询 API请求
// taobao.smartapp.table.meta.get
//
// 智能应用服务登记工作表元数据查询
type TaobaoSmartappTableMetaGetAPIRequest struct {
	model.Params
	// 工作表表单id
	_tableId string
}

// NewTaobaoSmartappTableMetaGetRequest 初始化TaobaoSmartappTableMetaGetAPIRequest对象
func NewTaobaoSmartappTableMetaGetRequest() *TaobaoSmartappTableMetaGetAPIRequest {
	return &TaobaoSmartappTableMetaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappTableMetaGetAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.meta.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappTableMetaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappTableMetaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableId is TableId Setter
// 工作表表单id
func (r *TaobaoSmartappTableMetaGetAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaoSmartappTableMetaGetAPIRequest) GetTableId() string {
	return r._tableId
}
