package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmActivityDataUpdateAPIRequest
私域导购数据回流接口 API请求
alibaba.lsy.crm.activity.data.update

私域导购数据回流接口 */
type AlibabaLsyCrmActivityDataUpdateAPIRequest struct {
	model.Params
	// 入参对象
	_reqDTO *NrtCrmActivityStatisticsDataReq
}

// New
