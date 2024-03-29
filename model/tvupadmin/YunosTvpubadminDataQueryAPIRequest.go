package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDataQueryAPIRequest 魔盒统计数据查询接口 API请求
// yunos.tvpubadmin.data.query
//
// 用于华数查询魔盒上的一些用户统计数据
type YunosTvpubadminDataQueryAPIRequest struct {
	model.Params
	// 表名
	_tableName string
	// 列名
	_columns string
	// UUID
	_uuid string
	// 日期
	_date string
	// 数据类型
	_dataTypeId int64
	// 页码
	_pageNo int64
	// 每页个数
	_pageSize int64
}

// NewYunosTvpubadminDataQueryRequest 初始化YunosTvpubadminDataQueryAPIRequest对象
func NewYunosTvpubadminDataQueryRequest() *YunosTvpubadminDataQueryAPIRequest {
	return &YunosTvpubadminDataQueryAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDataQueryAPIRequest) Reset() {
	r._tableName = ""
	r._columns = ""
	r._uuid = ""
	r._date = ""
	r._dataTypeId = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDataQueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.data.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDataQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDataQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableName is TableName Setter
// 表名
func (r *YunosTvpubadminDataQueryAPIRequest) SetTableName(_tableName string) error {
	r._tableName = _tableName
	r.Set("table_name", _tableName)
	return nil
}

// GetTableName TableName Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetTableName() string {
	return r._tableName
}

// SetColumns is Columns Setter
// 列名
func (r *YunosTvpubadminDataQueryAPIRequest) SetColumns(_columns string) error {
	r._columns = _columns
	r.Set("columns", _columns)
	return nil
}

// GetColumns Columns Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetColumns() string {
	return r._columns
}

// SetUuid is Uuid Setter
// UUID
func (r *YunosTvpubadminDataQueryAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetUuid() string {
	return r._uuid
}

// SetDate is Date Setter
// 日期
func (r *YunosTvpubadminDataQueryAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetDate() string {
	return r._date
}

// SetDataTypeId is DataTypeId Setter
// 数据类型
func (r *YunosTvpubadminDataQueryAPIRequest) SetDataTypeId(_dataTypeId int64) error {
	r._dataTypeId = _dataTypeId
	r.Set("data_type_id", _dataTypeId)
	return nil
}

// GetDataTypeId DataTypeId Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetDataTypeId() int64 {
	return r._dataTypeId
}

// SetPageNo is PageNo Setter
// 页码
func (r *YunosTvpubadminDataQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页个数
func (r *YunosTvpubadminDataQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunosTvpubadminDataQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolYunosTvpubadminDataQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDataQueryRequest()
	},
}

// GetYunosTvpubadminDataQueryRequest 从 sync.Pool 获取 YunosTvpubadminDataQueryAPIRequest
func GetYunosTvpubadminDataQueryAPIRequest() *YunosTvpubadminDataQueryAPIRequest {
	return poolYunosTvpubadminDataQueryAPIRequest.Get().(*YunosTvpubadminDataQueryAPIRequest)
}

// ReleaseYunosTvpubadminDataQueryAPIRequest 将 YunosTvpubadminDataQueryAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDataQueryAPIRequest(v *YunosTvpubadminDataQueryAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDataQueryAPIRequest.Put(v)
}
