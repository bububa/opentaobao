package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiMvMusiclistGetAPIRequest 乐馆mv列表 API请求
// alibaba.xiami.api.mv.musiclist.get
//
// 乐馆mv列表
type AlibabaXiamiApiMvMusiclistGetAPIRequest struct {
	model.Params
	// 语种, 有all, chinese, musician, english, japanese, korea
	_type string
	// 分组, 有recommend、hot、new
	_order string
	// 每页记录
	_limit int64
	// 当前页
	_page int64
}

// NewAlibabaXiamiApiMvMusiclistGetRequest 初始化AlibabaXiamiApiMvMusiclistGetAPIRequest对象
func NewAlibabaXiamiApiMvMusiclistGetRequest() *AlibabaXiamiApiMvMusiclistGetAPIRequest {
	return &AlibabaXiamiApiMvMusiclistGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiMvMusiclistGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.mv.musiclist.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiMvMusiclistGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Type Setter
// 语种, 有all, chinese, musician, english, japanese, korea
func (r *AlibabaXiamiApiMvMusiclistGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaXiamiApiMvMusiclistGetAPIRequest) GetType() string {
	return r._type
}

// Set is Order Setter
// 分组, 有recommend、hot、new
func (r *AlibabaXiamiApiMvMusiclistGetAPIRequest) SetOrder(_order string) error {
	r._order = _order
	r.Set("order", _order)
	return nil
}

// Get Order Getter
func (r AlibabaXiamiApiMvMusiclistGetAPIRequest) GetOrder() string {
	return r._order
}

// Set is Limit Setter
// 每页记录
func (r *AlibabaXiamiApiMvMusiclistGetAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// Get Limit Getter
func (r AlibabaXiamiApiMvMusiclistGetAPIRequest) GetLimit() int64 {
	return r._limit
}

// Set is Page Setter
// 当前页
func (r *AlibabaXiamiApiMvMusiclistGetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaXiamiApiMvMusiclistGetAPIRequest) GetPage() int64 {
	return r._page
}
