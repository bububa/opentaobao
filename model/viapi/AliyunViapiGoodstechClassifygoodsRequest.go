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
type AliyunViapiGoodstechClassifygoodsRequest struct {
    model.Params
    // 待检测图片链接
    imageUrl   string
}

// 初始化AliyunViapiGoodstechClassifygoodsRequest对象
func NewAliyunViapiGoodstechClassifygoodsRequest() *AliyunViapiGoodstechClassifygoodsRequest{
    return &AliyunViapiGoodstechClassifygoodsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechClassifygoodsRequest) GetApiMethodName() string {
    return "aliyun.viapi.goodstech.classifygoods"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechClassifygoodsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechClassifygoodsRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiGoodstechClassifygoodsRequest) GetImageUrl() string {
    return r.imageUrl
}
