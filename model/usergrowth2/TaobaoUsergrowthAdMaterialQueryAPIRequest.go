package usergrowth2

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsergrowthAdMaterialQueryAPIRequest) Reset() {
	r._outerCreativeId = ""
	r._channelId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.material.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsergrowthAdMaterialQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoUsergrowthAdMaterialQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsergrowthAdMaterialQueryRequest()
	},
}

// GetTaobaoUsergrowthAdMaterialQueryRequest 从 sync.Pool 获取 TaobaoUsergrowthAdMaterialQueryAPIRequest
func GetTaobaoUsergrowthAdMaterialQueryAPIRequest() *TaobaoUsergrowthAdMaterialQueryAPIRequest {
	return poolTaobaoUsergrowthAdMaterialQueryAPIRequest.Get().(*TaobaoUsergrowthAdMaterialQueryAPIRequest)
}

// ReleaseTaobaoUsergrowthAdMaterialQueryAPIRequest 将 TaobaoUsergrowthAdMaterialQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsergrowthAdMaterialQueryAPIRequest(v *TaobaoUsergrowthAdMaterialQueryAPIRequest) {
	v.Reset()
	poolTaobaoUsergrowthAdMaterialQueryAPIRequest.Put(v)
}
