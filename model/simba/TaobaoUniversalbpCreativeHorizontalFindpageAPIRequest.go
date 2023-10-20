package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest 横向管理创意分页查询 API请求
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
type TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// creativeQueryVO
	_creativeQueryVO *CreativeQueryVo
}

// NewTaobaoUniversalbpCreativeHorizontalFindpageRequest 初始化TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest对象
func NewTaobaoUniversalbpCreativeHorizontalFindpageRequest() *TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest {
	return &TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) Reset() {
	r._topServiceContext = nil
	r._creativeQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.creative.horizontal.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCreativeQueryVO is CreativeQueryVO Setter
// creativeQueryVO
func (r *TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) SetCreativeQueryVO(_creativeQueryVO *CreativeQueryVo) error {
	r._creativeQueryVO = _creativeQueryVO
	r.Set("creative_query_v_o", _creativeQueryVO)
	return nil
}

// GetCreativeQueryVO CreativeQueryVO Getter
func (r TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) GetCreativeQueryVO() *CreativeQueryVo {
	return r._creativeQueryVO
}

var poolTaobaoUniversalbpCreativeHorizontalFindpageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpCreativeHorizontalFindpageRequest()
	},
}

// GetTaobaoUniversalbpCreativeHorizontalFindpageRequest 从 sync.Pool 获取 TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest
func GetTaobaoUniversalbpCreativeHorizontalFindpageAPIRequest() *TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest {
	return poolTaobaoUniversalbpCreativeHorizontalFindpageAPIRequest.Get().(*TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest)
}

// ReleaseTaobaoUniversalbpCreativeHorizontalFindpageAPIRequest 将 TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpCreativeHorizontalFindpageAPIRequest(v *TaobaoUniversalbpCreativeHorizontalFindpageAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpCreativeHorizontalFindpageAPIRequest.Put(v)
}
