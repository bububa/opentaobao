package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangereceivegetAPIRequest 卖家查询换货列表 API请求
// tmall.exchange.receive.get
//
// 卖家查询换货列表
type TmallexchangereceivegetAPIRequest struct {
	model.Params
	// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
	_fields []string
	// 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
	_disputeStatusArray []string
	// 退款单号ID列表，最多只能输入20个id
	_refundIdArray []string
	// 查询修改时间段的结束时间点
	_endGmtModifedTime string
	// 快递单号
	_logisticNo string
	// 买家昵称
	_buyerNick string
	// 查询申请时间段的开始时间点
	_startCreatedTime string
	// 查询申请时间段的结束时间点
	_endCreatedTime string
	// 买家的openId
	_buyerOpenUid string
	// 查询修改时间段的开始时间点
	_startGmtModifiedTime string
	// 每页条数
	_pageSize int64
	// 页码
	_pageNo int64
	// 正向订单号
	_bizOrderId int64
}

// NewTmallexchangereceivegetRequest 初始化TmallexchangereceivegetAPIRequest对象
func NewTmallexchangereceivegetRequest() *TmallexchangereceivegetAPIRequest {
	return &TmallexchangereceivegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangereceivegetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.receive.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangereceivegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangereceivegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick
func (r *TmallexchangereceivegetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangereceivegetAPIRequest) GetFields() []string {
	return r._fields
}

// SetDisputeStatusArray is DisputeStatusArray Setter
// 换货状态，具体包括：换货待处理(1), 待买家退货(2), 买家已退货，待收货(3),  换货关闭(4), 换货成功(5), 待买家修改(6), 待发出换货商品(12), 待买家收货(13), 请退款(14)
func (r *TmallexchangereceivegetAPIRequest) SetDisputeStatusArray(_disputeStatusArray []string) error {
	r._disputeStatusArray = _disputeStatusArray
	r.Set("dispute_status_array", _disputeStatusArray)
	return nil
}

// GetDisputeStatusArray DisputeStatusArray Getter
func (r TmallexchangereceivegetAPIRequest) GetDisputeStatusArray() []string {
	return r._disputeStatusArray
}

// SetRefundIdArray is RefundIdArray Setter
// 退款单号ID列表，最多只能输入20个id
func (r *TmallexchangereceivegetAPIRequest) SetRefundIdArray(_refundIdArray []string) error {
	r._refundIdArray = _refundIdArray
	r.Set("refund_id_array", _refundIdArray)
	return nil
}

// GetRefundIdArray RefundIdArray Getter
func (r TmallexchangereceivegetAPIRequest) GetRefundIdArray() []string {
	return r._refundIdArray
}

// SetEndGmtModifedTime is EndGmtModifedTime Setter
// 查询修改时间段的结束时间点
func (r *TmallexchangereceivegetAPIRequest) SetEndGmtModifedTime(_endGmtModifedTime string) error {
	r._endGmtModifedTime = _endGmtModifedTime
	r.Set("end_gmt_modifed_time", _endGmtModifedTime)
	return nil
}

// GetEndGmtModifedTime EndGmtModifedTime Getter
func (r TmallexchangereceivegetAPIRequest) GetEndGmtModifedTime() string {
	return r._endGmtModifedTime
}

// SetLogisticNo is LogisticNo Setter
// 快递单号
func (r *TmallexchangereceivegetAPIRequest) SetLogisticNo(_logisticNo string) error {
	r._logisticNo = _logisticNo
	r.Set("logistic_no", _logisticNo)
	return nil
}

// GetLogisticNo LogisticNo Getter
func (r TmallexchangereceivegetAPIRequest) GetLogisticNo() string {
	return r._logisticNo
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TmallexchangereceivegetAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TmallexchangereceivegetAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetStartCreatedTime is StartCreatedTime Setter
// 查询申请时间段的开始时间点
func (r *TmallexchangereceivegetAPIRequest) SetStartCreatedTime(_startCreatedTime string) error {
	r._startCreatedTime = _startCreatedTime
	r.Set("start_created_time", _startCreatedTime)
	return nil
}

// GetStartCreatedTime StartCreatedTime Getter
func (r TmallexchangereceivegetAPIRequest) GetStartCreatedTime() string {
	return r._startCreatedTime
}

// SetEndCreatedTime is EndCreatedTime Setter
// 查询申请时间段的结束时间点
func (r *TmallexchangereceivegetAPIRequest) SetEndCreatedTime(_endCreatedTime string) error {
	r._endCreatedTime = _endCreatedTime
	r.Set("end_created_time", _endCreatedTime)
	return nil
}

// GetEndCreatedTime EndCreatedTime Getter
func (r TmallexchangereceivegetAPIRequest) GetEndCreatedTime() string {
	return r._endCreatedTime
}

// SetBuyerOpenUid is BuyerOpenUid Setter
// 买家的openId
func (r *TmallexchangereceivegetAPIRequest) SetBuyerOpenUid(_buyerOpenUid string) error {
	r._buyerOpenUid = _buyerOpenUid
	r.Set("buyer_open_uid", _buyerOpenUid)
	return nil
}

// GetBuyerOpenUid BuyerOpenUid Getter
func (r TmallexchangereceivegetAPIRequest) GetBuyerOpenUid() string {
	return r._buyerOpenUid
}

// SetStartGmtModifiedTime is StartGmtModifiedTime Setter
// 查询修改时间段的开始时间点
func (r *TmallexchangereceivegetAPIRequest) SetStartGmtModifiedTime(_startGmtModifiedTime string) error {
	r._startGmtModifiedTime = _startGmtModifiedTime
	r.Set("start_gmt_modified_time", _startGmtModifiedTime)
	return nil
}

// GetStartGmtModifiedTime StartGmtModifiedTime Getter
func (r TmallexchangereceivegetAPIRequest) GetStartGmtModifiedTime() string {
	return r._startGmtModifiedTime
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TmallexchangereceivegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallexchangereceivegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TmallexchangereceivegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TmallexchangereceivegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetBizOrderId is BizOrderId Setter
// 正向订单号
func (r *TmallexchangereceivegetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallexchangereceivegetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
