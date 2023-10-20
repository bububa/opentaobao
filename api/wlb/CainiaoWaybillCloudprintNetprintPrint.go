package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Cainiaowaybillcloudprintnetprintprint 网络打印机打印接口
// cainiao.waybill.cloudprint.netprint.print
//
// 打印接口
func Cainiaowaybillcloudprintnetprintprint(clt *core.SDKClient, req *wlb.CainiaowaybillcloudprintnetprintprintAPIRequest, session string) (*wlb.CainiaowaybillcloudprintnetprintprintAPIResponse, error) {
	var resp wlb.CainiaowaybillcloudprintnetprintprintAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
