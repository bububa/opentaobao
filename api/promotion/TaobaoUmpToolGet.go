package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpToolGet 查询工具
// taobao.ump.tool.get
//
// 根据工具id获取一个工具对象
func TaobaoUmpToolGet(clt *core.SDKClient, req *promotion.TaobaoUmpToolGetAPIRequest, resp *promotion.TaobaoUmpToolGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
