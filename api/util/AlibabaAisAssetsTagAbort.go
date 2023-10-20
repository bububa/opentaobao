package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Alibabaaisassetstagabort 基础设施资产标签废弃
// alibaba.ais.assets.tag.abort
//
// 提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
func Alibabaaisassetstagabort(clt *core.SDKClient, req *util.AlibabaaisassetstagabortAPIRequest, session string) (*util.AlibabaaisassetstagabortAPIResponse, error) {
	var resp util.AlibabaaisassetstagabortAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
