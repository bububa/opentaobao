package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
3D款式基本信息同步API APIRequest
tmall.trend.style.basicinfo.upload

3D款式基本信息同步至天猫趋势中心
*/
type TmallTrendStyleBasicinfoUploadRequest struct {
    model.Params

    // 款式基本信息列表，单次同步最多1000条
    styleBasicInfoBoList   []StyleBasicInfoBo 

}

func NewTmallTrendStyleBasicinfoUploadRequest() *TmallTrendStyleBasicinfoUploadRequest{
    return &TmallTrendStyleBasicinfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTrendStyleBasicinfoUploadRequest) GetApiMethodName() string {
    return "tmall.trend.style.basicinfo.upload"
}

func (r TmallTrendStyleBasicinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTrendStyleBasicinfoUploadRequest) SetStyleBasicInfoBoList(styleBasicInfoBoList []StyleBasicInfoBo) error {
    r.styleBasicInfoBoList = styleBasicInfoBoList
    r.Set("style_basic_info_bo_list", styleBasicInfoBoList)
    return nil
}

func (r TmallTrendStyleBasicinfoUploadRequest) GetStyleBasicInfoBoList() []StyleBasicInfoBo {
    return r.styleBasicInfoBoList
}

