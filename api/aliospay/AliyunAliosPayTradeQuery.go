package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayTradeQuery 查询支付结果接口
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
func AliyunAliosPayTradeQuery(clt *core.SDKClient, req *aliospay.AliyunAliosPayTradeQueryAPIRequest, session string) (*aliospay.AliyunAliosPayTradeQueryAPIResponse, error) {
	var resp aliospay.AliyunAliosPayTradeQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
