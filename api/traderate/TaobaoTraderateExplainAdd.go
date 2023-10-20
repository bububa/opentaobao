package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// Taobaotraderateexplainadd 商城评价解释接口
// taobao.traderate.explain.add
//
// 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
func Taobaotraderateexplainadd(clt *core.SDKClient, req *traderate.TaobaotraderateexplainaddAPIRequest, session string) (*traderate.TaobaotraderateexplainaddAPIResponse, error) {
	var resp traderate.TaobaotraderateexplainaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
