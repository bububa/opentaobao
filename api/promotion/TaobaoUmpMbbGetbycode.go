package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumpmbbgetbycode 根据营销积木块代码获取积木块
// taobao.ump.mbb.getbycode
//
// 根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。
func Taobaoumpmbbgetbycode(clt *core.SDKClient, req *promotion.TaobaoumpmbbgetbycodeAPIRequest, session string) (*promotion.TaobaoumpmbbgetbycodeAPIResponse, error) {
	var resp promotion.TaobaoumpmbbgetbycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
