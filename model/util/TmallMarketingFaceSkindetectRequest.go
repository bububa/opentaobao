package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
肌肤检测 API请求
tmall.marketing.face.skindetect

提供人脸肌肤属性报告
*/
type TmallMarketingFaceSkindetectAPIRequest struct {
    model.Params
    // 图片的base64（必须以base64,开头）
    _image   string
    // isv标识
    _source   string
    // 前置摄像头1，后置摄像头0
    _frontCamera   string
    // 混淆nick
    _mixnick   string
}

// 初始化TmallMarketingFaceSkindetectAPIRequest对象
func NewTmallMarketingFaceSkindetectRequest() *TmallMarketingFaceSkindetectAPIRequest{
    return &TmallMarketingFaceSkindetectAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMarketingFaceSkindetectAPIRequest) GetApiMethodName() string {
    return "tmall.marketing.face.skindetect"
}

// IRequest interface 方法, 获取API参数
func (r TmallMarketingFaceSkindetectAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Image Setter
// 图片的base64（必须以base64,开头）
func (r *TmallMarketingFaceSkindetectAPIRequest) SetImage(_image string) error {
    r._image = _image
    r.Set("image", _image)
    return nil
}

// Image Getter
func (r TmallMarketingFaceSkindetectAPIRequest) GetImage() string {
    return r._image
}
// Source Setter
// isv标识
func (r *TmallMarketingFaceSkindetectAPIRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TmallMarketingFaceSkindetectAPIRequest) GetSource() string {
    return r._source
}
// FrontCamera Setter
// 前置摄像头1，后置摄像头0
func (r *TmallMarketingFaceSkindetectAPIRequest) SetFrontCamera(_frontCamera string) error {
    r._frontCamera = _frontCamera
    r.Set("front_camera", _frontCamera)
    return nil
}

// FrontCamera Getter
func (r TmallMarketingFaceSkindetectAPIRequest) GetFrontCamera() string {
    return r._frontCamera
}
// Mixnick Setter
// 混淆nick
func (r *TmallMarketingFaceSkindetectAPIRequest) SetMixnick(_mixnick string) error {
    r._mixnick = _mixnick
    r.Set("mixnick", _mixnick)
    return nil
}

// Mixnick Getter
func (r TmallMarketingFaceSkindetectAPIRequest) GetMixnick() string {
    return r._mixnick
}
