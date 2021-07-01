package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpPresalespackageConsignAPIRequest
预售预包尾款推单发货 API请求
alibaba.ascp.presalespackage.consign

预售预包尾款发货后推单处理 */
type AlibabaAscpPresalespackageConsignAPIRequest struct {
	model.Params
	// 入参
	_requestParams *Requestparams
}

// New
