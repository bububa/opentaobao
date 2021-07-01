package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* TaobaoUmpToolGet
查询工具
taobao.ump.tool.get

根据工具id获取一个工具对象 */
func TaobaoUmpToolGet(clt *core.SDKClient, req *promotion.TaobaoUmpToolGetAPIRequest, session string) (*promotion.TaobaoUmpToolGetAPIResponse, error) {
	var resp promotion.TaobaoUmpToolGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
