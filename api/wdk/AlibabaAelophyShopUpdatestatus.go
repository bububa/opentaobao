package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyShopUpdatestatus 更新渠道店营业状态
// alibaba.aelophy.shop.updatestatus
//
// 更新渠道店营业状态
func AlibabaAelophyShopUpdatestatus(clt *core.SDKClient, req *wdk.AlibabaAelophyShopUpdatestatusAPIRequest, session string) (*wdk.AlibabaAelophyShopUpdatestatusAPIResponse, error) {
	var resp wdk.AlibabaAelophyShopUpdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
