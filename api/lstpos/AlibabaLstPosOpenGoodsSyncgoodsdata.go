package lstpos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstpos"
)

/* 
门店商品批量同步接口(最多10条商品信息) 
alibaba.lst.pos.open.goods.syncgoodsdata

门店商品批量同步接口(最多10条商品信息)
*/
func AlibabaLstPosOpenGoodsSyncgoodsdata(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataRequest, session string) (*lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse, error) {
    var resp lstpos.AlibabaLstPosOpenGoodsSyncgoodsdataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
