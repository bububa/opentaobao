package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
肌肤检测 APIRequest
tmall.marketing.face.skindetect

提供人脸肌肤属性报告
*/
type TmallMarketingFaceSkindetectRequest struct {
    model.Params

    // 图片的base64（必须以base64,开头）
    image   string 

    // isv标识
    source   string 

    // 前置摄像头1，后置摄像头0
    frontCamera   string 

    // 混淆nick
    mixnick   string 

}

func NewTmallMarketingFaceSkindetectRequest() *TmallMarketingFaceSkindetectRequest{
    return &TmallMarketingFaceSkindetectRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMarketingFaceSkindetectRequest) GetApiMethodName() string {
    return "tmall.marketing.face.skindetect"
}

func (r TmallMarketingFaceSkindetectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMarketingFaceSkindetectRequest) SetImage(image string) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TmallMarketingFaceSkindetectRequest) GetImage() string {
    return r.image
}

func (r *TmallMarketingFaceSkindetectRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TmallMarketingFaceSkindetectRequest) GetSource() string {
    return r.source
}

func (r *TmallMarketingFaceSkindetectRequest) SetFrontCamera(frontCamera string) error {
    r.frontCamera = frontCamera
    r.Set("front_camera", frontCamera)
    return nil
}

func (r TmallMarketingFaceSkindetectRequest) GetFrontCamera() string {
    return r.frontCamera
}

func (r *TmallMarketingFaceSkindetectRequest) SetMixnick(mixnick string) error {
    r.mixnick = mixnick
    r.Set("mixnick", mixnick)
    return nil
}

func (r TmallMarketingFaceSkindetectRequest) GetMixnick() string {
    return r.mixnick
}

