package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractmediaartworkAPIRequest 原图相关鉴权接口 API请求
// alibaba.interact.media.artwork
//
// 拍摄并上传原图相关鉴权接口
type AlibabainteractmediaartworkAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabainteractmediaartworkRequest 初始化AlibabainteractmediaartworkAPIRequest对象
func NewAlibabainteractmediaartworkRequest() *AlibabainteractmediaartworkAPIRequest {
	return &AlibabainteractmediaartworkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractmediaartworkAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.media.artwork"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractmediaartworkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractmediaartworkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabainteractmediaartworkAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabainteractmediaartworkAPIRequest) GetId() string {
	return r._id
}
