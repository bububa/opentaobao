package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolqueryactivity 查找商品池活动
// alibaba.hm.marketing.itempool.queryactivity
//
// 查找商品池活动
func Alibabahmmarketingitempoolqueryactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolqueryactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolqueryactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
