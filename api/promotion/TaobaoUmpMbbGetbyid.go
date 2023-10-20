package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpmbbgetbyid 获取营销积木块
// taobao.ump.mbb.getbyid
//
// 根据积木块id获取营销积木块。
func Taobaoumpmbbgetbyid(clt *core.SDKClient, req *promotion.TaobaoumpmbbgetbyidAPIRequest, session string) (*promotion.TaobaoumpmbbgetbyidAPIResponse, error) {
	var resp promotion.TaobaoumpmbbgetbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
