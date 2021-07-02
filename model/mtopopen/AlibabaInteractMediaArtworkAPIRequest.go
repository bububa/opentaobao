package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractMediaArtworkAPIRequest 原图相关鉴权接口 API请求
// alibaba.interact.media.artwork
//
// 拍摄并上传原图相关鉴权接口
type AlibabaInteractMediaArtworkAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractMediaArtworkRequest 初始化AlibabaInteractMediaArtworkAPIRequest对象
func NewAlibabaInteractMediaArtworkRequest() *AlibabaInteractMediaArtworkAPIRequest {
	return &AlibabaInteractMediaArtworkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractMediaArtworkAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.media.artwork"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractMediaArtworkAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 系统自动生成
func (r *AlibabaInteractMediaArtworkAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaInteractMediaArtworkAPIRequest) GetId() string {
	return r._id
}
