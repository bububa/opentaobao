package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleParttimeSynclogAPIRequest
兼职同步日志 API请求
alibaba.idle.parttime.synclog

提供给供应商查询的接口 */
type AlibabaIdleParttimeSynclogAPIRequest struct {
	model.Params
	// 查询岗位同步开始时间
	_startTime int64
	// 查询岗位同步结束时间
	_endTime int64
	// 查询的类型, 0:岗位
	_type int64
	// 页大小
	_pageSize int64
	// 第几页, 从0开始
	_pageNum int64
	// 同步的id
	_syncIds []int64
}

// New
