package exchange

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeMessagesGetAPIRequest 查询换货订单留言列表 API请求
// tmall.exchange.messages.get
//
// 查询换货订单留言列表
type TmallExchangeMessagesGetAPIRequest struct {
	model.Params
	// 留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6)
	_operatorRoles []string
	// 返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
	_fields []string
	// 每页条数
	_pageSize int64
	// 换货单号ID
	_disputeId int64
	// 页码
	_pageNo int64
}

// NewTmallExchangeMessagesGetRequest 初始化TmallExchangeMessagesGetAPIRequest对象
func NewTmallExchangeMessagesGetRequest() *TmallExchangeMessagesGetAPIRequest {
	return &TmallExchangeMessagesGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallExchangeMessagesGetAPIRequest) Reset() {
	r._operatorRoles = r._operatorRoles[:0]
	r._fields = r._fields[:0]
	r._pageSize = 0
	r._disputeId = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeMessagesGetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.messages.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeMessagesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallExchangeMessagesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperatorRoles is OperatorRoles Setter
// 留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6)
func (r *TmallExchangeMessagesGetAPIRequest) SetOperatorRoles(_operatorRoles []string) error {
	r._operatorRoles = _operatorRoles
	r.Set("operator_roles", _operatorRoles)
	return nil
}

// GetOperatorRoles OperatorRoles Getter
func (r TmallExchangeMessagesGetAPIRequest) GetOperatorRoles() []string {
	return r._operatorRoles
}

// SetFields is Fields Setter
// 返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
func (r *TmallExchangeMessagesGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeMessagesGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TmallExchangeMessagesGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallExchangeMessagesGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeMessagesGetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeMessagesGetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TmallExchangeMessagesGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TmallExchangeMessagesGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTmallExchangeMessagesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallExchangeMessagesGetRequest()
	},
}

// GetTmallExchangeMessagesGetRequest 从 sync.Pool 获取 TmallExchangeMessagesGetAPIRequest
func GetTmallExchangeMessagesGetAPIRequest() *TmallExchangeMessagesGetAPIRequest {
	return poolTmallExchangeMessagesGetAPIRequest.Get().(*TmallExchangeMessagesGetAPIRequest)
}

// ReleaseTmallExchangeMessagesGetAPIRequest 将 TmallExchangeMessagesGetAPIRequest 放入 sync.Pool
func ReleaseTmallExchangeMessagesGetAPIRequest(v *TmallExchangeMessagesGetAPIRequest) {
	v.Reset()
	poolTmallExchangeMessagesGetAPIRequest.Put(v)
}
