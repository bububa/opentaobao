package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单记录导出 APIRequest
taobao.vas.order.search

用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。<br/>建议用于查询前一日的历史记录，不适合用作实时数据查询。<br/>现在只能查询90天以内的数据<br/>该接口限制每分钟所有appkey调用总和只能有800次。
*/
type TaobaoVasOrderSearchRequest struct {
    model.Params

    // 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
    articleCode   string 

    // 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
    itemCode   string 

    // 淘宝会员名
    nick   string 

    // 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
    startCreated   string 

    // 订单创建时间（订购时间）结束值
    endCreated   string 

    // 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
    bizType   int64 

    // 订单号
    bizOrderId   int64 

    // 子订单号
    orderId   int64 

    // 一页包含的记录数
    pageSize   int64 

    // 页码
    pageNo   int64 

}

func NewTaobaoVasOrderSearchRequest() *TaobaoVasOrderSearchRequest{
    return &TaobaoVasOrderSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVasOrderSearchRequest) GetApiMethodName() string {
    return "taobao.vas.order.search"
}

func (r TaobaoVasOrderSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVasOrderSearchRequest) SetArticleCode(articleCode string) error {
    r.articleCode = articleCode
    r.Set("article_code", articleCode)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetArticleCode() string {
    return r.articleCode
}

func (r *TaobaoVasOrderSearchRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetItemCode() string {
    return r.itemCode
}

func (r *TaobaoVasOrderSearchRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoVasOrderSearchRequest) SetStartCreated(startCreated string) error {
    r.startCreated = startCreated
    r.Set("start_created", startCreated)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetStartCreated() string {
    return r.startCreated
}

func (r *TaobaoVasOrderSearchRequest) SetEndCreated(endCreated string) error {
    r.endCreated = endCreated
    r.Set("end_created", endCreated)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetEndCreated() string {
    return r.endCreated
}

func (r *TaobaoVasOrderSearchRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoVasOrderSearchRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *TaobaoVasOrderSearchRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVasOrderSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoVasOrderSearchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoVasOrderSearchRequest) GetPageNo() int64 {
    return r.pageNo
}

