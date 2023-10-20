package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmItemsSync 天天特卖商品信息同步
// aliyun.industry.tttm.items.sync
//
// 天天特卖商品信息同步
func AliyunIndustryTttmItemsSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmItemsSyncAPIRequest, resp *tttm.AliyunIndustryTttmItemsSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
