package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaonbaddappointdeliverfeedbackcodes 服务质量反馈编码列表
// cainiao.nbadd.appointdeliver.feedbackcodes
//
// 服务质量反馈编码列表，建议获取数据后缓存在本地，定时刷新即可
func Cainiaonbaddappointdeliverfeedbackcodes(clt *core.SDKClient, req *cainiaolocker.CainiaonbaddappointdeliverfeedbackcodesAPIRequest, session string) (*cainiaolocker.CainiaonbaddappointdeliverfeedbackcodesAPIResponse, error) {
	var resp cainiaolocker.CainiaonbaddappointdeliverfeedbackcodesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
