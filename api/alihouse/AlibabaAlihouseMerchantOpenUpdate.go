package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantOpenUpdate 非融合店进件升级成融合店
// alibaba.alihouse.merchant.open.update
//
// 非融合店进件升级成融合店
func AlibabaAlihouseMerchantOpenUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantOpenUpdateAPIRequest, session string) (*alihouse.AlibabaAlihouseMerchantOpenUpdateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseMerchantOpenUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
