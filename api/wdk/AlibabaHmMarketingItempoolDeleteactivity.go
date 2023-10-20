package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempooldeleteactivity 删除商品池活动
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
func Alibabahmmarketingitempooldeleteactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempooldeleteactivityAPIRequest, session string) (*wdk.AlibabahmmarketingitempooldeleteactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempooldeleteactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
