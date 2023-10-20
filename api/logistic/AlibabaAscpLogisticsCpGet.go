package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaAscpLogisticsCpGet 快递公司资源列表查询接口
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
func AlibabaAscpLogisticsCpGet(clt *core.SDKClient, req *logistic.AlibabaAscpLogisticsCpGetAPIRequest, session string) (*logistic.AlibabaAscpLogisticsCpGetAPIResponse, error) {
	var resp logistic.AlibabaAscpLogisticsCpGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
