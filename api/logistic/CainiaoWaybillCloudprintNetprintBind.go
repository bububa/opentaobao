package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// CainiaoWaybillCloudprintNetprintBind 网络打印机绑定
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
func CainiaoWaybillCloudprintNetprintBind(clt *core.SDKClient, req *logistic.CainiaoWaybillCloudprintNetprintBindAPIRequest, session string) (*logistic.CainiaoWaybillCloudprintNetprintBindAPIResponse, error) {
	var resp logistic.CainiaoWaybillCloudprintNetprintBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
