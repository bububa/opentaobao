package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinPoiSearchAPIRequest 景点poi搜索-阿信 API请求
// taobao.alitrip.travel.axin.poi.search
//
// 给阿信提供景点poi搜索
type TaobaoAlitripTravelAxinPoiSearchAPIRequest struct {
	model.Params
	// 搜索关键词
	_keyWord string
}

// NewTaobaoAlitripTravelAxinPoiSearchRequest 初始化TaobaoAlitripTravelAxinPoiSearchAPIRequest对象
func NewTaobaoAlitripTravelAxinPoiSearchRequest() *TaobaoAlitripTravelAxinPoiSearchAPIRequest {
	return &TaobaoAlitripTravelAxinPoiSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinPoiSearchAPIRequest) Reset() {
	r._keyWord = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.poi.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyWord is KeyWord Setter
// 搜索关键词
func (r *TaobaoAlitripTravelAxinPoiSearchAPIRequest) SetKeyWord(_keyWord string) error {
	r._keyWord = _keyWord
	r.Set("key_word", _keyWord)
	return nil
}

// GetKeyWord KeyWord Getter
func (r TaobaoAlitripTravelAxinPoiSearchAPIRequest) GetKeyWord() string {
	return r._keyWord
}

var poolTaobaoAlitripTravelAxinPoiSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinPoiSearchRequest()
	},
}

// GetTaobaoAlitripTravelAxinPoiSearchRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinPoiSearchAPIRequest
func GetTaobaoAlitripTravelAxinPoiSearchAPIRequest() *TaobaoAlitripTravelAxinPoiSearchAPIRequest {
	return poolTaobaoAlitripTravelAxinPoiSearchAPIRequest.Get().(*TaobaoAlitripTravelAxinPoiSearchAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinPoiSearchAPIRequest 将 TaobaoAlitripTravelAxinPoiSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinPoiSearchAPIRequest(v *TaobaoAlitripTravelAxinPoiSearchAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinPoiSearchAPIRequest.Put(v)
}
