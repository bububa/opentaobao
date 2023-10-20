package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousemerchantopenupdate 非融合店进件升级成融合店
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
func Alibabaalihousemerchantopenupdate(clt *core.SDKClient, req *alihouse.AlibabaalihousemerchantopenupdateAPIRequest, session string) (*alihouse.AlibabaalihousemerchantopenupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihousemerchantopenupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
