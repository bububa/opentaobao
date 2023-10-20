package viapi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunviapigoodstechclassifygoodsAPIRequest 商品分类 API请求
// aliyun.viapi.goodstech.classifygoods
//
// 可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunviapigoodstechclassifygoodsAPIRequest struct {
	model.Params
	// 待检测图片链接
	_imageUrl string
}

// NewAliyunviapigoodstechclassifygoodsRequest 初始化AliyunviapigoodstechclassifygoodsAPIRequest对象
func NewAliyunviapigoodstechclassifygoodsRequest() *AliyunviapigoodstechclassifygoodsAPIRequest {
	return &AliyunviapigoodstechclassifygoodsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunviapigoodstechclassifygoodsAPIRequest) GetApiMethodName() string {
	return "aliyun.viapi.goodstech.classifygoods"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunviapigoodstechclassifygoodsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunviapigoodstechclassifygoodsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrl is ImageUrl Setter
// 待检测图片链接
func (r *AliyunviapigoodstechclassifygoodsAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliyunviapigoodstechclassifygoodsAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
