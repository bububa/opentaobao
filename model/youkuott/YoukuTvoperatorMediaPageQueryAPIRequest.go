package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuTvoperatorMediaPageQueryAPIRequest
运营商全量媒资分页查询 API请求
youku.tvoperator.media.page.query

分页获取渠道全量媒资 */
type YoukuTvoperatorMediaPageQueryAPIRequest struct {
	model.Params
	// 系统信息（和服务提供方确认)
	_systemInfo string
	// 从第一页开始
	_pageNo int64
	// 页面大小
	_pageSize int64
	// 节目programId
	_programId int64
}

// New
