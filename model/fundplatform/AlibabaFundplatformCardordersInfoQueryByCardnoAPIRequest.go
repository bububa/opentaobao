package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest
通过卡号查询卡信息 API请求
alibaba.fundplatform.cardorders.info.query.by.cardno

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息 */
type AlibabaFundplatformCardordersInfoQueryByCardnoAPIRequest struct {
	model.Params
	// 请求结构体
	_parameters *CardMakingInfoQueryRequest
}

// New
