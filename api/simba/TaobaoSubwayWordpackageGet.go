package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayWordpackageGet 获取词包列表
// taobao.subway.wordpackage.get
//
// 获取流量智选、捡漏词包等词包列表
func TaobaoSubwayWordpackageGet(clt *core.SDKClient, req *simba.TaobaoSubwayWordpackageGetAPIRequest, resp *simba.TaobaoSubwayWordpackageGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
