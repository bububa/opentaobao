package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangindustrywaybillcreate 服务商开运单
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
func Alibabadchainaoxiangindustrywaybillcreate(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangindustrywaybillcreateAPIRequest, session string) (*ascp.AlibabadchainaoxiangindustrywaybillcreateAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangindustrywaybillcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
