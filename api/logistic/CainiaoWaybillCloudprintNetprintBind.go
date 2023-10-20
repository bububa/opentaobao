package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoWaybillCloudprintNetprintBind 网络打印机绑定
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
func CainiaoWaybillCloudprintNetprintBind(clt *core.SDKClient, req *logistic.CainiaoWaybillCloudprintNetprintBindAPIRequest, resp *logistic.CainiaoWaybillCloudprintNetprintBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
