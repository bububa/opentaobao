package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinAxbVendorExceptionNoSync 中心化供应商异常号码状态同步接口
// alibaba.aliqin.axb.vendor.exception.no.sync
//
// 用于中心化供应商同步异常号码
func AlibabaAliqinAxbVendorExceptionNoSync(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorExceptionNoSyncAPIRequest, session string) (*alicom.AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse, error) {
	var resp alicom.AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
