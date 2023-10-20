package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarFpcarRestpayReceive 门店线下已收尾款
// tmall.car.fpcar.restpay.receive
//
// 提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣)
func TmallCarFpcarRestpayReceive(clt *core.SDKClient, req *servicecenter.TmallCarFpcarRestpayReceiveAPIRequest, resp *servicecenter.TmallCarFpcarRestpayReceiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
