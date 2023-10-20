package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewcancel 订单取消
// alibaba.ele.enterprise.ordernew.cancel
//
// 订单取消
func Alibabaeleenterpriseordernewcancel(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewcancelAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewcancelAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
