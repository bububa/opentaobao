package baichuanctg

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuanctg"
)

/* 
微博输出头条数据 
alibaba.baichuan.ctg.toutiao.content

百川头条内容获取
*/
func AlibabaBaichuanCtgToutiaoContent(clt *core.SDKClient, req *baichuanctg.AlibabaBaichuanCtgToutiaoContentAPIRequest, session string) (*baichuanctg.AlibabaBaichuanCtgToutiaoContentAPIResponse, error) {
    var resp baichuanctg.AlibabaBaichuanCtgToutiaoContentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
