package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// Taobaotraderatelistadd 针对父子订单新增批量评价
// taobao.traderate.list.add
//
// 针对父子订单新增批量评价(&lt;font color=&#34;red&#34;&gt;注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价&lt;/font&gt;)
func Taobaotraderatelistadd(clt *core.SDKClient, req *traderate.TaobaotraderatelistaddAPIRequest, session string) (*traderate.TaobaotraderatelistaddAPIResponse, error) {
	var resp traderate.TaobaotraderatelistaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
