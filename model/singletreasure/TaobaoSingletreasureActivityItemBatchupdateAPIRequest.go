package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityItemBatchupdateAPIRequest
批量修改商品接口 API请求
taobao.singletreasure.activity.item.batchupdate

批量修改商品优惠接口 */
type TaobaoSingletreasureActivityItemBatchupdateAPIRequest struct {
	model.Params
	// 系统入参
	_itemDetailInfo *ItemDetailInfoBatchCreateDto
}

// New
