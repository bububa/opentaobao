package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindataqueryAPIRequest 魔盒统计数据查询接口 API请求
// yunos.tvpubadmin.data.query
//
// 用于华数查询魔盒上的一些用户统计数据
type YunostvpubadmindataqueryAPIRequest struct {
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

// NewYunostvpubadmindataqueryRequest 初始化YunostvpubadmindataqueryAPIRequest对象
func NewYunostvpubadmindataqueryRequest() *YunostvpubadmindataqueryAPIRequest {
	return &YunostvpubadmindataqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindataqueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.data.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindataqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindataqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTableName is TableName Setter
// 表名
func (r *YunostvpubadmindataqueryAPIRequest) SetTableName(_tableName string) error {
	r._tableName = _tableName
	r.Set("table_name", _tableName)
	return nil
}

// GetTableName TableName Getter
func (r YunostvpubadmindataqueryAPIRequest) GetTableName() string {
	return r._tableName
}

// SetColumns is Columns Setter
// 列名
func (r *YunostvpubadmindataqueryAPIRequest) SetColumns(_columns string) error {
	r._columns = _columns
	r.Set("columns", _columns)
	return nil
}

// GetColumns Columns Getter
func (r YunostvpubadmindataqueryAPIRequest) GetColumns() string {
	return r._columns
}

// SetUuid is Uuid Setter
// UUID
func (r *YunostvpubadmindataqueryAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r YunostvpubadmindataqueryAPIRequest) GetUuid() string {
	return r._uuid
}

// SetDate is Date Setter
// 日期
func (r *YunostvpubadmindataqueryAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r YunostvpubadmindataqueryAPIRequest) GetDate() string {
	return r._date
}

// SetDataTypeId is DataTypeId Setter
// 数据类型
func (r *YunostvpubadmindataqueryAPIRequest) SetDataTypeId(_dataTypeId int64) error {
	r._dataTypeId = _dataTypeId
	r.Set("data_type_id", _dataTypeId)
	return nil
}

// GetDataTypeId DataTypeId Getter
func (r YunostvpubadmindataqueryAPIRequest) GetDataTypeId() int64 {
	return r._dataTypeId
}

// SetPageNo is PageNo Setter
// 页码
func (r *YunostvpubadmindataqueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r YunostvpubadmindataqueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页个数
func (r *YunostvpubadmindataqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r YunostvpubadmindataqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
