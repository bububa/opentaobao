package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleappraiseorderperform 闲鱼验货宝服务商订单履约
// alibaba.idle.appraise.order.perform
//
// 闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
func Alibabaidleappraiseorderperform(clt *core.SDKClient, req *idle.AlibabaidleappraiseorderperformAPIRequest, session string) (*idle.AlibabaidleappraiseorderperformAPIResponse, error) {
	var resp idle.AlibabaidleappraiseorderperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
