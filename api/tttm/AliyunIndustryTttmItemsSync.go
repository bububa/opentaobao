package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// Aliyunindustrytttmitemssync 天天特卖商品信息同步
// aliyun.industry.tttm.items.sync
//
// 天天特卖商品信息同步
func Aliyunindustrytttmitemssync(clt *core.SDKClient, req *tttm.AliyunindustrytttmitemssyncAPIRequest, session string) (*tttm.AliyunindustrytttmitemssyncAPIResponse, error) {
	var resp tttm.AliyunindustrytttmitemssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
