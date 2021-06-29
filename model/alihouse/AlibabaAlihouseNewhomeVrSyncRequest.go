package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
VR关系数据同步 APIRequest
alibaba.alihouse.newhome.vr.sync

对接易居VR关系数据迁移
*/
type AlibabaAlihouseNewhomeVrSyncRequest struct {
    model.Params

    // VR提取码
    extractedCode   string 

    // 户型ID
    layoutInfoId   string 

    // 是否生效
    isValid   string 

    // 封面图
    coverImage   string 

    // vr展示链接
    vrUrl   string 

    // 营销户型图
    markingLayoutImg   string 

    // 是否是单层 1 是 0 否
    isSingleLayout   int64 

}

func NewAlibabaAlihouseNewhomeVrSyncRequest() *AlibabaAlihouseNewhomeVrSyncRequest{
    return &AlibabaAlihouseNewhomeVrSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.vr.sync"
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetExtractedCode(extractedCode string) error {
    r.extractedCode = extractedCode
    r.Set("extracted_code", extractedCode)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetExtractedCode() string {
    return r.extractedCode
}

func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetLayoutInfoId(layoutInfoId string) error {
    r.layoutInfoId = layoutInfoId
    r.Set("layout_info_id", layoutInfoId)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetLayoutInfoId() string {
    return r.layoutInfoId
}

func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetIsValid(isValid string) error {
    r.isValid = isValid
    r.Set("is_valid", isValid)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetIsValid() string {
    return r.isValid
}

func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetCoverImage(coverImage string) error {
    r.coverImage = coverImage
    r.Set("cover_image", coverImage)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetCoverImage() string {
    return r.coverImage
}

func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetVrUrl(vrUrl string) error {
    r.vrUrl = vrUrl
    r.Set("vr_url", vrUrl)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetVrUrl() string {
    return r.vrUrl
}

func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetMarkingLayoutImg(markingLayoutImg string) error {
    r.markingLayoutImg = markingLayoutImg
    r.Set("marking_layout_img", markingLayoutImg)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetMarkingLayoutImg() string {
    return r.markingLayoutImg
}

func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetIsSingleLayout(isSingleLayout int64) error {
    r.isSingleLayout = isSingleLayout
    r.Set("is_single_layout", isSingleLayout)
    return nil
}

func (r AlibabaAlihouseNewhomeVrSyncRequest) GetIsSingleLayout() int64 {
    return r.isSingleLayout
}

