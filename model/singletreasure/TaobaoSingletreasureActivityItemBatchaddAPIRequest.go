package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityItemBatchaddAPIRequest
批量添加商品接口 API请求
taobao.singletreasure.activity.item.batchadd

向活动中批量添加商品优惠 */
type TaobaoSingletreasureActivityItemBatchaddAPIRequest struct {
	model.Params
	// 系统入参
	_itemDetailInfo *ItemDetailInfoBatchCreateDto
}

// New
