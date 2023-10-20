package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest 查看资源包列表 API请求
// taobao.universalbp.adzone.horizontal.findpage
//
// 查看已存在的计划上设置的资源包列表
type TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// adzoneRefQueryVO
	_adzoneRefQueryVO *AdzoneRefQueryVo
}

// NewTaobaoUniversalbpAdzoneHorizontalFindpageRequest 初始化TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest对象
func NewTaobaoUniversalbpAdzoneHorizontalFindpageRequest() *TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest {
	return &TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) Reset() {
	r._topServiceContext = nil
	r._adzoneRefQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.adzone.horizontal.findpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetAdzoneRefQueryVO is AdzoneRefQueryVO Setter
// adzoneRefQueryVO
func (r *TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) SetAdzoneRefQueryVO(_adzoneRefQueryVO *AdzoneRefQueryVo) error {
	r._adzoneRefQueryVO = _adzoneRefQueryVO
	r.Set("adzone_ref_query_v_o", _adzoneRefQueryVO)
	return nil
}

// GetAdzoneRefQueryVO AdzoneRefQueryVO Getter
func (r TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) GetAdzoneRefQueryVO() *AdzoneRefQueryVo {
	return r._adzoneRefQueryVO
}

var poolTaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpAdzoneHorizontalFindpageRequest()
	},
}

// GetTaobaoUniversalbpAdzoneHorizontalFindpageRequest 从 sync.Pool 获取 TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest
func GetTaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest() *TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest {
	return poolTaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest.Get().(*TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest)
}

// ReleaseTaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest 将 TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest(v *TaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpAdzoneHorizontalFindpageAPIRequest.Put(v)
}
