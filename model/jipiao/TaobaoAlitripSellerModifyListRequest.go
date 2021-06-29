package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商订单】改签订单列表 APIRequest
taobao.alitrip.seller.modify.list

提供机票代理商查询改签订单列表
*/
type TaobaoAlitripSellerModifyListRequest struct {
    model.Params

    // 申请单ID
    applyId   int64 

    // 淘宝订单号
    orderId   int64 

    // 改签发起时间的查询结束日期 和 更新时间必选其一
    applyDateEnd   string 

    // 改签发起时间的查询开始日期 和 更新时间必选其一
    applyDateStart   string 

    // 页码
    currentPage   int64 

    // 乘客起飞时间的查询结束日期
    flyDateEnd   string 

    // 乘客起飞时间的查询开始日期
    flyDateStart   string 

    // 每页记录数
    pageSize   int64 

    // 1：初始状态，2：已改签成功，3：已拒绝，4：未付款（已回填退票费），5：已付款
    status   int64 

    // 记录修改结束时间  和 改签发起时间必选其一
    modifyDateEnd   string 

    // 记录修改起始时间 和 改签发起时间必选其一
    modifyDateStart   string 

}

func NewTaobaoAlitripSellerModifyListRequest() *TaobaoAlitripSellerModifyListRequest{
    return &TaobaoAlitripSellerModifyListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripSellerModifyListRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.modify.list"
}

func (r TaobaoAlitripSellerModifyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripSellerModifyListRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *TaobaoAlitripSellerModifyListRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoAlitripSellerModifyListRequest) SetApplyDateEnd(applyDateEnd string) error {
    r.applyDateEnd = applyDateEnd
    r.Set("apply_date_end", applyDateEnd)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetApplyDateEnd() string {
    return r.applyDateEnd
}

func (r *TaobaoAlitripSellerModifyListRequest) SetApplyDateStart(applyDateStart string) error {
    r.applyDateStart = applyDateStart
    r.Set("apply_date_start", applyDateStart)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetApplyDateStart() string {
    return r.applyDateStart
}

func (r *TaobaoAlitripSellerModifyListRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoAlitripSellerModifyListRequest) SetFlyDateEnd(flyDateEnd string) error {
    r.flyDateEnd = flyDateEnd
    r.Set("fly_date_end", flyDateEnd)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetFlyDateEnd() string {
    return r.flyDateEnd
}

func (r *TaobaoAlitripSellerModifyListRequest) SetFlyDateStart(flyDateStart string) error {
    r.flyDateStart = flyDateStart
    r.Set("fly_date_start", flyDateStart)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetFlyDateStart() string {
    return r.flyDateStart
}

func (r *TaobaoAlitripSellerModifyListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoAlitripSellerModifyListRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoAlitripSellerModifyListRequest) SetModifyDateEnd(modifyDateEnd string) error {
    r.modifyDateEnd = modifyDateEnd
    r.Set("modify_date_end", modifyDateEnd)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetModifyDateEnd() string {
    return r.modifyDateEnd
}

func (r *TaobaoAlitripSellerModifyListRequest) SetModifyDateStart(modifyDateStart string) error {
    r.modifyDateStart = modifyDateStart
    r.Set("modify_date_start", modifyDateStart)
    return nil
}

func (r TaobaoAlitripSellerModifyListRequest) GetModifyDateStart() string {
    return r.modifyDateStart
}

