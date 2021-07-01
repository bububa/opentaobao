package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeApplyorderGetsAPIRequest
获取采购申请单列表 API请求
tmall.channel.trade.applyorder.gets

分页查询采购申请单列表 */
type TmallChannelTradeApplyorderGetsAPIRequest struct {
	model.Params
	// 交易类型
	_tradeType int64
	// 申请单号
	_channelPurchaseApplyOrderNo string
	// 每页显示数量
	_pageSize int64
	// 查询第几页
	_pageNumber int64
	// 是否需要分页
	_needPagination bool
	// 审核状态列表
	_auditStatusList []int64
	// 分销商nick
	_distributorNick string
	// 渠道
	_channel int64
}

// NewTmallChannelTradeApplyorderGetsRequest 初始化TmallChannelTradeApplyorderGetsAPIRequest对象
func NewTmallChannelTradeApplyorderGetsRequest() *TmallChannelTradeApplyorderGetsAPIRequest {
	return &TmallChannelTradeApplyorderGetsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.applyorder.gets"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TradeType Setter
// 交易类型
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetTradeType(_tradeType int64) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// Get TradeType Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetTradeType() int64 {
	return r._tradeType
}

// Set is ChannelPurchaseApplyOrderNo Setter
// 申请单号
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetChannelPurchaseApplyOrderNo(_channelPurchaseApplyOrderNo string) error {
	r._channelPurchaseApplyOrderNo = _channelPurchaseApplyOrderNo
	r.Set("channel_purchase_apply_order_no", _channelPurchaseApplyOrderNo)
	return nil
}

// Get ChannelPurchaseApplyOrderNo Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetChannelPurchaseApplyOrderNo() string {
	return r._channelPurchaseApplyOrderNo
}

// Set is PageSize Setter
// 每页显示数量
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageNumber Setter
// 查询第几页
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// Get PageNumber Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}

// Set is NeedPagination Setter
// 是否需要分页
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetNeedPagination(_needPagination bool) error {
	r._needPagination = _needPagination
	r.Set("need_pagination", _needPagination)
	return nil
}

// Get NeedPagination Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetNeedPagination() bool {
	return r._needPagination
}

// Set is AuditStatusList Setter
// 审核状态列表
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetAuditStatusList(_auditStatusList []int64) error {
	r._auditStatusList = _auditStatusList
	r.Set("audit_status_list", _auditStatusList)
	return nil
}

// Get AuditStatusList Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetAuditStatusList() []int64 {
	return r._auditStatusList
}

// Set is DistributorNick Setter
// 分销商nick
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetDistributorNick(_distributorNick string) error {
	r._distributorNick = _distributorNick
	r.Set("distributor_nick", _distributorNick)
	return nil
}

// Get DistributorNick Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetDistributorNick() string {
	return r._distributorNick
}

// Set is Channel Setter
// 渠道
func (r *TmallChannelTradeApplyorderGetsAPIRequest) SetChannel(_channel int64) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r TmallChannelTradeApplyorderGetsAPIRequest) GetChannel() int64 {
	return r._channel
}
