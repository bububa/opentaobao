package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandoverparcelquery 获取交接单小包信息
// cainiao.global.handover.parcel.query
//
// 提供给ISV通过该接口查询小包信息
func Cainiaoglobalhandoverparcelquery(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandoverparcelqueryAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandoverparcelqueryAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandoverparcelqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
