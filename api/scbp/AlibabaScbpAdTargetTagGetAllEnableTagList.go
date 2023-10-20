package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadtargettaggetallenabletaglist 查询所有可添加标签信息
// alibaba.scbp.ad.target.tag.get.all.enable.tag.list
//
// 查询标签数据
func Alibabascbpadtargettaggetallenabletaglist(clt *core.SDKClient, req *scbp.AlibabascbpadtargettaggetallenabletaglistAPIRequest, session string) (*scbp.AlibabascbpadtargettaggetallenabletaglistAPIResponse, error) {
	var resp scbp.AlibabascbpadtargettaggetallenabletaglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
