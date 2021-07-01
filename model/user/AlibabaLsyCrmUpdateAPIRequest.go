package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmUpdateAPIRequest
跟进客资状态接口 API请求
alibaba.lsy.crm.update

同步客资状态接口 */
type AlibabaLsyCrmUpdateAPIRequest struct {
	model.Params
	// 更新客资状态对象
	_nrtUpdateRecordStatusDto *NrtUpdateRecordStatusDto
}

// New
