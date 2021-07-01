package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsEfspCountsubsidyAPIRequest
计算补助金额 API请求
alibaba.alisports.efsp.countsubsidy

计算补助金额 */
type AlibabaAlisportsEfspCountsubsidyAPIRequest struct {
	model.Params
	// 订单金额
	_sumAmount int64
	// 健身房ID
	_gymId string
	// 企业ID
	_enterpriseId string
	// alipayId
	_alipayId string
	// 健身房所在省市
	_districtCode string
}

// New
