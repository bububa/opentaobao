package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRefundsApplyGetAPIRequest
查询买家申请的退款列表 API请求
taobao.refunds.apply.get

查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型 */
type TaobaoRefundsApplyGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee
	_fields []string
	// 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。<br/>WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) <br/>WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) <br/>WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) <br/>SELLER_REFUSE_BUYER(卖家拒绝退款) <br/>CLOSED(退款关闭) <br/>SUCCESS(退款成功)
	_status string
	// 卖家昵称
	_sellerNick string
	// 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。<br/>fixed(一口价) <br/>auction(拍卖) <br/>guarantee_trade(一口价、拍卖) <br/>independent_simple_trade(旺店入门版交易) <br/>independent_shop_trade(旺店标准版交易) <br/>auto_delivery(自动发货) <br/>ec(直冲) <br/>cod(货到付款) <br/>fenxiao(分销) <br/>game_equipment(游戏装备) <br/>shopex_trade(ShopEX交易) <br/>netcn_trade(万网交易) <br/>external_trade(统一外部交易)
	_type string
	// 页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始
	_pageNo int64
	// 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
	_pageSize int64
}

// New
