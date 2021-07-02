package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAppraiseOrderQuery 闲鱼验货担保订单详情查询V1
// alibaba.idle.appraise.order.query
//
// 鉴定商调用该接口获取订单状态
func AlibabaIdleAppraiseOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleAppraiseOrderQueryAPIRequest, session string) (*idle.AlibabaIdleAppraiseOrderQueryAPIResponse, error) {
	var resp idle.AlibabaIdleAppraiseOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
