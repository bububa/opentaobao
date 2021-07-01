package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthHealthRecordHaveAPIRequest
判断用户的慢健康健康档案是否建设完成 API请求
alibaba.alihealth.health.record.have

判断用户的慢健康健康档案是否建设完成 */
type AlibabaAlihealthHealthRecordHaveAPIRequest struct {
	model.Params
	// 入参
	_request1 *HaveRecordRequest
}

// New
