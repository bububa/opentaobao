package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaServicecenterWorkcardRelatedskuQuery 查询工单关联的服务项
// alibaba.servicecenter.workcard.relatedsku.query
//
// 查询工单关联的服务项
func AlibabaServicecenterWorkcardRelatedskuQuery(clt *core.SDKClient, req *tmallsc.AlibabaServicecenterWorkcardRelatedskuQueryAPIRequest, session string) (*tmallsc.AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse, error) {
	var resp tmallsc.AlibabaServicecenterWorkcardRelatedskuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
