package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交图片翻译任务 API请求
alibaba.seaking.imagetranslate.submit

提交图片翻译任务
*/
type AlibabaSeakingImagetranslateSubmitAPIRequest struct {
    model.Params
    // token来源站点
    _tokenFrom   string
    // 子任务列表
    _imageTranslateDetailList   []ImageTranslateDetailDTO
    // 用户token
    _token   string
}

// 初始化AlibabaSeakingImagetranslateSubmitAPIRequest对象
func NewAlibabaSeakingImagetranslateSubmitRequest() *AlibabaSeakingImagetranslateSubmitAPIRequest{
    return &AlibabaSeakingImagetranslateSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagetranslateSubmitAPIRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagetranslateSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingImagetranslateSubmitAPIRequest) SetTokenFrom(_tokenFrom string) error {
    r._tokenFrom = _tokenFrom
    r.Set("token_from", _tokenFrom)
    return nil
}

// TokenFrom Getter
func (r AlibabaSeakingImagetranslateSubmitAPIRequest) GetTokenFrom() string {
    return r._tokenFrom
}
// ImageTranslateDetailList Setter
// 子任务列表
func (r *AlibabaSeakingImagetranslateSubmitAPIRequest) SetImageTranslateDetailList(_imageTranslateDetailList []ImageTranslateDetailDTO) error {
    r._imageTranslateDetailList = _imageTranslateDetailList
    r.Set("image_translate_detail_list", _imageTranslateDetailList)
    return nil
}

// ImageTranslateDetailList Getter
func (r AlibabaSeakingImagetranslateSubmitAPIRequest) GetImageTranslateDetailList() []ImageTranslateDetailDTO {
    return r._imageTranslateDetailList
}
// Token Setter
// 用户token
func (r *AlibabaSeakingImagetranslateSubmitAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaSeakingImagetranslateSubmitAPIRequest) GetToken() string {
    return r._token
}
