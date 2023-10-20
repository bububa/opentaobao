package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityCreate 活动创建接口
// taobao.singletreasure.activity.create
//
// 创建优惠活动
func TaobaoSingletreasureActivityCreate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityCreateAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
