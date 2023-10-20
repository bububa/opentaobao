package eleenterpriseordernew

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterpriseordernew"
)

// AlibabaEleEnterpriseOrdernewGettrackinginfo 订单配送信息跟踪
// alibaba.ele.enterprise.ordernew.gettrackinginfo
//
// 订单配送信息跟踪
func AlibabaEleEnterpriseOrdernewGettrackinginfo(clt *core.SDKClient, req *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest, resp *eleenterpriseordernew.AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
