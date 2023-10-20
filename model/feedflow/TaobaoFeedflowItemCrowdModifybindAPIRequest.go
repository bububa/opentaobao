package feedflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCrowdModifybindAPIRequest 修改人群出价或状态 API请求
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
type TaobaoFeedflowItemCrowdModifybindAPIRequest struct {
	model.Params
	// 人群信息
	_crowds []CrowdDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemCrowdModifybindRequest 初始化TaobaoFeedflowItemCrowdModifybindAPIRequest对象
func NewTaobaoFeedflowItemCrowdModifybindRequest() *TaobaoFeedflowItemCrowdModifybindAPIRequest {
	return &TaobaoFeedflowItemCrowdModifybindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCrowdModifybindAPIRequest) Reset() {
	r._crowds = r._crowds[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.crowd.modifybind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrowds is Crowds Setter
// 人群信息
func (r *TaobaoFeedflowItemCrowdModifybindAPIRequest) SetCrowds(_crowds []CrowdDto) error {
	r._crowds = _crowds
	r.Set("crowds", _crowds)
	return nil
}

// GetCrowds Crowds Getter
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetCrowds() []CrowdDto {
	return r._crowds
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemCrowdModifybindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemCrowdModifybindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemCrowdModifybindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCrowdModifybindRequest()
	},
}

// GetTaobaoFeedflowItemCrowdModifybindRequest 从 sync.Pool 获取 TaobaoFeedflowItemCrowdModifybindAPIRequest
func GetTaobaoFeedflowItemCrowdModifybindAPIRequest() *TaobaoFeedflowItemCrowdModifybindAPIRequest {
	return poolTaobaoFeedflowItemCrowdModifybindAPIRequest.Get().(*TaobaoFeedflowItemCrowdModifybindAPIRequest)
}

// ReleaseTaobaoFeedflowItemCrowdModifybindAPIRequest 将 TaobaoFeedflowItemCrowdModifybindAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCrowdModifybindAPIRequest(v *TaobaoFeedflowItemCrowdModifybindAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCrowdModifybindAPIRequest.Put(v)
}
