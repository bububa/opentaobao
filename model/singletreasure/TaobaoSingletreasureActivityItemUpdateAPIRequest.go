package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityItemUpdateAPIRequest
更新单品优惠接口 API请求
taobao.singletreasure.activity.item.update

更新单品优惠接口 */
type TaobaoSingletreasureActivityItemUpdateAPIRequest struct {
	model.Params
	// 修改接口的入参对象
	_itemDetailInfo *ItemDetailInfoCreateDto
}

// New
