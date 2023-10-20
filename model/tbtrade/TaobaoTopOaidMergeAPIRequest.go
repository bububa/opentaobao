package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopoaidmergeAPIRequest OAID订单合并 API请求
// taobao.top.oaid.merge
//
// 基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
type TaobaotopoaidmergeAPIRequest struct {
	model.Params
	// 合单请求列表，最多支持100个。
	_mergeList []OrderMerge
}

// NewTaobaotopoaidmergeRequest 初始化TaobaotopoaidmergeAPIRequest对象
func NewTaobaotopoaidmergeRequest() *TaobaotopoaidmergeAPIRequest {
	return &TaobaotopoaidmergeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopoaidmergeAPIRequest) GetApiMethodName() string {
	return "taobao.top.oaid.merge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopoaidmergeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopoaidmergeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMergeList is MergeList Setter
// 合单请求列表，最多支持100个。
func (r *TaobaotopoaidmergeAPIRequest) SetMergeList(_mergeList []OrderMerge) error {
	r._mergeList = _mergeList
	r.Set("merge_list", _mergeList)
	return nil
}

// GetMergeList MergeList Getter
func (r TaobaotopoaidmergeAPIRequest) GetMergeList() []OrderMerge {
	return r._mergeList
}
