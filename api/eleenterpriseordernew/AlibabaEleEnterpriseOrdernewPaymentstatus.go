package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewpaymentstatus 设置订单支付
// alibaba.ele.enterprise.ordernew.paymentstatus
//
// 设置订单支付成功
func Alibabaeleenterpriseordernewpaymentstatus(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewpaymentstatusAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewpaymentstatusAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewpaymentstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
