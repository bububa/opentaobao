package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingitempoolitemremoveasync 商品池删除商品
// alibaba.wdk.marketing.itempool.item.remove.async
//
// 新模型下删除商品
func Alibabawdkmarketingitempoolitemremoveasync(clt *core.SDKClient, req *wdk.AlibabawdkmarketingitempoolitemremoveasyncAPIRequest, session string) (*wdk.AlibabawdkmarketingitempoolitemremoveasyncAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingitempoolitemremoveasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
