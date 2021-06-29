package tmalltrend

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmalltrend"
)

/* 
3D款式基本信息同步API 
tmall.trend.style.basicinfo.upload

3D款式基本信息同步至天猫趋势中心
*/
func TmallTrendStyleBasicinfoUpload(clt *core.SDKClient, req *tmalltrend.TmallTrendStyleBasicinfoUploadRequest, session string) (*tmalltrend.TmallTrendStyleBasicinfoUploadAPIResponse, error) {
    var resp tmalltrend.TmallTrendStyleBasicinfoUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
