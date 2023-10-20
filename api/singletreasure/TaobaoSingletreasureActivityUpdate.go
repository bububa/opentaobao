package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityUpdate 修改活动接口
// taobao.singletreasure.activity.update
//
// 修改活动接口
func TaobaoSingletreasureActivityUpdate(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityUpdateAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
