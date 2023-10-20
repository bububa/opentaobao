package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingbuygiftitemremoveasync 批量删除买赠商品
// alibaba.hm.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
func Alibabahmmarketingbuygiftitemremoveasync(clt *core.SDKClient, req *wdk.AlibabahmmarketingbuygiftitemremoveasyncAPIRequest, session string) (*wdk.AlibabahmmarketingbuygiftitemremoveasyncAPIResponse, error) {
	var resp wdk.AlibabahmmarketingbuygiftitemremoveasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
