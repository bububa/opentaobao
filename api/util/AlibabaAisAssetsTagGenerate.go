package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Alibabaaisassetstaggenerate 基础设施资产标签生成
// alibaba.ais.assets.tag.generate
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
func Alibabaaisassetstaggenerate(clt *core.SDKClient, req *util.AlibabaaisassetstaggenerateAPIRequest, session string) (*util.AlibabaaisassetstaggenerateAPIResponse, error) {
	var resp util.AlibabaaisassetstaggenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
