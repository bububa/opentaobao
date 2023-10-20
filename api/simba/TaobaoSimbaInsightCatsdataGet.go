package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbainsightcatsdataget 获取类目的大盘数据
// taobao.simba.insight.catsdata.get
//
// 根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。
func Taobaosimbainsightcatsdataget(clt *core.SDKClient, req *simba.TaobaosimbainsightcatsdatagetAPIRequest, session string) (*simba.TaobaosimbainsightcatsdatagetAPIResponse, error) {
	var resp simba.TaobaosimbainsightcatsdatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
