package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradessoldgetAPIRequest 查询卖家已卖出的交易数据（商家应用使用） API请求
// taobao.open.trades.sold.get
//
// 搜索当前会话用户作为卖家已卖出的交易数据（只能获取到三个月以内的交易信息）&lt;br/&gt;
// 1. 返回的数据结果是以订单的创建时间倒序排列的。&lt;br/&gt;
// 注意：type字段的说明，如果该字段不传，接口默认只查4种类型订单，非默认查询的订单是不返回。遇到订单查不到的情况的，通常都是这个原因造成。解决办法就是type加上订单类型就可正常返回了。&lt;br/&gt;
// 2.入参fields中传入buyer_nick ，才能返回buyer_open_id
type TaobaoopentradessoldgetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0,处方药未审核
	_fields string
	// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
	_startCreated string
	// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
	_endCreated string
	// 交易状态（<a href="http://open.taobao.com/doc/detail.htm?id=102856" target="_blank">查看可选值</a>），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
	_status string
	// 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade,auto_delivery,ec,cod,step 这5 种的交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>可选值：fixed(一口价)auction(拍卖)guarantee_trade(一口价、拍卖)step(分阶段付款，万人团，阶梯团订单）independent_simple_trade(旺店入门版交易)independent_shop_trade(旺店标准版交易)auto_delivery(自动发货)ec(直冲)cod(货到付款)game_equipment(游戏装备)shopex_trade(ShopEX交易)netcn_trade(万网交易)external_trade(统一外部交易)instant_trade (即时到账)b2c_cod(大商家货到付款)hotel_trade(酒店类型交易)super_market_trade(商超交易)super_market_cod_trade(商超货到付款交易)taohua(淘花网交易类型）waimai(外卖交易类型）o2o_offlinetrade（O2O交易）nopaid（即时到帐/趣味猜交易类型）step (万人团) eticket(电子凭证) tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。pre_auth_type(预授权0元购) lazada（获取lazada订单类型）
	_type string
	// 买家的openId
	_buyerOpenId string
	// 页码。取值范围:大于零的整数; 默认值:1
	_pageNo int64
	// 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
	_pageSize int64
	// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。
	_useHasNext bool
}

// NewTaobaoopentradessoldgetRequest 初始化TaobaoopentradessoldgetAPIRequest对象
func NewTaobaoopentradessoldgetRequest() *TaobaoopentradessoldgetAPIRequest {
	return &TaobaoopentradessoldgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradessoldgetAPIRequest) GetApiMethodName() string {
	return "taobao.open.trades.sold.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradessoldgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradessoldgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。rx_audit_status=0,处方药未审核
func (r *TaobaoopentradessoldgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoopentradessoldgetAPIRequest) GetFields() string {
	return r._fields
}

// SetStartCreated is StartCreated Setter
// 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoopentradessoldgetAPIRequest) SetStartCreated(_startCreated string) error {
	r._startCreated = _startCreated
	r.Set("start_created", _startCreated)
	return nil
}

// GetStartCreated StartCreated Getter
func (r TaobaoopentradessoldgetAPIRequest) GetStartCreated() string {
	return r._startCreated
}

// SetEndCreated is EndCreated Setter
// 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss
func (r *TaobaoopentradessoldgetAPIRequest) SetEndCreated(_endCreated string) error {
	r._endCreated = _endCreated
	r.Set("end_created", _endCreated)
	return nil
}

// GetEndCreated EndCreated Getter
func (r TaobaoopentradessoldgetAPIRequest) GetEndCreated() string {
	return r._endCreated
}

// SetStatus is Status Setter
// 交易状态（&lt;a href=&#34;http://open.taobao.com/doc/detail.htm?id=102856&#34; target=&#34;_blank&#34;&gt;查看可选值&lt;/a&gt;），默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
func (r *TaobaoopentradessoldgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoopentradessoldgetAPIRequest) GetStatus() string {
	return r._status
}

// SetType is Type Setter
// 交易类型列表，同时查询多种交易类型可用逗号分隔。&lt;span style=&#34;color:red;font-weight: bold;&#34;&gt;默认同时查询guarantee_trade,auto_delivery,ec,cod,step 这5 种的交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。&lt;/span&gt;可选值：fixed(一口价)auction(拍卖)guarantee_trade(一口价、拍卖)step(分阶段付款，万人团，阶梯团订单）independent_simple_trade(旺店入门版交易)independent_shop_trade(旺店标准版交易)auto_delivery(自动发货)ec(直冲)cod(货到付款)game_equipment(游戏装备)shopex_trade(ShopEX交易)netcn_trade(万网交易)external_trade(统一外部交易)instant_trade (即时到账)b2c_cod(大商家货到付款)hotel_trade(酒店类型交易)super_market_trade(商超交易)super_market_cod_trade(商超货到付款交易)taohua(淘花网交易类型）waimai(外卖交易类型）o2o_offlinetrade（O2O交易）nopaid（即时到帐/趣味猜交易类型）step (万人团) eticket(电子凭证) tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。pre_auth_type(预授权0元购) lazada（获取lazada订单类型）
func (r *TaobaoopentradessoldgetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoopentradessoldgetAPIRequest) GetType() string {
	return r._type
}

// SetBuyerOpenId is BuyerOpenId Setter
// 买家的openId
func (r *TaobaoopentradessoldgetAPIRequest) SetBuyerOpenId(_buyerOpenId string) error {
	r._buyerOpenId = _buyerOpenId
	r.Set("buyer_open_id", _buyerOpenId)
	return nil
}

// GetBuyerOpenId BuyerOpenId Getter
func (r TaobaoopentradessoldgetAPIRequest) GetBuyerOpenId() string {
	return r._buyerOpenId
}

// SetPageNo is PageNo Setter
// 页码。取值范围:大于零的整数; 默认值:1
func (r *TaobaoopentradessoldgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoopentradessoldgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100
func (r *TaobaoopentradessoldgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoopentradessoldgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetUseHasNext is UseHasNext Setter
// 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。
func (r *TaobaoopentradessoldgetAPIRequest) SetUseHasNext(_useHasNext bool) error {
	r._useHasNext = _useHasNext
	r.Set("use_has_next", _useHasNext)
	return nil
}

// GetUseHasNext UseHasNext Getter
func (r TaobaoopentradessoldgetAPIRequest) GetUseHasNext() bool {
	return r._useHasNext
}
