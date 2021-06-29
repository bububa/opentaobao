package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
券模板查询 
tmall.nrt.coupontemplate.query

新零售场景，商家拉取在新零售工作台设置的券数据
*/
func TmallNrtCoupontemplateQuery(clt *core.SDKClient, req *nrt.TmallNrtCoupontemplateQueryRequest, session string) (*nrt.TmallNrtCoupontemplateQueryAPIResponse, error) {
    var resp nrt.TmallNrtCoupontemplateQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
