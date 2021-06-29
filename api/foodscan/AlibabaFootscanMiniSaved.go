package foodscan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/foodscan"
)

/* 
更新报告状态 
alibaba.footscan.mini.saved

更新报告状态接口
*/
func AlibabaFootscanMiniSaved(clt *core.SDKClient, req *foodscan.AlibabaFootscanMiniSavedRequest, session string) (*foodscan.AlibabaFootscanMiniSavedAPIResponse, error) {
    var resp foodscan.AlibabaFootscanMiniSavedAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
