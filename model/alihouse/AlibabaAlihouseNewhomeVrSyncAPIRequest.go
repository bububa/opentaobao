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
type AlibabaAlihouseNewhomeVrSyncAPIRequest struct {
    model.Params
    // VR提取码
    _extractedCode   string
    // 户型ID
    _layoutInfoId   string
    // 是否生效
    _isValid   string
    // 封面图
    _coverImage   string
    // vr展示链接
    _vrUrl   string
    // 营销户型图
    _markingLayoutImg   string
    // 是否是单层 1 是 0 否
    _isSingleLayout   int64
}

// 初始化AlibabaAlihouseNewhomeVrSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeVrSyncRequest() *AlibabaAlihouseNewhomeVrSyncAPIRequest{
    return &AlibabaAlihouseNewhomeVrSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.vr.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtractedCode Setter
// VR提取码
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetExtractedCode(_extractedCode string) error {
    r._extractedCode = _extractedCode
    r.Set("extracted_code", _extractedCode)
    return nil
}

// ExtractedCode Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetExtractedCode() string {
    return r._extractedCode
}
// LayoutInfoId Setter
// 户型ID
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetLayoutInfoId(_layoutInfoId string) error {
    r._layoutInfoId = _layoutInfoId
    r.Set("layout_info_id", _layoutInfoId)
    return nil
}

// LayoutInfoId Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetLayoutInfoId() string {
    return r._layoutInfoId
}
// IsValid Setter
// 是否生效
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetIsValid(_isValid string) error {
    r._isValid = _isValid
    r.Set("is_valid", _isValid)
    return nil
}

// IsValid Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetIsValid() string {
    return r._isValid
}
// CoverImage Setter
// 封面图
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetCoverImage(_coverImage string) error {
    r._coverImage = _coverImage
    r.Set("cover_image", _coverImage)
    return nil
}

// CoverImage Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetCoverImage() string {
    return r._coverImage
}
// VrUrl Setter
// vr展示链接
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetVrUrl(_vrUrl string) error {
    r._vrUrl = _vrUrl
    r.Set("vr_url", _vrUrl)
    return nil
}

// VrUrl Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetVrUrl() string {
    return r._vrUrl
}
// MarkingLayoutImg Setter
// 营销户型图
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetMarkingLayoutImg(_markingLayoutImg string) error {
    r._markingLayoutImg = _markingLayoutImg
    r.Set("marking_layout_img", _markingLayoutImg)
    return nil
}

// MarkingLayoutImg Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetMarkingLayoutImg() string {
    return r._markingLayoutImg
}
// IsSingleLayout Setter
// 是否是单层 1 是 0 否
func (r *AlibabaAlihouseNewhomeVrSyncAPIRequest) SetIsSingleLayout(_isSingleLayout int64) error {
    r._isSingleLayout = _isSingleLayout
    r.Set("is_single_layout", _isSingleLayout)
    return nil
}

// IsSingleLayout Getter
func (r AlibabaAlihouseNewhomeVrSyncAPIRequest) GetIsSingleLayout() int64 {
    return r._isSingleLayout
}
