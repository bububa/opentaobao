package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpStdcategoryFindlistAPIRequest 人群相关类目查询 API请求
// taobao.universalbp.stdcategory.findlist
//
// 入参商品信息，出参商品所属类别
type TaobaoUniversalbpStdcategoryFindlistAPIRequest struct {
	model.Params
	// topServiceContext
	_topServiceContext *TopServiceContext
	// stdCategoryQueryVO
	_stdCategoryQueryVO *StdCategoryQueryVo
}

// NewTaobaoUniversalbpStdcategoryFindlistRequest 初始化TaobaoUniversalbpStdcategoryFindlistAPIRequest对象
func NewTaobaoUniversalbpStdcategoryFindlistRequest() *TaobaoUniversalbpStdcategoryFindlistAPIRequest {
	return &TaobaoUniversalbpStdcategoryFindlistAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUniversalbpStdcategoryFindlistAPIRequest) Reset() {
	r._topServiceContext = nil
	r._stdCategoryQueryVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUniversalbpStdcategoryFindlistAPIRequest) GetApiMethodName() string {
	return "taobao.universalbp.stdcategory.findlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUniversalbpStdcategoryFindlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUniversalbpStdcategoryFindlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopServiceContext is TopServiceContext Setter
// topServiceContext
func (r *TaobaoUniversalbpStdcategoryFindlistAPIRequest) SetTopServiceContext(_topServiceContext *TopServiceContext) error {
	r._topServiceContext = _topServiceContext
	r.Set("top_service_context", _topServiceContext)
	return nil
}

// GetTopServiceContext TopServiceContext Getter
func (r TaobaoUniversalbpStdcategoryFindlistAPIRequest) GetTopServiceContext() *TopServiceContext {
	return r._topServiceContext
}

// SetStdCategoryQueryVO is StdCategoryQueryVO Setter
// stdCategoryQueryVO
func (r *TaobaoUniversalbpStdcategoryFindlistAPIRequest) SetStdCategoryQueryVO(_stdCategoryQueryVO *StdCategoryQueryVo) error {
	r._stdCategoryQueryVO = _stdCategoryQueryVO
	r.Set("std_category_query_v_o", _stdCategoryQueryVO)
	return nil
}

// GetStdCategoryQueryVO StdCategoryQueryVO Getter
func (r TaobaoUniversalbpStdcategoryFindlistAPIRequest) GetStdCategoryQueryVO() *StdCategoryQueryVo {
	return r._stdCategoryQueryVO
}

var poolTaobaoUniversalbpStdcategoryFindlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUniversalbpStdcategoryFindlistRequest()
	},
}

// GetTaobaoUniversalbpStdcategoryFindlistRequest 从 sync.Pool 获取 TaobaoUniversalbpStdcategoryFindlistAPIRequest
func GetTaobaoUniversalbpStdcategoryFindlistAPIRequest() *TaobaoUniversalbpStdcategoryFindlistAPIRequest {
	return poolTaobaoUniversalbpStdcategoryFindlistAPIRequest.Get().(*TaobaoUniversalbpStdcategoryFindlistAPIRequest)
}

// ReleaseTaobaoUniversalbpStdcategoryFindlistAPIRequest 将 TaobaoUniversalbpStdcategoryFindlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoUniversalbpStdcategoryFindlistAPIRequest(v *TaobaoUniversalbpStdcategoryFindlistAPIRequest) {
	v.Reset()
	poolTaobaoUniversalbpStdcategoryFindlistAPIRequest.Put(v)
}
