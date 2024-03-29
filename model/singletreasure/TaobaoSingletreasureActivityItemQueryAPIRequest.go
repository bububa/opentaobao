package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityItemQueryAPIRequest 查询活动下的优惠信息 API请求
// taobao.singletreasure.activity.item.query
//
// 分页查询活动下的商品优惠信息
type TaobaoSingletreasureActivityItemQueryAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
	// 页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTaobaoSingletreasureActivityItemQueryRequest 初始化TaobaoSingletreasureActivityItemQueryAPIRequest对象
func NewTaobaoSingletreasureActivityItemQueryRequest() *TaobaoSingletreasureActivityItemQueryAPIRequest {
	return &TaobaoSingletreasureActivityItemQueryAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) Reset() {
	r._activityId = 0
	r._pageSize = 0
	r._pageNumber = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动Id
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetPageSize is PageSize Setter
// 页大小
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TaobaoSingletreasureActivityItemQueryAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaoSingletreasureActivityItemQueryAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}

var poolTaobaoSingletreasureActivityItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityItemQueryRequest()
	},
}

// GetTaobaoSingletreasureActivityItemQueryRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityItemQueryAPIRequest
func GetTaobaoSingletreasureActivityItemQueryAPIRequest() *TaobaoSingletreasureActivityItemQueryAPIRequest {
	return poolTaobaoSingletreasureActivityItemQueryAPIRequest.Get().(*TaobaoSingletreasureActivityItemQueryAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityItemQueryAPIRequest 将 TaobaoSingletreasureActivityItemQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityItemQueryAPIRequest(v *TaobaoSingletreasureActivityItemQueryAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityItemQueryAPIRequest.Put(v)
}
