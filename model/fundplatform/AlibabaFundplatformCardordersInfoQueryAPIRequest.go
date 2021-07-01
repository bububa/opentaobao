package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardordersInfoQueryAPIRequest
根据制卡单分页查询卡信息 API请求
alibaba.fundplatform.cardorders.info.query

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息 */
type AlibabaFundplatformCardordersInfoQueryAPIRequest struct {
	model.Params
	// 请求结构体
	_parameters *CardMakingInfoQueryRequest
}

// New
