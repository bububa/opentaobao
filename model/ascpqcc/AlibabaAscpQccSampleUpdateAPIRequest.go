package ascpqcc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpQccSampleUpdateAPIRequest
品控中心更新样品信息 API请求
alibaba.ascp.qcc.sample.update

品控中心更新样品信息 */
type AlibabaAscpQccSampleUpdateAPIRequest struct {
	model.Params
	// 更新请求参数
	_updateRequest *UpdateSampleRequest
}

// New
