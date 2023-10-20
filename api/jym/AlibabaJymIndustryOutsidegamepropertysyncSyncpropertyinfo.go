package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfo 外部上报游戏属性信息
// alibaba.jym.industry.outsidegamepropertysync.syncpropertyinfo
//
// 外部上报游戏属性信息
func AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfo(clt *core.SDKClient, req *jym.AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIRequest, session string) (*jym.AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse, error) {
	var resp jym.AlibabaJymIndustryOutsidegamepropertysyncSyncpropertyinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
