package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinpoisearchAPIRequest 景点poi搜索-阿信 API请求
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
type TaobaoalitriptravelaxinpoisearchAPIRequest struct {
	model.Params
	// 搜索关键词
	_keyWord string
}

// NewTaobaoalitriptravelaxinpoisearchRequest 初始化TaobaoalitriptravelaxinpoisearchAPIRequest对象
func NewTaobaoalitriptravelaxinpoisearchRequest() *TaobaoalitriptravelaxinpoisearchAPIRequest {
	return &TaobaoalitriptravelaxinpoisearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinpoisearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.poi.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinpoisearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinpoisearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyWord is KeyWord Setter
// 搜索关键词
func (r *TaobaoalitriptravelaxinpoisearchAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// GetKeyWord KeyWord Getter
func (r TaobaoalitriptravelaxinpoisearchAPIRequest) GetKeyWord() string {
	return r._keyWord
}
