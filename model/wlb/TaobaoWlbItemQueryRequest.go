package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商品 APIRequest
taobao.wlb.item.query

根据状态、卖家、SKU等信息查询商品列表
*/
type TaobaoWlbItemQueryRequest struct {
    model.Params

    // 是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理
    isSku   string 

    // 只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理
    status   string 

    // ITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理
    itemType   string 

    // 商品名称
    name   string 

    // 商品前台销售名字
    title   string 

    // 商家编码
    itemCode   string 

    // 父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品
    parentId   int64 

    // 当前页
    pageNo   int64 

    // 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
    pageSize   int64 

}

func NewTaobaoWlbItemQueryRequest() *TaobaoWlbItemQueryRequest{
    return &TaobaoWlbItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbItemQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.item.query"
}

func (r TaobaoWlbItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbItemQueryRequest) SetIsSku(isSku string) error {
    r.isSku = isSku
    r.Set("is_sku", isSku)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetIsSku() string {
    return r.isSku
}

func (r *TaobaoWlbItemQueryRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetStatus() string {
    return r.status
}

func (r *TaobaoWlbItemQueryRequest) SetItemType(itemType string) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetItemType() string {
    return r.itemType
}

func (r *TaobaoWlbItemQueryRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetName() string {
    return r.name
}

func (r *TaobaoWlbItemQueryRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoWlbItemQueryRequest) SetItemCode(itemCode string) error {
    r.itemCode = itemCode
    r.Set("item_code", itemCode)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetItemCode() string {
    return r.itemCode
}

func (r *TaobaoWlbItemQueryRequest) SetParentId(parentId int64) error {
    r.parentId = parentId
    r.Set("parent_id", parentId)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetParentId() int64 {
    return r.parentId
}

func (r *TaobaoWlbItemQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbItemQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbItemQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

