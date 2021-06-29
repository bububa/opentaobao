package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
VR关系数据同步 API请求
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

// 初始化AlibabaAlihouseNewhomeVrSyncRequest对象
func NewAlibabaAlihouseNewhomeVrSyncRequest() *AlibabaAlihouseNewhomeVrSyncRequest{
    return &AlibabaAlihouseNewhomeVrSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.vr.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtractedCode Setter
// VR提取码
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetExtractedCode(extractedCode string) error {
    r.extractedCode = extractedCode
    r.Set("extracted_code", extractedCode)
    return nil
}

// ExtractedCode Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetExtractedCode() string {
    return r.extractedCode
}
// LayoutInfoId Setter
// 户型ID
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetLayoutInfoId(layoutInfoId string) error {
    r.layoutInfoId = layoutInfoId
    r.Set("layout_info_id", layoutInfoId)
    return nil
}

// LayoutInfoId Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetLayoutInfoId() string {
    return r.layoutInfoId
}
// IsValid Setter
// 是否生效
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetIsValid(isValid string) error {
    r.isValid = isValid
    r.Set("is_valid", isValid)
    return nil
}

// IsValid Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetIsValid() string {
    return r.isValid
}
// CoverImage Setter
// 封面图
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetCoverImage(coverImage string) error {
    r.coverImage = coverImage
    r.Set("cover_image", coverImage)
    return nil
}

// CoverImage Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetCoverImage() string {
    return r.coverImage
}
// VrUrl Setter
// vr展示链接
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetVrUrl(vrUrl string) error {
    r.vrUrl = vrUrl
    r.Set("vr_url", vrUrl)
    return nil
}

// VrUrl Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetVrUrl() string {
    return r.vrUrl
}
// MarkingLayoutImg Setter
// 营销户型图
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetMarkingLayoutImg(markingLayoutImg string) error {
    r.markingLayoutImg = markingLayoutImg
    r.Set("marking_layout_img", markingLayoutImg)
    return nil
}

// MarkingLayoutImg Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetMarkingLayoutImg() string {
    return r.markingLayoutImg
}
// IsSingleLayout Setter
// 是否是单层 1 是 0 否
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetIsSingleLayout(isSingleLayout int64) error {
    r.isSingleLayout = isSingleLayout
    r.Set("is_single_layout", isSingleLayout)
    return nil
}

// IsSingleLayout Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetIsSingleLayout() int64 {
    return r.isSingleLayout
}
