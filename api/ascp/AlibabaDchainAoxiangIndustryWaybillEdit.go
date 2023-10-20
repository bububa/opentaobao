package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangindustrywaybilledit 服务商编辑运单
// alibaba.dchain.aoxiang.industry.waybill.edit
//
// 服务商编辑运单
func Alibabadchainaoxiangindustrywaybilledit(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangindustrywaybilleditAPIRequest, session string) (*ascp.AlibabadchainaoxiangindustrywaybilleditAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangindustrywaybilleditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
