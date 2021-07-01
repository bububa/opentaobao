package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmActivityGetbaseinfoAPIRequest
ISV查询活动 API请求
alibaba.lsy.crm.activity.getbaseinfo

ISV查询活动 */
type AlibabaLsyCrmActivityGetbaseinfoAPIRequest struct {
	model.Params
	// 入参
	_nrtQueryActivityReq *NrtQueryActivityReq
}

// New
