package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivitySellerinfo 商家信息推送
// alibaba.lsy.crm.activity.sellerinfo
//
// 本地团商家信息推送
func AlibabaLsyCrmActivitySellerinfo(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivitySellerinfoAPIRequest, session string) (*tmallnr.AlibabaLsyCrmActivitySellerinfoAPIResponse, error) {
	var resp tmallnr.AlibabaLsyCrmActivitySellerinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
