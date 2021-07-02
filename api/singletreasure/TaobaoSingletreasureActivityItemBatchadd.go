package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityItemBatchadd 批量添加商品接口
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
func TaobaoSingletreasureActivityItemBatchadd(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityItemBatchaddAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityItemBatchaddAPIResponse, error) {
	var resp singletreasure.TaobaoSingletreasureActivityItemBatchaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
