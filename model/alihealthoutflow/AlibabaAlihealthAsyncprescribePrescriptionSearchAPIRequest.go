package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest
异步开方处方查询 API请求
alibaba.alihealth.asyncprescribe.prescription.search

异步开方处方查询 */
type AlibabaAlihealthAsyncprescribePrescriptionSearchAPIRequest struct {
	model.Params
	// 查询入参
	_searchRequest *AsyncPrescribeSearchRequest
}

// New
