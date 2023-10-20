package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// Aliyunindustrytttmstocksync 天天特卖库存同步接口
// aliyun.industry.tttm.stock.sync
//
// 天天特卖库存同步接口
func Aliyunindustrytttmstocksync(clt *core.SDKClient, req *tttm.AliyunindustrytttmstocksyncAPIRequest, session string) (*tttm.AliyunindustrytttmstocksyncAPIResponse, error) {
	var resp tttm.AliyunindustrytttmstocksyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
