package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// Aliyunindustrytttmproductssync 天天特卖货品信息同步
// aliyun.industry.tttm.products.sync
//
// 天天特卖货品信息同步
func Aliyunindustrytttmproductssync(clt *core.SDKClient, req *tttm.AliyunindustrytttmproductssyncAPIRequest, session string) (*tttm.AliyunindustrytttmproductssyncAPIResponse, error) {
	var resp tttm.AliyunindustrytttmproductssyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
