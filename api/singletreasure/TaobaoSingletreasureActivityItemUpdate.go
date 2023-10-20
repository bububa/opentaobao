package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemUpdate 更新单品优惠接口
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
func TaobaoSingletreasureActivityItemUpdate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemUpdateAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityItemUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
