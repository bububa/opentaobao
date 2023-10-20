package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugtracetopyljgqueryupbilldetail 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】
// alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
func Alibabaalihealthdrugtracetopyljgqueryupbilldetail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugtracetopyljgqueryupbilldetailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugtracetopyljgqueryupbilldetailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugtracetopyljgqueryupbilldetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
