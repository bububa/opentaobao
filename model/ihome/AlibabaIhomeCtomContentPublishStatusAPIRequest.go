package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomContentPublishStatusAPIRequest
实拍图发布审核状态查询API API请求
alibaba.ihome.ctom.content.publish.status

实拍图发布审核状态查询API */
type AlibabaIhomeCtomContentPublishStatusAPIRequest struct {
	model.Params
	// 要查询投稿状态的ID列表
	_idList []int64
}

// New
