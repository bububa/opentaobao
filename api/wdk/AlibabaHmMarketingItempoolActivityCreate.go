package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolActivityCreate 创建活动新接口
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func AlibabaHmMarketingItempoolActivityCreate(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolActivityCreateAPIRequest, resp *wdk.AlibabaHmMarketingItempoolActivityCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
