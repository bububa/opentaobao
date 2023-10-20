package baichuanctg

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuanctg"
)

// AlibabaBaichuanCtgToutiaoContent 微博输出头条数据
// alibaba.baichuan.ctg.toutiao.content
//
// 百川头条内容获取
func AlibabaBaichuanCtgToutiaoContent(clt *core.SDKClient, req *baichuanctg.AlibabaBaichuanCtgToutiaoContentAPIRequest, resp *baichuanctg.AlibabaBaichuanCtgToutiaoContentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
