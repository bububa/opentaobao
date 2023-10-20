package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableFulldataGetAPIRequest 智能应用工作表地址查询 API请求
// taobao.smartapp.table.fulldata.get
//
// 智能应用工作表地址查询
type TaobaoSmartappTableFulldataGetAPIRequest struct {
	model.Params
	// 数据空间表名
	_tableId string
	// 数据主键ID
	_id string
}

// NewTaobaoSmartappTableFulldataGetRequest 初始化TaobaoSmartappTableFulldataGetAPIRequest对象
func NewTaobaoSmartappTableFulldataGetRequest() *TaobaoSmartappTableFulldataGetAPIRequest {
	return &TaobaoSmartappTableFulldataGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSmartappTableFulldataGetAPIRequest) Reset() {
	r._tableId = ""
	r._id = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappTableFulldataGetAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.fulldata.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappTableFulldataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappTableFulldataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableId is TableId Setter
// 数据空间表名
func (r *TaobaoSmartappTableFulldataGetAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaoSmartappTableFulldataGetAPIRequest) GetTableId() string {
	return r._tableId
}

// SetId is Id Setter
// 数据主键ID
func (r *TaobaoSmartappTableFulldataGetAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoSmartappTableFulldataGetAPIRequest) GetId() string {
	return r._id
}

var poolTaobaoSmartappTableFulldataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSmartappTableFulldataGetRequest()
	},
}

// GetTaobaoSmartappTableFulldataGetRequest 从 sync.Pool 获取 TaobaoSmartappTableFulldataGetAPIRequest
func GetTaobaoSmartappTableFulldataGetAPIRequest() *TaobaoSmartappTableFulldataGetAPIRequest {
	return poolTaobaoSmartappTableFulldataGetAPIRequest.Get().(*TaobaoSmartappTableFulldataGetAPIRequest)
}

// ReleaseTaobaoSmartappTableFulldataGetAPIRequest 将 TaobaoSmartappTableFulldataGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSmartappTableFulldataGetAPIRequest(v *TaobaoSmartappTableFulldataGetAPIRequest) {
	v.Reset()
	poolTaobaoSmartappTableFulldataGetAPIRequest.Put(v)
}
