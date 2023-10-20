package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbainsightwordsdataget 获取关键词的大盘数据
// taobao.simba.insight.wordsdata.get
//
// 获取关键词的详细数据
func Taobaosimbainsightwordsdataget(clt *core.SDKClient, req *simba.TaobaosimbainsightwordsdatagetAPIRequest, session string) (*simba.TaobaosimbainsightwordsdatagetAPIResponse, error) {
	var resp simba.TaobaosimbainsightwordsdatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
