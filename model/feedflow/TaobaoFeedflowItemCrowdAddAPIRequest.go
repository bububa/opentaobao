package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdAddAPIRequest 单品单元下，新增定向人群 API请求
// taobao.feedflow.item.crowd.add
//
// 单品单元下，新增定向人群
type TaobaoFeedflowItemCrowdAddAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemCrowdAddRequest 初始化TaobaoFeedflowItemCrowdAddAPIRequest对象
func NewTaobaoFeedflowItemCrowdAddRequest() *TaobaoFeedflowItemCrowdAddAPIRequest {
	return &TaobaoFeedflowItemCrowdAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdAddAPIRequest) Reset() {
	r._crowds = r._crowds[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群列表
func (r *TaobaoFeedflowItemCrowdAddAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemCrowdAddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemCrowdAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdAddRequest()
	},
}

// GetTaobaoFeedflowItemCrowdAddRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdAddAPIRequest
func GetTaobaoFeedflowItemCrowdAddAPIRequest() *TaobaoFeedflowItemCrowdAddAPIRequest {
	return poolTaobaoFeedflowItemCrowdAddAPIRequest.Get().(*TaobaoFeedflowItemCrowdAddAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdAddAPIRequest 将 TaobaoFeedflowItemCrowdAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdAddAPIRequest(v *TaobaoFeedflowItemCrowdAddAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdAddAPIRequest.Put(v)
}
