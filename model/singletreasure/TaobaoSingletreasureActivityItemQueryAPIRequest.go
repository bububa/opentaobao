package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityItemQueryAPIRequest
查询活动下的优惠信息 API请求
taobao.singletreasure.activity.item.query

分页查询活动下的商品优惠信息 */
type TaobaoSingletreasureActivityItemQueryAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
	// 页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// New
