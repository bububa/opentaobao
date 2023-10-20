package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpRangeAdd 添加活动范围
// taobao.ump.range.add
//
// 指定某项活动中，某个商家的某些类型物品（指定商品或者别的）参加或者不参加活动。当活动详情的参与类型为部分或者部分不参加的时候，需要指定具体哪部分参加或者不参加，使用本接口完成操作。比如部分商品满就送，这里的range用来指定具体哪些商品参加满就送活动。
func TaobaoUmpRangeAdd(clt *core.SDKClient, req *promotion.TaobaoUmpRangeAddAPIRequest, resp *promotion.TaobaoUmpRangeAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
