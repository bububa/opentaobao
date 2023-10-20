package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintclientinfoput 云打印客户端监控信息收集
// cainiao.cloudprint.clientinfo.put
//
// 云打印客户端监控信息收集
func Cainiaocloudprintclientinfoput(clt *core.SDKClient, req *waybill.CainiaocloudprintclientinfoputAPIRequest, session string) (*waybill.CainiaocloudprintclientinfoputAPIResponse, error) {
	var resp waybill.CainiaocloudprintclientinfoputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
