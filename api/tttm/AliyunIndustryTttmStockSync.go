package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmStockSync 天天特卖库存同步接口
// aliyun.industry.tttm.stock.sync
//
// 天天特卖库存同步接口
func AliyunIndustryTttmStockSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmStockSyncAPIRequest, resp *tttm.AliyunIndustryTttmStockSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
