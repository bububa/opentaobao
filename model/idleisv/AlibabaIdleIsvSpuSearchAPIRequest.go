package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvspusearchAPIRequest spu搜索接口 API请求
// alibaba.idle.isv.spu.search
//
// 搜索的品牌和型号，供服务商进行选择
type AlibabaidleisvspusearchAPIRequest struct {
	model.Params
	// 闲鱼渠道类目的id
	_channelCatId string
	// 搜索的文本
	_searchText string
}

// NewAlibabaidleisvspusearchRequest 初始化AlibabaidleisvspusearchAPIRequest对象
func NewAlibabaidleisvspusearchRequest() *AlibabaidleisvspusearchAPIRequest {
	return &AlibabaidleisvspusearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvspusearchAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.spu.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvspusearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvspusearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCatId is ChannelCatId Setter
// 闲鱼渠道类目的id
func (r *AlibabaidleisvspusearchAPIRequest) SetChannelCatId(_channelCatId string) error {
	r._channelCatId = _channelCatId
	r.Set("channel_cat_id", _channelCatId)
	return nil
}

// GetChannelCatId ChannelCatId Getter
func (r AlibabaidleisvspusearchAPIRequest) GetChannelCatId() string {
	return r._channelCatId
}

// SetSearchText is SearchText Setter
// 搜索的文本
func (r *AlibabaidleisvspusearchAPIRequest) SetSearchText(_searchText string) error {
	r._searchText = _searchText
	r.Set("search_text", _searchText)
	return nil
}

// GetSearchText SearchText Getter
func (r AlibabaidleisvspusearchAPIRequest) GetSearchText() string {
	return r._searchText
}
