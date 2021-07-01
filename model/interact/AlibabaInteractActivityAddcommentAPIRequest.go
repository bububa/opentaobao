package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityAddcommentAPIRequest
微淘评论接口 API请求
alibaba.interact.activity.addcomment

发表评论，并返回楼层 */
type AlibabaInteractActivityAddcommentAPIRequest struct {
	model.Params
	// 该字段为评论内容
	_content string
	// 评论feedid
	_feedId int64
	// 发评论的业务id
	_bizId string
}

// New
