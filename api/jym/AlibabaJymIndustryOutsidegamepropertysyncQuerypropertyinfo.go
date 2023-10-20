package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfo 外部查询游戏属性库属性信息
// alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo
//
// 外部查询游戏属性库属性信息
func AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfo(clt *core.SDKClient, req *jym.AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest, resp *jym.AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
