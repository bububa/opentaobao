package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaAscpLogisticsCpGet 快递公司资源列表查询接口
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
func AlibabaAscpLogisticsCpGet(clt *core.SDKClient, req *logistic.AlibabaAscpLogisticsCpGetAPIRequest, resp *logistic.AlibabaAscpLogisticsCpGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
