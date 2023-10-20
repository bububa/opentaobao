package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverCloudprintGet 获取面单云打印数据
// cainiao.global.handover.cloudprint.get
//
// 提供给ISV通过该接口获取面单云打印数据
func CainiaoGlobalHandoverCloudprintGet(clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverCloudprintGetAPIRequest, session string) (*cainiaohandover.CainiaoGlobalHandoverCloudprintGetAPIResponse, error) {
	var resp cainiaohandover.CainiaoGlobalHandoverCloudprintGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
