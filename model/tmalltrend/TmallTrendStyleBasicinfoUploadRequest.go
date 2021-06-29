package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
3D款式基本信息同步API API请求
tmall.trend.style.basicinfo.upload

3D款式基本信息同步至天猫趋势中心
*/
type TmallTrendStyleBasicinfoUploadRequest struct {
    model.Params
    // 款式基本信息列表，单次同步最多1000条
    _styleBasicInfoBoList   []StyleBasicInfoBo
}

// 初始化TmallTrendStyleBasicinfoUploadRequest对象
func NewTmallTrendStyleBasicinfoUploadRequest() *TmallTrendStyleBasicinfoUploadRequest{
    return &TmallTrendStyleBasicinfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTrendStyleBasicinfoUploadRequest) GetApiMethodName() string {
    return "tmall.trend.style.basicinfo.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTrendStyleBasicinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StyleBasicInfoBoList Setter
// 款式基本信息列表，单次同步最多1000条
func (r *TmallTrendStyleBasicinfoUploadRequest) SetStyleBasicInfoBoList(_styleBasicInfoBoList []StyleBasicInfoBo) error {
    r._styleBasicInfoBoList = _styleBasicInfoBoList
    r.Set("style_basic_info_bo_list", _styleBasicInfoBoList)
    return nil
}

// StyleBasicInfoBoList Getter
func (r TmallTrendStyleBasicinfoUploadRequest) GetStyleBasicInfoBoList() []StyleBasicInfoBo {
    return r._styleBasicInfoBoList
}
