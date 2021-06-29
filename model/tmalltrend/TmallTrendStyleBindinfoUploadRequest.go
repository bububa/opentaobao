package tmalltrend

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
趋势词&款式绑定信息同步API API请求
tmall.trend.style.bindinfo.upload

趋势词&款式(服饰行业)绑定信息同步至平台
*/
type TmallTrendStyleBindinfoUploadRequest struct {
    model.Params
    // 趋势词&款式绑定信息列表，一次最多1000条
    _trendStyleBindInfoBoList   []TrendStyleBindInfoBO
}

// 初始化TmallTrendStyleBindinfoUploadRequest对象
func NewTmallTrendStyleBindinfoUploadRequest() *TmallTrendStyleBindinfoUploadRequest{
    return &TmallTrendStyleBindinfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTrendStyleBindinfoUploadRequest) GetApiMethodName() string {
    return "tmall.trend.style.bindinfo.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTrendStyleBindinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrendStyleBindInfoBoList Setter
// 趋势词&款式绑定信息列表，一次最多1000条
func (r *TmallTrendStyleBindinfoUploadRequest) SetTrendStyleBindInfoBoList(_trendStyleBindInfoBoList []TrendStyleBindInfoBO) error {
    r._trendStyleBindInfoBoList = _trendStyleBindInfoBoList
    r.Set("trend_style_bind_info_bo_list", _trendStyleBindInfoBoList)
    return nil
}

// TrendStyleBindInfoBoList Getter
func (r TmallTrendStyleBindinfoUploadRequest) GetTrendStyleBindInfoBoList() []TrendStyleBindInfoBO {
    return r._trendStyleBindInfoBoList
}
