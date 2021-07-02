package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOaidMergeAPIRequest OAID订单合并 API请求
// taobao.top.oaid.merge
//
// 基于OAID（收件人ID， Open Addressee ID)做订单合并，确保相同收件人信息的订单合并到相同组。
type TaobaoTopOaidMergeAPIRequest struct {
	model.Params
	// 合单请求列表，最多支持100个。
	_mergeList []OrderMerge
}

// NewTaobaoTopOaidMergeRequest 初始化TaobaoTopOaidMergeAPIRequest对象
func NewTaobaoTopOaidMergeRequest() *TaobaoTopOaidMergeAPIRequest {
	return &TaobaoTopOaidMergeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOaidMergeAPIRequest) GetApiMethodName() string {
	return "taobao.top.oaid.merge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOaidMergeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MergeList Setter
// 合单请求列表，最多支持100个。
func (r *TaobaoTopOaidMergeAPIRequest) SetMergeList(_mergeList []OrderMerge) error {
	r._mergeList = _mergeList
	r.Set("merge_list", _mergeList)
	return nil
}

// Get MergeList Getter
func (r TaobaoTopOaidMergeAPIRequest) GetMergeList() []OrderMerge {
	return r._mergeList
}
