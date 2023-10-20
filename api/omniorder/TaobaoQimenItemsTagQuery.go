package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoQimenItemsTagQuery 打标结果查询-商品维度
// taobao.qimen.items.tag.query
//
// 调用该接口，查询某个/某批商品上的标
func TaobaoQimenItemsTagQuery(clt *core.SDKClient, req *omniorder.TaobaoQimenItemsTagQueryAPIRequest, resp *omniorder.TaobaoQimenItemsTagQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
