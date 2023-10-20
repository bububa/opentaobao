package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainisvwmsorderprocessreport 仓作业信息同步
// alibaba.dchain.isv.wms.orderprocess.report
//
// 仓作业信息同步
func Alibabadchainisvwmsorderprocessreport(clt *core.SDKClient, req *ascp.AlibabadchainisvwmsorderprocessreportAPIRequest, session string) (*ascp.AlibabadchainisvwmsorderprocessreportAPIResponse, error) {
	var resp ascp.AlibabadchainisvwmsorderprocessreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
