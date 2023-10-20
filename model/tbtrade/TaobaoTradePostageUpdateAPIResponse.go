package tbtrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradepostageupdateAPIResponse 修改交易邮费价格 API返回值
// taobao.trade.postage.update
//
// 修改订单邮费接口，通过传入订单编号和邮费价格，修改订单的邮费，返回修改时间modified,邮费post_fee,总费用total_fee。
// &lt;br/&gt; &lt;span style=&#34;color:red&#34;&gt; API取消加邮费功能通知：http://open.taobao.com/support/announcement_detail.htm?tid=24750&lt;/span&gt;
type TaobaotradepostageupdateAPIResponse struct {
	model.CommonResponse
	TaobaotradepostageupdateAPIResponseModel
}

// TaobaotradepostageupdateAPIResponseModel is 修改交易邮费价格 成功返回结果
type TaobaotradepostageupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_postage_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}
