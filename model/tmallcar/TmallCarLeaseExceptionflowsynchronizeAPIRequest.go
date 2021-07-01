package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseExceptionflowsynchronizeAPIRequest
天猫开新车租后异常流线下处理状态通知接口 API请求
tmall.car.lease.exceptionflowsynchronize

天猫开新车租后异常流线下处理状态通知接口 */
type TmallCarLeaseExceptionflowsynchronizeAPIRequest struct {
	model.Params
	// 天猫开新车订单id
	_orderId int64
	// 1:开始切换为异常流，2:线下处理完成
	_status int64
	// 异常流类型,0.退车,1.买断,2.分期，3.续租
	_flowType int64
	// 切换原因描述
	_desc string
}

// New
