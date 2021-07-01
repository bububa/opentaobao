package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsBotsSkilsListAPIRequest
对外设备获取技能列表 API请求
alibaba.ailabs.bots.skils.list

获取ai开放平台技能列表 */
type AlibabaAilabsBotsSkilsListAPIRequest struct {
	model.Params
	// 当前页
	_pageIndex int64
	// 分页单位
	_pageSize int64
}

// New
