package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCreativeDeleteAPIRequest 信息流删除创意 API请求
// taobao.feedflow.item.creative.delete
//
// 信息流删除创意
type TaobaoFeedflowItemCreativeDeleteAPIRequest struct {
	model.Params
	// 创意id列表
	_creativeIdList []string
}

// NewTaobaoFeedflowItemCreativeDeleteRequest 初始化TaobaoFeedflowItemCreativeDeleteAPIRequest对象
func NewTaobaoFeedflowItemCreativeDeleteRequest() *TaobaoFeedflowItemCreativeDeleteAPIRequest {
	return &TaobaoFeedflowItemCreativeDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCreativeIdList is CreativeIdList Setter
// 创意id列表
func (r *TaobaoFeedflowItemCreativeDeleteAPIRequest) SetCreativeIdList(_creativeIdList []string) error {
	r._creativeIdList = _creativeIdList
	r.Set("creative_id_list", _creativeIdList)
	return nil
}

// GetCreativeIdList CreativeIdList Getter
func (r TaobaoFeedflowItemCreativeDeleteAPIRequest) GetCreativeIdList() []string {
	return r._creativeIdList
}
