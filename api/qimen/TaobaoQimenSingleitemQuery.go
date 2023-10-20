package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenSingleitemQuery 商品查询接口
// taobao.qimen.singleitem.query
//
// 商品查询接口
func TaobaoQimenSingleitemQuery(clt *core.SDKClient, req *qimen.TaobaoQimenSingleitemQueryAPIRequest, resp *qimen.TaobaoQimenSingleitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
