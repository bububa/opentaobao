package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Taobaoalitripieagentrefundnewreceive 商家退票受理申请(对外)
// taobao.alitrip.ie.agent.refund.new.receive
//
// 允许代理商通过top接口受理退票申请
func Taobaoalitripieagentrefundnewreceive(clt *core.SDKClient, req *ieagency.TaobaoalitripieagentrefundnewreceiveAPIRequest, session string) (*ieagency.TaobaoalitripieagentrefundnewreceiveAPIResponse, error) {
	var resp ieagency.TaobaoalitripieagentrefundnewreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
