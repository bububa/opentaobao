package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthAdMaterialQueryAPIRequest 素材审核结果查询 API请求
// taobao.usergrowth.ad.material.query
//
// 素材审核结果查询
type TaobaoUsergrowthAdMaterialQueryAPIRequest struct {
	model.Params
	// 渠道-创意id
	_outerCreativeId string
	// 渠道id
	_channelId int64
}

// NewTaobaoUsergrowthAdMaterialQueryRequest 初始化TaobaoUsergrowthAdMaterialQueryAPIRequest对象
func NewTaobaoUsergrowthAdMaterialQueryRequest() *TaobaoUsergrowthAdMaterialQueryAPIRequest {
	return &TaobaoUsergrowthAdMaterialQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.material.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterCreativeId is OuterCreativeId Setter
// 渠道-创意id
func (r *TaobaoUsergrowthAdMaterialQueryAPIRequest) SetOuterCreativeId(_outerCreativeId string) error {
	r._outerCreativeId = _outerCreativeId
	r.Set("outer_creative_id", _outerCreativeId)
	return nil
}

// GetOuterCreativeId OuterCreativeId Getter
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetOuterCreativeId() string {
	return r._outerCreativeId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *TaobaoUsergrowthAdMaterialQueryAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetChannelId() int64 {
	return r._channelId
}
