package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskmetasGetAPIRequest 任务元查询接口 API请求
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
type TaobaoQianniuTaskmetasGetAPIRequest struct {
	model.Params
	// 发起任务人的uid
	_senderUid int64
	// 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
	_fields string
	// 分页数，最大100
	_pageSize int64
	// 当前页码
	_currentPage int64
	// 排序字段。gmt_create,priority等
	_orderBy string
	// 升降序。asc为升，desc为降
	_orderType string
	// 0为未完成。2为完成。4为取消。不填为所有
	_status int64
	// 任务类型
	_bizType string
	// 按关键字搜索
	_keyWord string
	// 客户端的版本信息
	_clientInfo string
	// 接收人uid
	_receiverUid int64
	// 任务元ID，多个以逗号分离
	_metaIds string
}

// NewTaobaoQianniuTaskmetasGetRequest 初始化TaobaoQianniuTaskmetasGetAPIRequest对象
func NewTaobaoQianniuTaskmetasGetRequest() *TaobaoQianniuTaskmetasGetAPIRequest {
	return &TaobaoQianniuTaskmetasGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.taskmetas.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSenderUid is SenderUid Setter
// 发起任务人的uid
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetSenderUid(_senderUid int64) error {
	r._senderUid = _senderUid
	r.Set("sender_uid", _senderUid)
	return nil
}

// GetSenderUid SenderUid Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetSenderUid() int64 {
	return r._senderUid
}

// SetFields is Fields Setter
// 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetFields() string {
	return r._fields
}

// SetPageSize is PageSize Setter
// 分页数，最大100
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页码
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetOrderBy is OrderBy Setter
// 排序字段。gmt_create,priority等
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetOrderType is OrderType Setter
// 升降序。asc为升，desc为降
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetStatus is Status Setter
// 0为未完成。2为完成。4为取消。不填为所有
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetStatus() int64 {
	return r._status
}

// SetBizType is BizType Setter
// 任务类型
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetBizType() string {
	return r._bizType
}

// SetKeyWord is KeyWord Setter
// 按关键字搜索
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// GetKeyWord KeyWord Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetKeyWord() string {
	return r._keyWord
}

// SetClientInfo is ClientInfo Setter
// 客户端的版本信息
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetClientInfo(_clientInfo string) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetClientInfo() string {
	return r._clientInfo
}

// SetReceiverUid is ReceiverUid Setter
// 接收人uid
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetReceiverUid(_receiverUid int64) error {
	r._receiverUid = _receiverUid
	r.Set("receiver_uid", _receiverUid)
	return nil
}

// GetReceiverUid ReceiverUid Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetReceiverUid() int64 {
	return r._receiverUid
}

// SetMetaIds is MetaIds Setter
// 任务元ID，多个以逗号分离
func (r *TaobaoQianniuTaskmetasGetAPIRequest) SetMetaIds(_metaIds string) error {
	r._metaIds = _metaIds
	r.Set("meta_ids", _metaIds)
	return nil
}

// GetMetaIds MetaIds Getter
func (r TaobaoQianniuTaskmetasGetAPIRequest) GetMetaIds() string {
	return r._metaIds
}
