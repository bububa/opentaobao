package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsFulfillQaOrderCreateAPIRequest
创单接口 API请求
alibaba.tcls.fulfill.qa.order.create

根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据） */
type AlibabaTclsFulfillQaOrderCreateAPIRequest struct {
	model.Params
	// 原始履约单号
	_fulfillOrderId string
	// 目标ip
	_targetIp string
	// 执行人姓名
	_creator string
	// 执行人工号
	_jobNo string
}

// New
