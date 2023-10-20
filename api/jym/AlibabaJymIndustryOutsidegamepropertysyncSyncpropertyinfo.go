package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymindustryoutsidegamepropertysyncsyncpropertyinfo 外部上报游戏属性信息
// alibaba.jym.industry.outsidegamepropertysync.syncpropertyinfo
//
// 外部上报游戏属性信息
func Alibabajymindustryoutsidegamepropertysyncsyncpropertyinfo(clt *core.SDKClient, req *jym.AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIRequest, session string) (*jym.AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIResponse, error) {
	var resp jym.AlibabajymindustryoutsidegamepropertysyncsyncpropertyinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
