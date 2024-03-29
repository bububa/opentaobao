package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivitiesListGetAPIRequest 营销活动列表查询 API请求
// taobao.ump.activities.list.get
//
// 按照营销活动id的列表ids，查询对应的营销活动列表。
type TaobaoUmpActivitiesListGetAPIRequest struct {
	model.Params
	// 营销活动id列表。
	_ids []int64
}

// NewTaobaoUmpActivitiesListGetRequest 初始化TaobaoUmpActivitiesListGetAPIRequest对象
func NewTaobaoUmpActivitiesListGetRequest() *TaobaoUmpActivitiesListGetAPIRequest {
	return &TaobaoUmpActivitiesListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpActivitiesListGetAPIRequest) Reset() {
	r._ids = r._ids[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivitiesListGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activities.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivitiesListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpActivitiesListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 营销活动id列表。
func (r *TaobaoUmpActivitiesListGetAPIRequest) SetIds(_ids []int64) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoUmpActivitiesListGetAPIRequest) GetIds() []int64 {
	return r._ids
}

var poolTaobaoUmpActivitiesListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpActivitiesListGetRequest()
	},
}

// GetTaobaoUmpActivitiesListGetRequest 从 sync.Pool 获取 TaobaoUmpActivitiesListGetAPIRequest
func GetTaobaoUmpActivitiesListGetAPIRequest() *TaobaoUmpActivitiesListGetAPIRequest {
	return poolTaobaoUmpActivitiesListGetAPIRequest.Get().(*TaobaoUmpActivitiesListGetAPIRequest)
}

// ReleaseTaobaoUmpActivitiesListGetAPIRequest 将 TaobaoUmpActivitiesListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpActivitiesListGetAPIRequest(v *TaobaoUmpActivitiesListGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpActivitiesListGetAPIRequest.Put(v)
}
