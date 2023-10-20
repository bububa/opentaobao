package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// Taobaochanneltradeprepayofflineadd 渠道分销供应商上传线下流水预存款（增加）
// taobao.channel.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
func Taobaochanneltradeprepayofflineadd(clt *core.SDKClient, req *tmallchannel.TaobaochanneltradeprepayofflineaddAPIRequest, session string) (*tmallchannel.TaobaochanneltradeprepayofflineaddAPIResponse, error) {
	var resp tmallchannel.TaobaochanneltradeprepayofflineaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
