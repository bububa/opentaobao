package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbainsightwordspricedataget 关键词按竞价区间的分布数据
// taobao.simba.insight.wordspricedata.get
//
// 获取关键词按竞价区间进行细分的数据
func Taobaosimbainsightwordspricedataget(clt *core.SDKClient, req *simba.TaobaosimbainsightwordspricedatagetAPIRequest, session string) (*simba.TaobaosimbainsightwordspricedatagetAPIResponse, error) {
	var resp simba.TaobaosimbainsightwordspricedatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
