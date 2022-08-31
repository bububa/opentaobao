package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradePostageUpdate 修改交易邮费价格
// taobao.trade.postage.update
//
// 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
// &lt;br/&gt; &lt;span style=&#34;color:red&#34;&gt; API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750&lt;/span&gt;
func TaobaoTradePostageUpdate(clt *core.SDKClient, req *trade.TaobaoTradePostageUpdateAPIRequest, session string) (*trade.TaobaoTradePostageUpdateAPIResponse, error) {
	var resp trade.TaobaoTradePostageUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
