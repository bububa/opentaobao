package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingTitlerewriteSubmitAPIRequest
提交标题改写任务 API请求
alibaba.seaking.titlerewrite.submit

提交标题改写任务 */
type AlibabaSeakingTitlerewriteSubmitAPIRequest struct {
	model.Params
	// 任务详情列表
	_titleRewriteDetailList []TitleRewriteDetailDto
	// token来源站点
	_tokenFrom string
	// 用户token
	_token string
}

// New
