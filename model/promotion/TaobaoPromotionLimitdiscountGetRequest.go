package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
限时打折查询 API请求
taobao.promotion.limitdiscount.get

分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。
*/
type TaobaoPromotionLimitdiscountGetRequest struct {
    model.Params
    // 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。
    limitDiscountId   int64
    // 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。
    status   string
    // 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。
    startTime   string
    // 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。
    endTime   string
    // 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。
    pageNumber   int64
}

// 初始化TaobaoPromotionLimitdiscountGetRequest对象
func NewTaobaoPromotionLimitdiscountGetRequest() *TaobaoPromotionLimitdiscountGetRequest{
    return &TaobaoPromotionLimitdiscountGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionLimitdiscountGetRequest) GetApiMethodName() string {
    return "taobao.promotion.limitdiscount.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionLimitdiscountGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LimitDiscountId Setter
// 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。
func (r *TaobaoPromotionLimitdiscountGetRequest) SetLimitDiscountId(limitDiscountId int64) error {
    r.limitDiscountId = limitDiscountId
    r.Set("limit_discount_id", limitDiscountId)
    return nil
}

// LimitDiscountId Getter
func (r TaobaoPromotionLimitdiscountGetRequest) GetLimitDiscountId() int64 {
    return r.limitDiscountId
}
// Status Setter
// 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。
func (r *TaobaoPromotionLimitdiscountGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoPromotionLimitdiscountGetRequest) GetStatus() string {
    return r.status
}
// StartTime Setter
// 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。
func (r *TaobaoPromotionLimitdiscountGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionLimitdiscountGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。
func (r *TaobaoPromotionLimitdiscountGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionLimitdiscountGetRequest) GetEndTime() string {
    return r.endTime
}
// PageNumber Setter
// 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。
func (r *TaobaoPromotionLimitdiscountGetRequest) SetPageNumber(pageNumber int64) error {
    r.pageNumber = pageNumber
    r.Set("page_number", pageNumber)
    return nil
}

// PageNumber Getter
func (r TaobaoPromotionLimitdiscountGetRequest) GetPageNumber() int64 {
    return r.pageNumber
}
