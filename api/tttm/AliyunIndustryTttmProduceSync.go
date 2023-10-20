package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmProduceSync 天天特卖生产进度同步
// aliyun.industry.tttm.produce.sync
//
// 天天特卖生产进度同步
func AliyunIndustryTttmProduceSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmProduceSyncAPIRequest, resp *tttm.AliyunIndustryTttmProduceSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
