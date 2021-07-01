package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAccountStatusUpdateAPIRequest
修改账户级别关键词推广状态 API请求
alibaba.scbp.account.status.update

修改账户级别关键词推广状态 */
type AlibabaScbpAccountStatusUpdateAPIRequest struct {
	model.Params
	// on:开启,off:暂停
	_status string
}

// New
