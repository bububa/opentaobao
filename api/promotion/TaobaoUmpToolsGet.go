package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpToolsGet 查询工具列表
// taobao.ump.tools.get
//
// 查询工具列表
func TaobaoUmpToolsGet(clt *core.SDKClient, req *promotion.TaobaoUmpToolsGetAPIRequest, session string) (*promotion.TaobaoUmpToolsGetAPIResponse, error) {
	var resp promotion.TaobaoUmpToolsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
