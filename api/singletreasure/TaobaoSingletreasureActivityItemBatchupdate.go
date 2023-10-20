package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemBatchupdate 批量修改商品接口
// taobao.singletreasure.activity.item.batchupdate
//
// 批量修改商品优惠接口
func TaobaoSingletreasureActivityItemBatchupdate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemBatchupdateAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityItemBatchupdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
