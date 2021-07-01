package aetask

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressInteractiveTaskDeliveryQueryAPIRequest
AE互动任务投放 API请求
aliexpress.interactive.task.delivery.query

将内部配置好的任务，如浏览商品，店铺投放给外部ISV */
type AliexpressInteractiveTaskDeliveryQueryAPIRequest struct {
	model.Params
	// 返回结果
	_requestDto *QueryDeliveryRequestDto
}

// New
