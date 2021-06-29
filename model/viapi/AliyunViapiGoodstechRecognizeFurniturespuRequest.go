package viapi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
家居SPU识别 API请求
aliyun.viapi.goodstech.recognize.furniturespu

对输入的家居模型图进行分类，目前类别数可达70类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
type AliyunViapiGoodstechRecognizeFurniturespuRequest struct {
    model.Params
    // 待检测图片链接
    imageUrl   string
    // 模型x方向的尺寸，单位cm，默认值100
    xLength   string
    // 模型y方向的尺寸，单位cm，默认值100
    yLength   string
    // 模型z方向的尺寸，单位cm，默认值100
    zLength   string
}

// 初始化AliyunViapiGoodstechRecognizeFurniturespuRequest对象
func NewAliyunViapiGoodstechRecognizeFurniturespuRequest() *AliyunViapiGoodstechRecognizeFurniturespuRequest{
    return &AliyunViapiGoodstechRecognizeFurniturespuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunViapiGoodstechRecognizeFurniturespuRequest) GetApiMethodName() string {
    return "aliyun.viapi.goodstech.recognize.furniturespu"
}

// IRequest interface 方法, 获取API参数
func (r AliyunViapiGoodstechRecognizeFurniturespuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ImageUrl Setter
// 待检测图片链接
func (r *AliyunViapiGoodstechRecognizeFurniturespuRequest) SetImageUrl(imageUrl string) error {
    r.imageUrl = imageUrl
    r.Set("image_url", imageUrl)
    return nil
}

// ImageUrl Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuRequest) GetImageUrl() string {
    return r.imageUrl
}
// XLength Setter
// 模型x方向的尺寸，单位cm，默认值100
func (r *AliyunViapiGoodstechRecognizeFurniturespuRequest) SetXLength(xLength string) error {
    r.xLength = xLength
    r.Set("x_length", xLength)
    return nil
}

// XLength Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuRequest) GetXLength() string {
    return r.xLength
}
// YLength Setter
// 模型y方向的尺寸，单位cm，默认值100
func (r *AliyunViapiGoodstechRecognizeFurniturespuRequest) SetYLength(yLength string) error {
    r.yLength = yLength
    r.Set("y_length", yLength)
    return nil
}

// YLength Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuRequest) GetYLength() string {
    return r.yLength
}
// ZLength Setter
// 模型z方向的尺寸，单位cm，默认值100
func (r *AliyunViapiGoodstechRecognizeFurniturespuRequest) SetZLength(zLength string) error {
    r.zLength = zLength
    r.Set("z_length", zLength)
    return nil
}

// ZLength Getter
func (r AliyunViapiGoodstechRecognizeFurniturespuRequest) GetZLength() string {
    return r.zLength
}
