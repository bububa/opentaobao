package viapi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiGoodstechClassifygoodsAPIRequest 商品分类 API请求
// aliyun.viapi.goodstech.classifygoods
//
// 可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiGoodstechClassifygoodsAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunViapiGoodstechClassifygoodsRequest 初始化AliyunViapiGoodstechClassifygoodsAPIRequest对象
func NewAliyunViapiGoodstechClassifygoodsRequest() *AliyunViapiGoodstechClassifygoodsAPIRequest {
	return &AliyunViapiGoodstechClassifygoodsAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunViapiGoodstechClassifygoodsAPIRequest) Reset() {
	r._imageUrl = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.goodstech.classifygoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechClassifygoodsAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetImageUrl() string {
	return r._imageUrl
}

var poolAliyunViapiGoodstechClassifygoodsAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunViapiGoodstechClassifygoodsRequest()
	},
}

// GetAliyunViapiGoodstechClassifygoodsRequest 从 sync.Pool 获取 AliyunViapiGoodstechClassifygoodsAPIRequest
func GetAliyunViapiGoodstechClassifygoodsAPIRequest() *AliyunViapiGoodstechClassifygoodsAPIRequest {
	return poolAliyunViapiGoodstechClassifygoodsAPIRequest.Get().(*AliyunViapiGoodstechClassifygoodsAPIRequest)
}

// ReleaseAliyunViapiGoodstechClassifygoodsAPIRequest 将 AliyunViapiGoodstechClassifygoodsAPIRequest 放入 sync.Pool
func ReleaseAliyunViapiGoodstechClassifygoodsAPIRequest(v *AliyunViapiGoodstechClassifygoodsAPIRequest) {
	v.Reset()
	poolAliyunViapiGoodstechClassifygoodsAPIRequest.Put(v)
}
