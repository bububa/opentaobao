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
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetExtractedCode(_extractedCode string) error {
    r._extractedCode = _extractedCode
    r.Set("extracted_code", _extractedCode)
    return nil
}

// ExtractedCode Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetExtractedCode() string {
    return r._extractedCode
}
// LayoutInfoId Setter
// 户型ID
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetLayoutInfoId(_layoutInfoId string) error {
    r._layoutInfoId = _layoutInfoId
    r.Set("layout_info_id", _layoutInfoId)
    return nil
}

// LayoutInfoId Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetLayoutInfoId() string {
    return r._layoutInfoId
}
// IsValid Setter
// 是否生效
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetIsValid(_isValid string) error {
    r._isValid = _isValid
    r.Set("is_valid", _isValid)
    return nil
}

// IsValid Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetIsValid() string {
    return r._isValid
}
// CoverImage Setter
// 封面图
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetCoverImage(_coverImage string) error {
    r._coverImage = _coverImage
    r.Set("cover_image", _coverImage)
    return nil
}

// CoverImage Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetCoverImage() string {
    return r._coverImage
}
// VrUrl Setter
// vr展示链接
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetVrUrl(_vrUrl string) error {
    r._vrUrl = _vrUrl
    r.Set("vr_url", _vrUrl)
    return nil
}

// VrUrl Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetVrUrl() string {
    return r._vrUrl
}
// MarkingLayoutImg Setter
// 营销户型图
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetMarkingLayoutImg(_markingLayoutImg string) error {
    r._markingLayoutImg = _markingLayoutImg
    r.Set("marking_layout_img", _markingLayoutImg)
    return nil
}

// MarkingLayoutImg Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetMarkingLayoutImg() string {
    return r._markingLayoutImg
}
// IsSingleLayout Setter
// 是否是单层 1 是 0 否
func (r *AlibabaAlihouseNewhomeVrSyncRequest) SetIsSingleLayout(_isSingleLayout int64) error {
    r._isSingleLayout = _isSingleLayout
    r.Set("is_single_layout", _isSingleLayout)
    return nil
}

// IsSingleLayout Getter
func (r AlibabaAlihouseNewhomeVrSyncRequest) GetIsSingleLayout() int64 {
    return r._isSingleLayout
}
