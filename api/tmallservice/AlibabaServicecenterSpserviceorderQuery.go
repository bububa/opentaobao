package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* AlibabaServicecenterSpserviceorderQuery
服务供应链服务单查询
alibaba.servicecenter.spserviceorder.query

服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务 */
func AlibabaServicecenterSpserviceorderQuery(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterSpserviceorderQueryAPIRequest, session string) (*tmallservice.AlibabaServicecenterSpserviceorderQueryAPIResponse, error) {
	var resp tmallservice.AlibabaServicecenterSpserviceorderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
