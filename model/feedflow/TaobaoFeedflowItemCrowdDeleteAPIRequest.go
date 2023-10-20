package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdDeleteAPIRequest 删除单品人群 API请求
// taobao.feedflow.item.crowd.delete
//
// 删除单品人群
type TaobaoFeedflowItemCrowdDeleteAPIRequest struct {
	model.Params
	// 人群结构
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemCrowdDeleteRequest 初始化TaobaoFeedflowItemCrowdDeleteAPIRequest对象
func NewTaobaoFeedflowItemCrowdDeleteRequest() *TaobaoFeedflowItemCrowdDeleteAPIRequest {
	return &TaobaoFeedflowItemCrowdDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdDeleteAPIRequest) Reset() {
	r._crowds = r._crowds[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群结构
func (r *TaobaoFeedflowItemCrowdDeleteAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdDeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemCrowdDeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemCrowdDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdDeleteRequest()
	},
}

// GetTaobaoFeedflowItemCrowdDeleteRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdDeleteAPIRequest
func GetTaobaoFeedflowItemCrowdDeleteAPIRequest() *TaobaoFeedflowItemCrowdDeleteAPIRequest {
	return poolTaobaoFeedflowItemCrowdDeleteAPIRequest.Get().(*TaobaoFeedflowItemCrowdDeleteAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdDeleteAPIRequest 将 TaobaoFeedflowItemCrowdDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdDeleteAPIRequest(v *TaobaoFeedflowItemCrowdDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdDeleteAPIRequest.Put(v)
}
