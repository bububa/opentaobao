package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

/* TaobaoSingletreasureActivityDelete
删除活动接口
taobao.singletreasure.activity.delete

删除优惠活动 */
func TaobaoSingletreasureActivityDelete(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityDeleteAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityDeleteAPIResponse, error) {
	var resp singletreasure.TaobaoSingletreasureActivityDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
