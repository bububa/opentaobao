package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest
异步开方处方详情 API请求
alibaba.alihealth.asyncprescribe.prescription.detail

异步开方处方查询 */
type AlibabaAlihealthAsyncprescribePrescriptionDetailAPIRequest struct {
	model.Params
	// 入参
	_detailRequest *AsyncPrescribeDetailRequest
}

// New
