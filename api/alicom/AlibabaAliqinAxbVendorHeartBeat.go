package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

/* AlibabaAliqinAxbVendorHeartBeat
供应商心跳上报接口
alibaba.aliqin.axb.vendor.heart.beat

供应商上报自己的心跳信息 */
func AlibabaAliqinAxbVendorHeartBeat(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorHeartBeatAPIRequest, session string) (*alicom.AlibabaAliqinAxbVendorHeartBeatAPIResponse, error) {
	var resp alicom.AlibabaAliqinAxbVendorHeartBeatAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
