package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaSscServicecenterServicestoreQuery 根据天猫id查询门店信息
// alibaba.ssc.servicecenter.servicestore.query
//
// 根据天猫id查询门店信息
func AlibabaSscServicecenterServicestoreQuery(clt *core.SDKClient, req *tmallservice.AlibabaSscServicecenterServicestoreQueryAPIRequest, session string) (*tmallservice.AlibabaSscServicecenterServicestoreQueryAPIResponse, error) {
	var resp tmallservice.AlibabaSscServicecenterServicestoreQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
