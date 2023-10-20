package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospaytradequery 查询支付结果接口
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
func Aliyunaliospaytradequery(clt *core.SDKClient, req *aliospay.AliyunaliospaytradequeryAPIRequest, session string) (*aliospay.AliyunaliospaytradequeryAPIResponse, error) {
	var resp aliospay.AliyunaliospaytradequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
