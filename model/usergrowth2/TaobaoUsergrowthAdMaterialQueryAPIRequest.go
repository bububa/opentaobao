package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthadmaterialqueryAPIRequest 素材审核结果查询 API请求
// taobao.usergrowth.ad.material.query
//
// 素材审核结果查询
type TaobaousergrowthadmaterialqueryAPIRequest struct {
	model.Params
	// 渠道-创意id
	_outerCreativeId string
	// 渠道id
	_channelId int64
}

// NewTaobaousergrowthadmaterialqueryRequest 初始化TaobaousergrowthadmaterialqueryAPIRequest对象
func NewTaobaousergrowthadmaterialqueryRequest() *TaobaousergrowthadmaterialqueryAPIRequest {
	return &TaobaousergrowthadmaterialqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousergrowthadmaterialqueryAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.material.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousergrowthadmaterialqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousergrowthadmaterialqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCreativeId is OuterCreativeId Setter
// 渠道-创意id
func (r *TaobaousergrowthadmaterialqueryAPIRequest) SetOuterCreativeId(_outerCreativeId string) error {
	r._outerCreativeId = _outerCreativeId
	r.Set("outer_creative_id", _outerCreativeId)
	return nil
}

// GetOuterCreativeId OuterCreativeId Getter
func (r TaobaousergrowthadmaterialqueryAPIRequest) GetOuterCreativeId() string {
	return r._outerCreativeId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *TaobaousergrowthadmaterialqueryAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaousergrowthadmaterialqueryAPIRequest) GetChannelId() int64 {
	return r._channelId
}
