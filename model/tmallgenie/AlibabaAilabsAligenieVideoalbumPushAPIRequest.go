package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieVideoalbumPushAPIRequest 天猫精灵内容库视频合辑数据推送接口 API请求
// alibaba.ailabs.aligenie.videoalbum.push
//
// 三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用
type AlibabaAilabsAligenieVideoalbumPushAPIRequest struct {
	model.Params
	// 视频合辑数据
	_param1 []RawVideoAlbum
}

// NewAlibabaAilabsAligenieVideoalbumPushRequest 初始化AlibabaAilabsAligenieVideoalbumPushAPIRequest对象
func NewAlibabaAilabsAligenieVideoalbumPushRequest() *AlibabaAilabsAligenieVideoalbumPushAPIRequest {
	return &AlibabaAilabsAligenieVideoalbumPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieVideoalbumPushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.videoalbum.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieVideoalbumPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieVideoalbumPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam1 is Param1 Setter
// 视频合辑数据
func (r *AlibabaAilabsAligenieVideoalbumPushAPIRequest) SetParam1(_param1 []RawVideoAlbum) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaAilabsAligenieVideoalbumPushAPIRequest) GetParam1() []RawVideoAlbum {
	return r._param1
}
