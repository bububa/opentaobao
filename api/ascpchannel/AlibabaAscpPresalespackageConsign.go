package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascppresalespackageconsign 预售预包尾款推单发货
// alibaba.ascp.presalespackage.consign
//
// 预售预包尾款发货后推单处理
func Alibabaascppresalespackageconsign(clt *core.SDKClient, req *ascpchannel.AlibabaascppresalespackageconsignAPIRequest, session string) (*ascpchannel.AlibabaascppresalespackageconsignAPIResponse, error) {
	var resp ascpchannel.AlibabaascppresalespackageconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
