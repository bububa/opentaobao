package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

/* TmallNrtPayMerchantStallSigningModify
三级商户进件修改
tmall.nrt.pay.merchant.stall.signing.modify

三级商户进件修改 */
func TmallNrtPayMerchantStallSigningModify(clt *core.SDKClient, req *nrt.TmallNrtPayMerchantStallSigningModifyAPIRequest, session string) (*nrt.TmallNrtPayMerchantStallSigningModifyAPIResponse, error) {
	var resp nrt.TmallNrtPayMerchantStallSigningModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
