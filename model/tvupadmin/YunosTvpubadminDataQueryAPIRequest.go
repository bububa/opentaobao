package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDataQueryAPIRequest
魔盒统计数据查询接口 API请求
yunos.tvpubadmin.data.query

用于华数查询魔盒上的一些用户统计数据 */
type YunosTvpubadminDataQueryAPIRequest struct {
	model.Params
	// 表名
	_tableName string
	// 列名
	_columns string
	// UUID
	_uuid string
	// 数据类型
	_dataTypeId int64
	// 日期
	_date string
	// 页码
	_pageNo int64
	// 每页个数
	_pageSize int64
}

// New
