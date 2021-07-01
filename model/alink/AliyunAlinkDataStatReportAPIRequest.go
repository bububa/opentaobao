package alink

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAlinkDataStatReportAPIRequest
外部离线统计数据上报 API请求
aliyun.alink.data.stat.report

外部合作厂商上报设备的明细数据，或者离线统计数据。 */
type AliyunAlinkDataStatReportAPIRequest struct {
	model.Params
	// 入参对象
	_paramBean *OuterDataBean
}

// New
