package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// Taobaochanneltradeprepayofflinereduce 渠道分销供应商上传线下流水预存款（减少）
// taobao.channel.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
func Taobaochanneltradeprepayofflinereduce(clt *core.SDKClient, req *tmallchannel.TaobaochanneltradeprepayofflinereduceAPIRequest, session string) (*tmallchannel.TaobaochanneltradeprepayofflinereduceAPIResponse, error) {
	var resp tmallchannel.TaobaochanneltradeprepayofflinereduceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
