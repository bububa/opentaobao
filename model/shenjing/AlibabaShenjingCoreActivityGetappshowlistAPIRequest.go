package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaShenjingCoreActivityGetappshowlistAPIRequest
获取神鲸活动列表 API请求
alibaba.shenjing.core.activity.getappshowlist

获取神鲸活动列表 */
type AlibabaShenjingCoreActivityGetappshowlistAPIRequest struct {
	model.Params
	// 验权对象
	_workBenchContext *WorkBenchContext
	// 时间戳
	_timestamp1 int64
	// 页码
	_page int64
	// 一页行数
	_size int64
}

// New
