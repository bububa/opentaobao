package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// Alibabaeleenterpriseordernewgettrackinginfo 订单配送信息跟踪
// alibaba.ele.enterprise.ordernew.gettrackinginfo
//
// 订单配送信息跟踪
func Alibabaeleenterpriseordernewgettrackinginfo(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaeleenterpriseordernewgettrackinginfoAPIRequest, session string) (*eleenterpriseordernew.AlibabaeleenterpriseordernewgettrackinginfoAPIResponse, error) {
	var resp eleenterpriseordernew.AlibabaeleenterpriseordernewgettrackinginfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
