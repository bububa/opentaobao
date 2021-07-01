package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseSynchronizeplansAPIRequest
同步租赁方案 API请求
tmall.car.lease.synchronizeplans

租赁公司同步还款计划 */
type TmallCarLeaseSynchronizeplansAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 租赁计划
	_plans []CarLeasePlanDo
}

// New
