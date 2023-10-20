package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoNbaddAppointdeliverFeedbackcodes 服务质量反馈编码列表
// cainiao.nbadd.appointdeliver.feedbackcodes
//
// 服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
func CainiaoNbaddAppointdeliverFeedbackcodes(clt *core.SDKClient, req *cainiaolocker.CainiaoNbaddAppointdeliverFeedbackcodesAPIRequest, session string) (*cainiaolocker.CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse, error) {
	var resp cainiaolocker.CainiaoNbaddAppointdeliverFeedbackcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
