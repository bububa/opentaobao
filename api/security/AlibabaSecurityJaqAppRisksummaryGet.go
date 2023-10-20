package security

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/security"
)

// Alibabasecurityjaqapprisksummaryget 应用风险概要信息查询接口
// alibaba.security.jaq.app.risksummary.get
//
// 用户通过alibaba.security.jaq.app.risk.scan接口提交应用进行风险扫描后，用此接口获取风险概要信息，本接口不返回风险详细信息
func Alibabasecurityjaqapprisksummaryget(clt *core.SDKClient, req *security.AlibabasecurityjaqapprisksummarygetAPIRequest, session string) (*security.AlibabasecurityjaqapprisksummarygetAPIResponse, error) {
	var resp security.AlibabasecurityjaqapprisksummarygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
