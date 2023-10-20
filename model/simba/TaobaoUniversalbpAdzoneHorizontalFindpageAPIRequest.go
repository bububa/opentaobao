package simba

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
