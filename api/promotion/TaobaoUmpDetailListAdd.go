package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpDetailListAdd 营销详情添加
// taobao.ump.detail.list.add
//
// 批量添加营销活动。替代单条添加营销详情的的API。此接口适用针对某个营销活动的多档设置，会按顺序插入detail。若在整个事务过程中出现断点，会将已插入完成的detail_id返回，注意记录这些id，并将其删除，会对交易过程中的优惠产生影响。
func TaobaoUmpDetailListAdd(clt *core.SDKClient, req *promotion.TaobaoUmpDetailListAddAPIRequest, resp *promotion.TaobaoUmpDetailListAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
