package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
魔盒统计数据查询接口 APIRequest
yunos.tvpubadmin.data.query

用于华数查询魔盒上的一些用户统计数据
*/
type YunosTvpubadminDataQueryRequest struct {
    model.Params

    // 表名
    tableName   string 

    // 列名
    columns   string 

    // UUID
    uuid   string 

    // 数据类型
    dataTypeId   int64 

    // 日期
    date   string 

    // 页码
    pageNo   int64 

    // 每页个数
    pageSize   int64 

}

func NewYunosTvpubadminDataQueryRequest() *YunosTvpubadminDataQueryRequest{
    return &YunosTvpubadminDataQueryRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDataQueryRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.data.query"
}

func (r YunosTvpubadminDataQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDataQueryRequest) SetTableName(tableName string) error {
    r.tableName = tableName
    r.Set("table_name", tableName)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetTableName() string {
    return r.tableName
}

func (r *YunosTvpubadminDataQueryRequest) SetColumns(columns string) error {
    r.columns = columns
    r.Set("columns", columns)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetColumns() string {
    return r.columns
}

func (r *YunosTvpubadminDataQueryRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetUuid() string {
    return r.uuid
}

func (r *YunosTvpubadminDataQueryRequest) SetDataTypeId(dataTypeId int64) error {
    r.dataTypeId = dataTypeId
    r.Set("data_type_id", dataTypeId)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetDataTypeId() int64 {
    return r.dataTypeId
}

func (r *YunosTvpubadminDataQueryRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetDate() string {
    return r.date
}

func (r *YunosTvpubadminDataQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *YunosTvpubadminDataQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r YunosTvpubadminDataQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

