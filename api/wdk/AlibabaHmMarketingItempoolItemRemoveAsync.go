package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingitempoolitemremoveasync 商品池删除商品
// alibaba.hm.marketing.itempool.item.remove.async
//
// 新模型下删除商品
func Alibabahmmarketingitempoolitemremoveasync(clt *core.SDKClient, req *wdk.AlibabahmmarketingitempoolitemremoveasyncAPIRequest, session string) (*wdk.AlibabahmmarketingitempoolitemremoveasyncAPIResponse, error) {
	var resp wdk.AlibabahmmarketingitempoolitemremoveasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
