package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayTradeQuery 查询支付结果接口
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
func AliyunAliosPayTradeQuery(clt *core.SDKClient, req *aliospay.AliyunAliosPayTradeQueryAPIRequest, resp *aliospay.AliyunAliosPayTradeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
