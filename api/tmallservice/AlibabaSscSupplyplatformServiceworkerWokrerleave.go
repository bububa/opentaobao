package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscSupplyplatformServiceworkerWokrerleave 工人请假
// alibaba.ssc.supplyplatform.serviceworker.wokrerleave
//
// 工人请假
func AlibabaSscSupplyplatformServiceworkerWokrerleave(clt *core.SDKClient, req *tmallservice.AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIRequest, session string) (*tmallservice.AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponse, error) {
	var resp tmallservice.AlibabaSscSupplyplatformServiceworkerWokrerleaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
