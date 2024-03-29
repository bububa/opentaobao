package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartappTableGetAPIRequest 智能应用服务登记工作表数据查询 API请求
// taobao.smartapp.table.get
//
// 智能应用服务登记工作表数据查询
type TaobaoSmartappTableGetAPIRequest struct {
	model.Params
	// 查询创建日期的起始时间，若“创建 起始时间与结束时间”与“修改 起始时间与结束时间”两组数据均传入，仅筛选创建时间
	_startCreatedDate string
	// 查询创建日期的终止时间
	_endCreatedDate string
	// 查询修改日期的起始时间，仅当“创建 起始时间与结束时间”均未传入时，筛选修改时间，并按修改时间倒叙排序
	_startModifiedDate string
	// 查询修改日期的终止时间
	_endModifiedDate string
	// 数据空间表名
	_tableId string
	// 数据主键ID
	_recordId string
	// 页面索引，若不传入默认为1
	_pageIndex int64
	// 页面容量，若不传入默认为200
	_pageSize int64
}

// NewTaobaoSmartappTableGetRequest 初始化TaobaoSmartappTableGetAPIRequest对象
func NewTaobaoSmartappTableGetRequest() *TaobaoSmartappTableGetAPIRequest {
	return &TaobaoSmartappTableGetAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSmartappTableGetAPIRequest) Reset() {
	r._startCreatedDate = ""
	r._endCreatedDate = ""
	r._startModifiedDate = ""
	r._endModifiedDate = ""
	r._tableId = ""
	r._recordId = ""
	r._pageIndex = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartappTableGetAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartappTableGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSmartappTableGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartCreatedDate is StartCreatedDate Setter
// 查询创建日期的起始时间，若“创建 起始时间与结束时间”与“修改 起始时间与结束时间”两组数据均传入，仅筛选创建时间
func (r *TaobaoSmartappTableGetAPIRequest) SetStartCreatedDate(_startCreatedDate string) error {
	r._startCreatedDate = _startCreatedDate
	r.Set("start_created_date", _startCreatedDate)
	return nil
}

// GetStartCreatedDate StartCreatedDate Getter
func (r TaobaoSmartappTableGetAPIRequest) GetStartCreatedDate() string {
	return r._startCreatedDate
}

// SetEndCreatedDate is EndCreatedDate Setter
// 查询创建日期的终止时间
func (r *TaobaoSmartappTableGetAPIRequest) SetEndCreatedDate(_endCreatedDate string) error {
	r._endCreatedDate = _endCreatedDate
	r.Set("end_created_date", _endCreatedDate)
	return nil
}

// GetEndCreatedDate EndCreatedDate Getter
func (r TaobaoSmartappTableGetAPIRequest) GetEndCreatedDate() string {
	return r._endCreatedDate
}

// SetStartModifiedDate is StartModifiedDate Setter
// 查询修改日期的起始时间，仅当“创建 起始时间与结束时间”均未传入时，筛选修改时间，并按修改时间倒叙排序
func (r *TaobaoSmartappTableGetAPIRequest) SetStartModifiedDate(_startModifiedDate string) error {
	r._startModifiedDate = _startModifiedDate
	r.Set("start_modified_date", _startModifiedDate)
	return nil
}

// GetStartModifiedDate StartModifiedDate Getter
func (r TaobaoSmartappTableGetAPIRequest) GetStartModifiedDate() string {
	return r._startModifiedDate
}

// SetEndModifiedDate is EndModifiedDate Setter
// 查询修改日期的终止时间
func (r *TaobaoSmartappTableGetAPIRequest) SetEndModifiedDate(_endModifiedDate string) error {
	r._endModifiedDate = _endModifiedDate
	r.Set("end_modified_date", _endModifiedDate)
	return nil
}

// GetEndModifiedDate EndModifiedDate Getter
func (r TaobaoSmartappTableGetAPIRequest) GetEndModifiedDate() string {
	return r._endModifiedDate
}

// SetTableId is TableId Setter
// 数据空间表名
func (r *TaobaoSmartappTableGetAPIRequest) SetTableId(_tableId string) error {
	r._tableId = _tableId
	r.Set("table_id", _tableId)
	return nil
}

// GetTableId TableId Getter
func (r TaobaoSmartappTableGetAPIRequest) GetTableId() string {
	return r._tableId
}

// SetRecordId is RecordId Setter
// 数据主键ID
func (r *TaobaoSmartappTableGetAPIRequest) SetRecordId(_recordId string) error {
	r._recordId = _recordId
	r.Set("record_id", _recordId)
	return nil
}

// GetRecordId RecordId Getter
func (r TaobaoSmartappTableGetAPIRequest) GetRecordId() string {
	return r._recordId
}

// SetPageIndex is PageIndex Setter
// 页面索引，若不传入默认为1
func (r *TaobaoSmartappTableGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoSmartappTableGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 页面容量，若不传入默认为200
func (r *TaobaoSmartappTableGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSmartappTableGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoSmartappTableGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSmartappTableGetRequest()
	},
}

// GetTaobaoSmartappTableGetRequest 从 sync.Pool 获取 TaobaoSmartappTableGetAPIRequest
func GetTaobaoSmartappTableGetAPIRequest() *TaobaoSmartappTableGetAPIRequest {
	return poolTaobaoSmartappTableGetAPIRequest.Get().(*TaobaoSmartappTableGetAPIRequest)
}

// ReleaseTaobaoSmartappTableGetAPIRequest 将 TaobaoSmartappTableGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSmartappTableGetAPIRequest(v *TaobaoSmartappTableGetAPIRequest) {
	v.Reset()
	poolTaobaoSmartappTableGetAPIRequest.Put(v)
}
