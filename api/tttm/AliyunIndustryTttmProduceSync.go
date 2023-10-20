package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// Aliyunindustrytttmproducesync 天天特卖生产进度同步
// aliyun.industry.tttm.produce.sync
//
// 天天特卖生产进度同步
func Aliyunindustrytttmproducesync(clt *core.SDKClient, req *tttm.AliyunindustrytttmproducesyncAPIRequest, session string) (*tttm.AliyunindustrytttmproducesyncAPIResponse, error) {
	var resp tttm.AliyunindustrytttmproducesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
