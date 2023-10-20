package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvpvlistAPIRequest 闲鱼已验货pv查询 API请求
// alibaba.idle.isv.pv.list
//
// 根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
type AlibabaidleisvpvlistAPIRequest struct {
	model.Params
	// 系统自动生成
	_brandModelInfo []IdleNewPubValueDo
	// 闲鱼渠道类目的id
	_channelCatId string
}

// NewAlibabaidleisvpvlistRequest 初始化AlibabaidleisvpvlistAPIRequest对象
func NewAlibabaidleisvpvlistRequest() *AlibabaidleisvpvlistAPIRequest {
	return &AlibabaidleisvpvlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleisvpvlistAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.pv.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleisvpvlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleisvpvlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandModelInfo is BrandModelInfo Setter
// 系统自动生成
func (r *AlibabaidleisvpvlistAPIRequest) SetBrandModelInfo(_brandModelInfo []IdleNewPubValueDo) error {
	r._brandModelInfo = _brandModelInfo
	r.Set("brand_model_info", _brandModelInfo)
	return nil
}

// GetBrandModelInfo BrandModelInfo Getter
func (r AlibabaidleisvpvlistAPIRequest) GetBrandModelInfo() []IdleNewPubValueDo {
	return r._brandModelInfo
}

// SetChannelCatId is ChannelCatId Setter
// 闲鱼渠道类目的id
func (r *AlibabaidleisvpvlistAPIRequest) SetChannelCatId(_channelCatId string) error {
	r._channelCatId = _channelCatId
	r.Set("channel_cat_id", _channelCatId)
	return nil
}

// GetChannelCatId ChannelCatId Getter
func (r AlibabaidleisvpvlistAPIRequest) GetChannelCatId() string {
	return r._channelCatId
}
