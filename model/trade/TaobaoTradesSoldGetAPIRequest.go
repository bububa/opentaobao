package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradesSoldGetAPIRequest
查询卖家已卖出的交易数据（根据创建时间） API请求
taobao.trades.sold.get

搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）
<br/>1. 返回的数据结果是以订单的创建时间倒序排列的。
<br/>2. 返回的数据结果只包含了订单的部分数据，可通过taobao.trade.fullinfo.get获取订单详情。
<br/>注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，解决办法就是type加上订单类型就可正常返回了。
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong> */
type TaobaoTradesSoldGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0,处方药未审核
	_fields string
	// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
	_startCreated string
	// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
	_endCreated string
	// 交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
	_status string
	// 买家昵称
	_buyerNick string
	// 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade,auto_delivery,ec,cod,step 这5 种的交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>可选值：fixed(一口价)auction(拍卖)guarantee_trade(一口价、拍卖)step(分阶段付款，万人团，阶梯团订单）independent_simple_trade(旺店入门版交易)independent_shop_trade(旺店标准版交易)auto_delivery(自动发货)ec(直冲)cod(货到付款)game_equipment(游戏装备)shopex_trade(ShopEX交易)netcn_trade(万网交易)external_trade(统一外部交易)instant_trade (即时到账)b2c_cod(大商家货到付款)hotel_trade(酒店类型交易)super_market_trade(商超交易)super_market_cod_trade(商超货到付款交易)taohua(淘花网交易类型）waimai(外卖交易类型）o2o_offlinetrade（O2O交易）nopaid（即时到帐/趣味猜交易类型）step (万人团) eticket(电子凭证) tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。pre_auth_type(预授权0元购) lazada（获取lazada订单类型）
	_type string
	// 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型
	_extType string
	// 评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>可选值：RATE_UNBUYER(买家未评)RATE_UNSELLER(卖家未评)RATE_BUYER_UNSELLER(买家已评，卖家未评)RATE_UNBUYER_SELLER(买家未评，卖家已评)RATE_BUYER_SELLER(买家已评,卖家已评)
	_rateStatus string
	// 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充）
	_tag string
	// 页码。默认值:1，可填整数，通过传入page_no来控制获取的页数，总页数=total_results÷page_size
	_pageNo int64
	// 每页条数。 默认值:40;最大值:100，可填整数。通过page_no和page_size组合多次调用实现翻页获取全量数据。
	_pageSize int64
	// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。
	_useHasNext bool
	// 买家的openId
	_buyerOpenId string
}

// New
