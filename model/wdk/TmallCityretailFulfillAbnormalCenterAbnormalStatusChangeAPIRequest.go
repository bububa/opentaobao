package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest
同城零售履约异常中心异常单处理结果回调接口 API请求
tmall.cityretail.fulfill.abnormal.center.abnormal.status.change

同城零售履约异常中心异常单处理结果回调接口 */
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeAPIRequest struct {
	model.Params
	// 入参
	_abnormalStatusChangeDto []AbnormalStatusChangeDto
}

// New
