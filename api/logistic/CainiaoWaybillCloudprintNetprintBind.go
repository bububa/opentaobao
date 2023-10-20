package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Cainiaowaybillcloudprintnetprintbind 网络打印机绑定
// cainiao.waybill.cloudprint.netprint.bind
//
// 绑定打印机接口
func Cainiaowaybillcloudprintnetprintbind(clt *core.SDKClient, req *logistic.CainiaowaybillcloudprintnetprintbindAPIRequest, session string) (*logistic.CainiaowaybillcloudprintnetprintbindAPIResponse, error) {
	var resp logistic.CainiaowaybillcloudprintnetprintbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
