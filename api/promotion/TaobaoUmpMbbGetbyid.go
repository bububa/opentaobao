package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpMbbGetbyid 获取营销积木块
// taobao.ump.mbb.getbyid
//
// 根据积木块id获取营销积木块。
func TaobaoUmpMbbGetbyid(clt *core.SDKClient, req *promotion.TaobaoUmpMbbGetbyidAPIRequest, resp *promotion.TaobaoUmpMbbGetbyidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
