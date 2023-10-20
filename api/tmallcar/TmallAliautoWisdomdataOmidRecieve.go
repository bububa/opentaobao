package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoWisdomdataOmidRecieve 大搜车车型参配数据接入
// tmall.aliauto.wisdomdata.omid.recieve
//
// 大搜车车型参配数据接入
func TmallAliautoWisdomdataOmidRecieve(clt *core.SDKClient, req *tmallcar.TmallAliautoWisdomdataOmidRecieveAPIRequest, resp *tmallcar.TmallAliautoWisdomdataOmidRecieveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
