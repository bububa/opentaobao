package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenCombineitemDelete 组合货品删除接口
// taobao.qimen.combineitem.delete
//
// 组合货品删除
func TaobaoQimenCombineitemDelete(clt *core.SDKClient, req *qimen.TaobaoQimenCombineitemDeleteAPIRequest, resp *qimen.TaobaoQimenCombineitemDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
