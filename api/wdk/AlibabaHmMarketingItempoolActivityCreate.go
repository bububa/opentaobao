package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItempoolActivityCreate 创建活动新接口
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func AlibabaHmMarketingItempoolActivityCreate(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItempoolActivityCreateAPIRequest, session string) (*wdk.AlibabaHmMarketingItempoolActivityCreateAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingItempoolActivityCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
