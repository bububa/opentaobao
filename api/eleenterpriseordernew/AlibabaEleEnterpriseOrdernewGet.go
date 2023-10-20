package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewget 查询订单详情
// alibaba.ele.enterprise.ordernew.get
//
// 查询订单详情
func Alibabaeleenterpriseordernewget(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewgetAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewgetAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
