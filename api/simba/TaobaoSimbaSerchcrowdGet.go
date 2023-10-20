package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSerchcrowdGet 根据推广单元id获取搜索溢价人群
// taobao.simba.serchcrowd.get
//
// 根据推广单元id获取搜索溢价人群
func TaobaoSimbaSerchcrowdGet(clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdGetAPIRequest, resp *simba.TaobaoSimbaSerchcrowdGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
