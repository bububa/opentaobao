package singletreasure

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/singletreasure"
)

// TaobaoSingletreasureActivityNameQuery 查询官方的活动名称接口
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
func TaobaoSingletreasureActivityNameQuery(clt *core.SDKClient, req *singletreasure.TaobaoSingletreasureActivityNameQueryAPIRequest, resp *singletreasure.TaobaoSingletreasureActivityNameQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
