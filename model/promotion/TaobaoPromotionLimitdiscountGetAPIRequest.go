package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionLimitdiscountGetAPIRequest 限时打折查询 API请求
// taobao.promotion.limitdiscount.get
//
// 分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。
type TaobaoPromotionLimitdiscountGetAPIRequest struct {
	model.Params
	// 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。
	_limitDiscountId int64
	// 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。
	_status string
	// 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。
	_startTime string
	// 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。
	_endTime string
	// 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。
	_pageNumber int64
}

// NewTaobaoPromotionLimitdiscountGetRequest 初始化TaobaoPromotionLimitdiscountGetAPIRequest对象
func NewTaobaoPromotionLimitdiscountGetRequest() *TaobaoPromotionLimitdiscountGetAPIRequest {
	return &TaobaoPromotionLimitdiscountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.limitdiscount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is LimitDiscountId Setter
// 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。
func (r *TaobaoPromotionLimitdiscountGetAPIRequest) SetLimitDiscountId(_limitDiscountId int64) error {
	r._limitDiscountId = _limitDiscountId
	r.Set("limit_discount_id", _limitDiscountId)
	return nil
}

// Get LimitDiscountId Getter
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetLimitDiscountId() int64 {
	return r._limitDiscountId
}

// Set is Status Setter
// 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。
func (r *TaobaoPromotionLimitdiscountGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetStatus() string {
	return r._status
}

// Set is StartTime Setter
// 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。
func (r *TaobaoPromotionLimitdiscountGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// Set is EndTime Setter
// 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。
func (r *TaobaoPromotionLimitdiscountGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is PageNumber Setter
// 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。
func (r *TaobaoPromotionLimitdiscountGetAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// Get PageNumber Getter
func (r TaobaoPromotionLimitdiscountGetAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
