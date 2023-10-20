package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdModifyAPIRequest 覆盖单元下同类型定向人群 API请求
// taobao.feedflow.item.crowd.modify
//
// 覆盖单元下同类型定向人群
type TaobaoFeedflowItemCrowdModifyAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemCrowdModifyRequest 初始化TaobaoFeedflowItemCrowdModifyAPIRequest对象
func NewTaobaoFeedflowItemCrowdModifyRequest() *TaobaoFeedflowItemCrowdModifyAPIRequest {
	return &TaobaoFeedflowItemCrowdModifyAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdModifyAPIRequest) Reset() {
	r._crowds = r._crowds[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdModifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群信息
func (r *TaobaoFeedflowItemCrowdModifyAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaoFeedflowItemCrowdModifyAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdModifyAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemCrowdModifyAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemCrowdModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdModifyRequest()
	},
}

// GetTaobaoFeedflowItemCrowdModifyRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdModifyAPIRequest
func GetTaobaoFeedflowItemCrowdModifyAPIRequest() *TaobaoFeedflowItemCrowdModifyAPIRequest {
	return poolTaobaoFeedflowItemCrowdModifyAPIRequest.Get().(*TaobaoFeedflowItemCrowdModifyAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdModifyAPIRequest 将 TaobaoFeedflowItemCrowdModifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdModifyAPIRequest(v *TaobaoFeedflowItemCrowdModifyAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdModifyAPIRequest.Put(v)
}
