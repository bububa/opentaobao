package mtopopen

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractMediaArtworkAPIRequest) Reset() {
	r._id = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractMediaArtworkAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.media.artwork"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractMediaArtworkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractMediaArtworkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 系统自动生成
func (r *AlibabaInteractMediaArtworkAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaInteractMediaArtworkAPIRequest) GetId() string {
	return r._id
}

var poolAlibabaInteractMediaArtworkAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractMediaArtworkRequest()
	},
}

// GetAlibabaInteractMediaArtworkRequest 从 sync.Pool 获取 AlibabaInteractMediaArtworkAPIRequest
func GetAlibabaInteractMediaArtworkAPIRequest() *AlibabaInteractMediaArtworkAPIRequest {
	return poolAlibabaInteractMediaArtworkAPIRequest.Get().(*AlibabaInteractMediaArtworkAPIRequest)
}

// ReleaseAlibabaInteractMediaArtworkAPIRequest 将 AlibabaInteractMediaArtworkAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractMediaArtworkAPIRequest(v *AlibabaInteractMediaArtworkAPIRequest) {
	v.Reset()
	poolAlibabaInteractMediaArtworkAPIRequest.Put(v)
}
