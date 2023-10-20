package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaAscpIndustryWorkerModify 送货入户并安装修改师傅信息
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
func AlibabaAscpIndustryWorkerModify(clt *core.SDKClient, req *ascp.AlibabaAscpIndustryWorkerModifyAPIRequest, resp *ascp.AlibabaAscpIndustryWorkerModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
