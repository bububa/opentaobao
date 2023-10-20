package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqianniutaskmetasgetAPIRequest 任务元查询接口 API请求
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
type TaobaoqianniutaskmetasgetAPIRequest struct {
	model.Params
	// 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
	_fields string
	// 排序字段。gmt_create,priority等
	_orderBy string
	// 升降序。asc为升，desc为降
	_orderType string
	// 任务类型
	_bizType string
	// 按关键字搜索
	_keyWord string
	// 客户端的版本信息
	_clientInfo string
	// 任务元ID，多个以逗号分离
	_metaIds string
	// 发起任务人的uid
	_senderUid int64
	// 分页数，最大100
	_pageSize int64
	// 当前页码
	_currentPage int64
	// 0为未完成。2为完成。4为取消。不填为所有
	_status int64
	// 接收人uid
	_receiverUid int64
}

// NewTaobaoqianniutaskmetasgetRequest 初始化TaobaoqianniutaskmetasgetAPIRequest对象
func NewTaobaoqianniutaskmetasgetRequest() *TaobaoqianniutaskmetasgetAPIRequest {
	return &TaobaoqianniutaskmetasgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqianniutaskmetasgetAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.taskmetas.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqianniutaskmetasgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqianniutaskmetasgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetFields() string {
	return r._fields
}

// SetOrderBy is OrderBy Setter
// 排序字段。gmt_create,priority等
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetOrderBy(_orderBy string) error {
	r._orderBy = _orderBy
	r.Set("order_by", _orderBy)
	return nil
}

// GetOrderBy OrderBy Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetOrderBy() string {
	return r._orderBy
}

// SetOrderType is OrderType Setter
// 升降序。asc为升，desc为降
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetBizType is BizType Setter
// 任务类型
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetBizType(_bizType string) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetBizType() string {
	return r._bizType
}

// SetKeyWord is KeyWord Setter
// 按关键字搜索
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// GetKeyWord KeyWord Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetKeyWord() string {
	return r._keyWord
}

// SetClientInfo is ClientInfo Setter
// 客户端的版本信息
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetClientInfo(_clientInfo string) error {
	r._clientInfo = _clientInfo
	r.Set("client_info", _clientInfo)
	return nil
}

// GetClientInfo ClientInfo Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetClientInfo() string {
	return r._clientInfo
}

// SetMetaIds is MetaIds Setter
// 任务元ID，多个以逗号分离
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetMetaIds(_metaIds string) error {
	r._metaIds = _metaIds
	r.Set("meta_ids", _metaIds)
	return nil
}

// GetMetaIds MetaIds Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetMetaIds() string {
	return r._metaIds
}

// SetSenderUid is SenderUid Setter
// 发起任务人的uid
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetSenderUid(_senderUid int64) error {
	r._senderUid = _senderUid
	r.Set("sender_uid", _senderUid)
	return nil
}

// GetSenderUid SenderUid Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetSenderUid() int64 {
	return r._senderUid
}

// SetPageSize is PageSize Setter
// 分页数，最大100
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页码
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetStatus is Status Setter
// 0为未完成。2为完成。4为取消。不填为所有
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetStatus() int64 {
	return r._status
}

// SetReceiverUid is ReceiverUid Setter
// 接收人uid
func (r *TaobaoqianniutaskmetasgetAPIRequest) SetReceiverUid(_receiverUid int64) error {
	r._receiverUid = _receiverUid
	r.Set("receiver_uid", _receiverUid)
	return nil
}

// GetReceiverUid ReceiverUid Getter
func (r TaobaoqianniutaskmetasgetAPIRequest) GetReceiverUid() int64 {
	return r._receiverUid
}
