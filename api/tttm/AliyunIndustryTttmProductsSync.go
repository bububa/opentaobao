package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmProductsSync 天天特卖货品信息同步
// aliyun.industry.tttm.products.sync
//
// 天天特卖货品信息同步
func AliyunIndustryTttmProductsSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmProductsSyncAPIRequest, session string) (*tttm.AliyunIndustryTttmProductsSyncAPIResponse, error) {
	var resp tttm.AliyunIndustryTttmProductsSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
