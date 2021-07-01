package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmActivityListGetAPIRequest
获取活动列表接口 API请求
alibaba.westcrm.activity.list.get

获取活动列表提供给友盟&互动屏 */
type AlibabaWestcrmActivityListGetAPIRequest struct {
	model.Params
	// 活动状态
	_status int64
	// 园区id
	_campusId int64
	// 排序方向
	_sord string
	// 页,默认第一页
	_page int64
	// 排序字段
	_sidx string
	// 分页偏移量 eq . limit offset ,rows
	_offset int64
	// 页大小,默认每页查询10条数据
	_rows int64
}

// New
