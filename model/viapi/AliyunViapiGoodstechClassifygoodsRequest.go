package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品分类 API请求
aliyun.viapi.goodstech.classifygoods

可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiGoodstechClassifygoodsAPIRequest struct {
    model.Params
    // 待检测图片链接
    _imageUrl   string
}

// 初始化AliyunViapiGoodstechClassifygoodsAPIRequest对象
func NewAliyunViapiGoodstechClassifygoodsRequest() *AliyunViapiGoodstechClassifygoodsAPIRequest{
    return &AliyunViapiGoodstechClassifygoodsAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetApiMethodName() string {
    return "aliyun.viapi.goodstech.classifygoods"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechClassifygoodsAPIRequest) SetImageUrl(_imageUrl string) error {
    r._imageUrl = _imageUrl
    r.Set("image_url", _imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiGoodstechClassifygoodsAPIRequest) GetImageUrl() string {
    return r._imageUrl
}
