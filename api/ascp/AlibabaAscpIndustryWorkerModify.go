package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryWorkerModify 送货入户并安装修改师傅信息
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
func AlibabaAscpIndustryWorkerModify(clt *core.SDKClient, req *ascp.AlibabaAscpIndustryWorkerModifyAPIRequest, session string) (*ascp.AlibabaAscpIndustryWorkerModifyAPIResponse, error) {
	var resp ascp.AlibabaAscpIndustryWorkerModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
