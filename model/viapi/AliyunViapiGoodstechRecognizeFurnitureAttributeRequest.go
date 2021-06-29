package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家居属性识别 API请求
aliyun.viapi.goodstech.recognize.furniture.attribute

识别输入的家居模型图的风格，目前支持16种风格识别。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiGoodstechRecognizeFurnitureAttributeRequest struct {
    model.Params
    // 待检测图片链接
    imageUrl   string
}

// 初始化AliyunViapiGoodstechRecognizeFurnitureAttributeRequest对象
func NewAliyunViapiGoodstechRecognizeFurnitureAttributeRequest() *AliyunViapiGoodstechRecognizeFurnitureAttributeRequest{
    return &AliyunViapiGoodstechRecognizeFurnitureAttributeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechRecognizeFurnitureAttributeRequest) GetApiMethodName() string {
    return "aliyun.viapi.goodstech.recognize.furniture.attribute"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechRecognizeFurnitureAttributeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechRecognizeFurnitureAttributeRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiGoodstechRecognizeFurnitureAttributeRequest) GetImageUrl() string {
    return r.imageUrl
}
