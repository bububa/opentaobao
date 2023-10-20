package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolcreateactivity 添加商品池活动
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
func Alibabahmmarketingitempoolcreateactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolcreateactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolcreateactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
