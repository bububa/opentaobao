package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolactivitycreate 创建活动新接口
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
func Alibabahmmarketingitempoolactivitycreate(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolactivitycreateAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolactivitycreateAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolactivitycreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
