package shenjing

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shenjing"
)

// AlibabaIbCampusTmiQuery IB智慧园区-查询TMI流水
// alibaba.ib.campus.tmi.query
//
// 获取特定银行账户的银行流水
func AlibabaIbCampusTmiQuery(clt *core.SDKClient, req *shenjing.AlibabaIbCampusTmiQueryAPIRequest, session string) (*shenjing.AlibabaIbCampusTmiQueryAPIResponse, error) {
	var resp shenjing.AlibabaIbCampusTmiQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
