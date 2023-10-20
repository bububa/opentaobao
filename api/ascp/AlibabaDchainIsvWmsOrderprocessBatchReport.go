package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainisvwmsorderprocessbatchreport 仓作业信息批量同步
// alibaba.dchain.isv.wms.orderprocess.batch.report
//
// 仓作业信息批量同步
func Alibabadchainisvwmsorderprocessbatchreport(clt *core.SDKClient, req *ascp.AlibabadchainisvwmsorderprocessbatchreportAPIRequest, session string) (*ascp.AlibabadchainisvwmsorderprocessbatchreportAPIResponse, error) {
	var resp ascp.AlibabadchainisvwmsorderprocessbatchreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
