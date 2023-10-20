package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgameliteplayavatarrecordreport Avatar形象保存地址回调
// alibaba.cgame.liteplay.avatar.record.report
//
// 新氢玩, 围观互动Avatar捏脸, 形象地址保存回调
func Alibabacgameliteplayavatarrecordreport(clt *core.SDKClient, req *cloudgame.AlibabacgameliteplayavatarrecordreportAPIRequest, session string) (*cloudgame.AlibabacgameliteplayavatarrecordreportAPIResponse, error) {
	var resp cloudgame.AlibabacgameliteplayavatarrecordreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
