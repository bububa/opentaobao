package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Taobaoumptoolget 查询工具
// taobao.ump.tool.get
//
// 根据工具id获取一个工具对象
func Taobaoumptoolget(clt *core.SDKClient, req *promotion.TaobaoumptoolgetAPIRequest, session string) (*promotion.TaobaoumptoolgetAPIResponse, error) {
	var resp promotion.TaobaoumptoolgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
