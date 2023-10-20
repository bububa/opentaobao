package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfo 外部查询游戏属性库属性信息
// alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo
//
// 外部查询游戏属性库属性信息
func AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfo(clt *core.SDKClient, req *jym.AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIRequest, session string) (*jym.AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse, error) {
	var resp jym.AlibabaJymIndustryOutsidegamepropertysyncQuerypropertyinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
