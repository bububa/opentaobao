package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthRxPrescriptionDetailAPIRequest
处方详情 API请求
alibaba.alihealth.rx.prescription.detail

获取处方结构化信息 */
type AlibabaAlihealthRxPrescriptionDetailAPIRequest struct {
	model.Params
	// 查询参数
	_query *RxPrescriptionQuery
}

// New
