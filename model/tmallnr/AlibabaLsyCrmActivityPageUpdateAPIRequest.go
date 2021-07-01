package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmActivityPageUpdateAPIRequest
ISV活动页面创建修改 API请求
alibaba.lsy.crm.activity.page.update

ISV活动页面创建修改 */
type AlibabaLsyCrmActivityPageUpdateAPIRequest struct {
	model.Params
	// 入参
	_nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq
}

// New
