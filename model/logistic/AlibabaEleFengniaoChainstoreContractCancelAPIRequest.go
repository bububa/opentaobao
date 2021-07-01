package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoChainstoreContractCancelAPIRequest
门店解约接口 API请求
alibaba.ele.fengniao.chainstore.contract.cancel

调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单 */
type AlibabaEleFengniaoChainstoreContractCancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *AlibabaEleFengniaoChainstoreContractCancelData
}

// New
