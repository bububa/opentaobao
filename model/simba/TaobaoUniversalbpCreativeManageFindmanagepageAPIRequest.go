package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest 创意库查询创意列表 API请求
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
type TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// creativeQueryVO
	_creativeQueryVO *CreativeQueryVo
}

// NewTaobaoUniversalbpCreativeManageFindmanagepageRequest 初始化TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest对象
func NewTaobaoUniversalbpCreativeManageFindmanagepageRequest() *TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest {
	return &TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) Reset() {
	r._topServiceContext = nil
	r._creativeQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.creative.manage.findmanagepage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetCreativeQueryVO is CreativeQueryVO Setter
// creativeQueryVO
func (r *TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) SetCreativeQueryVO(_creativeQueryVO *CreativeQueryVo) error {
	r._creativeQueryVO = _creativeQueryVO
	r.Set("creative_query_v_o", _creativeQueryVO)
	return nil
}

// GetCreativeQueryVO CreativeQueryVO Getter
func (r TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) GetCreativeQueryVO() *CreativeQueryVo {
	return r._creativeQueryVO
}

var poolTaobaoUniversalbpCreativeManageFindmanagepageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpCreativeManageFindmanagepageRequest()
	},
}

// GetTaobaoUniversalbpCreativeManageFindmanagepageRequest 从 sync.Pool 获取 TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest
func GetTaobaoUniversalbpCreativeManageFindmanagepageAPIRequest() *TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest {
	return poolTaobaoUniversalbpCreativeManageFindmanagepageAPIRequest.Get().(*TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest)
}

// ReleaseTaobaoUniversalbpCreativeManageFindmanagepageAPIRequest 将 TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpCreativeManageFindmanagepageAPIRequest(v *TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpCreativeManageFindmanagepageAPIRequest.Put(v)
}
