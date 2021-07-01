package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmActivityUpdateAPIRequest
ISV活动修改 API请求
alibaba.lsy.crm.activity.update

ISV活动修改 */
type AlibabaLsyCrmActivityUpdateAPIRequest struct {
	model.Params
	// 入参
	_nrtUpdateActivityReq *NrtUpdateActivityReq
}

// New
