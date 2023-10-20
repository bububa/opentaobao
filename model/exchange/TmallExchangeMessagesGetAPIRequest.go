package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangemessagesgetAPIRequest 查询换货订单留言列表 API请求
// tmall.exchange.messages.get
//
// 查询换货订单留言列表
type TmallexchangemessagesgetAPIRequest struct {
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

// NewTmallexchangemessagesgetRequest 初始化TmallexchangemessagesgetAPIRequest对象
func NewTmallexchangemessagesgetRequest() *TmallexchangemessagesgetAPIRequest {
	return &TmallexchangemessagesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangemessagesgetAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.messages.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangemessagesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangemessagesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperatorRoles is OperatorRoles Setter
// 留言创建角色。具体包括：卖家主账户(1)、卖家子账户(2)、小二(3)、买家(4)、系统(5)、系统超时(6)
func (r *TmallexchangemessagesgetAPIRequest) SetOperatorRoles(_operatorRoles []string) error {
	r._operatorRoles = _operatorRoles
	r.Set("operator_roles", _operatorRoles)
	return nil
}

// GetOperatorRoles OperatorRoles Getter
func (r TmallexchangemessagesgetAPIRequest) GetOperatorRoles() []string {
	return r._operatorRoles
}

// SetFields is Fields Setter
// 返回的字段。具体包括：id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
func (r *TmallexchangemessagesgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangemessagesgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TmallexchangemessagesgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallexchangemessagesgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangemessagesgetAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangemessagesgetAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TmallexchangemessagesgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TmallexchangemessagesgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
