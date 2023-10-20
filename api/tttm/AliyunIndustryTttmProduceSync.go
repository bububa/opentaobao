package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmProduceSync 天天特卖生产进度同步
// aliyun.industry.tttm.produce.sync
//
// 天天特卖生产进度同步
func AliyunIndustryTttmProduceSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmProduceSyncAPIRequest, session string) (*tttm.AliyunIndustryTttmProduceSyncAPIResponse, error) {
	var resp tttm.AliyunIndustryTttmProduceSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
