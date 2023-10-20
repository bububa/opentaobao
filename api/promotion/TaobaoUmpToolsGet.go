package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumptoolsget 查询工具列表
// taobao.ump.tools.get
//
// 查询工具列表
func Taobaoumptoolsget(clt *core.SDKClient, req *promotion.TaobaoumptoolsgetAPIRequest, session string) (*promotion.TaobaoumptoolsgetAPIResponse, error) {
	var resp promotion.TaobaoumptoolsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
