package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoqimenitemstagquery 打标结果查询-商品维度
// taobao.qimen.items.tag.query
//
// 调用该接口，查询某个/某批商品上的标
func Taobaoqimenitemstagquery(clt *core.SDKClient, req *omniorder.TaobaoqimenitemstagqueryAPIRequest, session string) (*omniorder.TaobaoqimenitemstagqueryAPIResponse, error) {
	var resp omniorder.TaobaoqimenitemstagqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
