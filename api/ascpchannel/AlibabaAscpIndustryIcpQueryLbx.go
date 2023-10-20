package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpindustryicpquerylbx icp订单号查询lbx订单号
// alibaba.ascp.industry.icp.query.lbx
//
// 根据icp订单号查询lbx订单号
func Alibabaascpindustryicpquerylbx(clt *core.SDKClient, req *ascpchannel.AlibabaascpindustryicpquerylbxAPIRequest, session string) (*ascpchannel.AlibabaascpindustryicpquerylbxAPIResponse, error) {
	var resp ascpchannel.AlibabaascpindustryicpquerylbxAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
