package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

/* TaobaoSingletreasureActivityNameQuery
查询官方的活动名称接口
taobao.singletreasure.activity.name.query

查询官方的活动名称列表接口 */
func TaobaoSingletreasureActivityNameQuery(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityNameQueryAPIRequest, session string) (*singletreasure.TaobaoSingletreasureActivityNameQueryAPIResponse, error) {
	var resp singletreasure.TaobaoSingletreasureActivityNameQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
