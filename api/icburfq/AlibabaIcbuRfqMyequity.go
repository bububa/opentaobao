package icburfq

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icburfq"
)

// AlibabaIcbuRfqMyequity 我的权益
// alibaba.icbu.rfq.myequity
//
// 查询供应商权益接口
func AlibabaIcbuRfqMyequity(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqMyequityAPIRequest, resp *icburfq.AlibabaIcbuRfqMyequityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
