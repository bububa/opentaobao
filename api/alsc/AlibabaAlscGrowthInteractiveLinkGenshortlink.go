package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveLinkGenshortlink 短链接口
// alibaba.alsc.growth.interactive.link.genshortlink
//
// 短链接口
func AlibabaAlscGrowthInteractiveLinkGenshortlink(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveLinkGenshortlinkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
