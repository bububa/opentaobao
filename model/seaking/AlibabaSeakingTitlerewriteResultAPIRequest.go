package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingTitlerewriteResultAPIRequest
获取标题改写任务结果 API请求
alibaba.seaking.titlerewrite.result

获取标题改写任务结果 */
type AlibabaSeakingTitlerewriteResultAPIRequest struct {
	model.Params
	// token来源站点
	_tokenFrom string
	// 任务id
	_taskId int64
	// 用户token
	_token string
}

// New
