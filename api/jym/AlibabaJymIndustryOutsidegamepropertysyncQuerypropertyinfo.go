package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymindustryoutsidegamepropertysyncquerypropertyinfo 外部查询游戏属性库属性信息
// alibaba.jym.industry.outsidegamepropertysync.querypropertyinfo
//
// 外部查询游戏属性库属性信息
func Alibabajymindustryoutsidegamepropertysyncquerypropertyinfo(clt *core.SDKClient, req *jym.AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIRequest, session string) (*jym.AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIResponse, error) {
	var resp jym.AlibabajymindustryoutsidegamepropertysyncquerypropertyinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
