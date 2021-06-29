package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交图片翻译任务 APIRequest
alibaba.seaking.imagetranslate.submit

提交图片翻译任务
*/
type AlibabaSeakingImagetranslateSubmitRequest struct {
    model.Params

    // token来源站点
    tokenFrom   string 

    // 子任务列表
    imageTranslateDetailList   []ImageTranslateDetailDto 

    // 用户token
    token   string 

}

func NewAlibabaSeakingImagetranslateSubmitRequest() *AlibabaSeakingImagetranslateSubmitRequest{
    return &AlibabaSeakingImagetranslateSubmitRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingImagetranslateSubmitRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate.submit"
}

func (r AlibabaSeakingImagetranslateSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingImagetranslateSubmitRequest) SetTokenFrom(tokenFrom string) error {
    r.tokenFrom = tokenFrom
    r.Set("token_from", tokenFrom)
    return nil
}

func (r AlibabaSeakingImagetranslateSubmitRequest) GetTokenFrom() string {
    return r.tokenFrom
}

func (r *AlibabaSeakingImagetranslateSubmitRequest) SetImageTranslateDetailList(imageTranslateDetailList []ImageTranslateDetailDto) error {
    r.imageTranslateDetailList = imageTranslateDetailList
    r.Set("image_translate_detail_list", imageTranslateDetailList)
    return nil
}

func (r AlibabaSeakingImagetranslateSubmitRequest) GetImageTranslateDetailList() []ImageTranslateDetailDto {
    return r.imageTranslateDetailList
}

func (r *AlibabaSeakingImagetranslateSubmitRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaSeakingImagetranslateSubmitRequest) GetToken() string {
    return r.token
}

