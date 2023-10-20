package feedflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemCreativeDeleteAPIRequest) Reset() {
	r._creativeIdList = r._creativeIdList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemCreativeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemCreativeDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemCreativeDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoFeedflowItemCreativeDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemCreativeDeleteRequest()
	},
}

// GetTaobaoFeedflowItemCreativeDeleteRequest 从 sync.Pool 获取 TaobaoFeedflowItemCreativeDeleteAPIRequest
func GetTaobaoFeedflowItemCreativeDeleteAPIRequest() *TaobaoFeedflowItemCreativeDeleteAPIRequest {
	return poolTaobaoFeedflowItemCreativeDeleteAPIRequest.Get().(*TaobaoFeedflowItemCreativeDeleteAPIRequest)
}

// ReleaseTaobaoFeedflowItemCreativeDeleteAPIRequest 将 TaobaoFeedflowItemCreativeDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemCreativeDeleteAPIRequest(v *TaobaoFeedflowItemCreativeDeleteAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemCreativeDeleteAPIRequest.Put(v)
}
