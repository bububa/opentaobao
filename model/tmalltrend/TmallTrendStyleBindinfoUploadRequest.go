package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
趋势词&款式绑定信息同步API APIRequest
tmall.trend.style.bindinfo.upload

趋势词&款式(服饰行业)绑定信息同步至平台
*/
type TmallTrendStyleBindinfoUploadRequest struct {
    model.Params

    // 趋势词&款式绑定信息列表，一次最多1000条
    trendStyleBindInfoBoList   []TrendStyleBindInfoBO 

}

func NewTmallTrendStyleBindinfoUploadRequest() *TmallTrendStyleBindinfoUploadRequest{
    return &TmallTrendStyleBindinfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTrendStyleBindinfoUploadRequest) GetApiMethodName() string {
    return "tmall.trend.style.bindinfo.upload"
}

func (r TmallTrendStyleBindinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTrendStyleBindinfoUploadRequest) SetTrendStyleBindInfoBoList(trendStyleBindInfoBoList []TrendStyleBindInfoBO) error {
    r.trendStyleBindInfoBoList = trendStyleBindInfoBoList
    r.Set("trend_style_bind_info_bo_list", trendStyleBindInfoBoList)
    return nil
}

func (r TmallTrendStyleBindinfoUploadRequest) GetTrendStyleBindInfoBoList() []TrendStyleBindInfoBO {
    return r.trendStyleBindInfoBoList
}

