package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangwmsorderprocessreport 回传发货单流水通知
// alibaba.dchain.aoxiang.wms.orderprocess.report
//
// 回传发货单流水通知
func Alibabadchainaoxiangwmsorderprocessreport(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangwmsorderprocessreportAPIRequest, session string) (*ascp.AlibabadchainaoxiangwmsorderprocessreportAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangwmsorderprocessreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
