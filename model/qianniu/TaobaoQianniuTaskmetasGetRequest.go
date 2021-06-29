package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
任务元查询接口 API请求
taobao.qianniu.taskmetas.get

任务元查询接口
*/
type TaobaoQianniuTaskmetasGetRequest struct {
    model.Params
    // 发起任务人的uid
    senderUid   int64
    // 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
    fields   string
    // 分页数，最大100
    pageSize   int64
    // 当前页码
    currentPage   int64
    // 排序字段。gmt_create,priority等
    orderBy   string
    // 升降序。asc为升，desc为降
    orderType   string
    // 0为未完成。2为完成。4为取消。不填为所有
    status   int64
    // 任务类型
    bizType   string
    // 按关键字搜索
    keyWord   string
    // 客户端的版本信息
    clientInfo   string
    // 接收人uid
    receiverUid   int64
    // 任务元ID，多个以逗号分离
    metaIds   string
}

// 初始化TaobaoQianniuTaskmetasGetRequest对象
func NewTaobaoQianniuTaskmetasGetRequest() *TaobaoQianniuTaskmetasGetRequest{
    return &TaobaoQianniuTaskmetasGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQianniuTaskmetasGetRequest) GetApiMethodName() string {
    return "taobao.qianniu.taskmetas.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQianniuTaskmetasGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SenderUid Setter
// 发起任务人的uid
func (r *TaobaoQianniuTaskmetasGetRequest) SetSenderUid(senderUid int64) error {
    r.senderUid = senderUid
    r.Set("sender_uid", senderUid)
    return nil
}

// SenderUid Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetSenderUid() int64 {
    return r.senderUid
}
// Fields Setter
// 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
func (r *TaobaoQianniuTaskmetasGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetFields() string {
    return r.fields
}
// PageSize Setter
// 分页数，最大100
func (r *TaobaoQianniuTaskmetasGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// CurrentPage Setter
// 当前页码
func (r *TaobaoQianniuTaskmetasGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// OrderBy Setter
// 排序字段。gmt_create,priority等
func (r *TaobaoQianniuTaskmetasGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

// OrderBy Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetOrderBy() string {
    return r.orderBy
}
// OrderType Setter
// 升降序。asc为升，desc为降
func (r *TaobaoQianniuTaskmetasGetRequest) SetOrderType(orderType string) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

// OrderType Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetOrderType() string {
    return r.orderType
}
// Status Setter
// 0为未完成。2为完成。4为取消。不填为所有
func (r *TaobaoQianniuTaskmetasGetRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetStatus() int64 {
    return r.status
}
// BizType Setter
// 任务类型
func (r *TaobaoQianniuTaskmetasGetRequest) SetBizType(bizType string) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetBizType() string {
    return r.bizType
}
// KeyWord Setter
// 按关键字搜索
func (r *TaobaoQianniuTaskmetasGetRequest) SetKeyWord(keyWord string) error {
    r.keyWord = keyWord
    r.Set("key_word", keyWord)
    return nil
}

// KeyWord Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetKeyWord() string {
    return r.keyWord
}
// ClientInfo Setter
// 客户端的版本信息
func (r *TaobaoQianniuTaskmetasGetRequest) SetClientInfo(clientInfo string) error {
    r.clientInfo = clientInfo
    r.Set("client_info", clientInfo)
    return nil
}

// ClientInfo Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetClientInfo() string {
    return r.clientInfo
}
// ReceiverUid Setter
// 接收人uid
func (r *TaobaoQianniuTaskmetasGetRequest) SetReceiverUid(receiverUid int64) error {
    r.receiverUid = receiverUid
    r.Set("receiver_uid", receiverUid)
    return nil
}

// ReceiverUid Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetReceiverUid() int64 {
    return r.receiverUid
}
// MetaIds Setter
// 任务元ID，多个以逗号分离
func (r *TaobaoQianniuTaskmetasGetRequest) SetMetaIds(metaIds string) error {
    r.metaIds = metaIds
    r.Set("meta_ids", metaIds)
    return nil
}

// MetaIds Getter
func (r TaobaoQianniuTaskmetasGetRequest) GetMetaIds() string {
    return r.metaIds
}
