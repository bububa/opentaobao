package lstpos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstpos"
)

/* 
分页查询用户全量的门店域商品接口(每页最多20条) 
alibaba.lst.pos.open.goods.getgoodsbypaging

分页查询用户全量的门店域商品接口(每页最多20条)
*/
func AlibabaLstPosOpenGoodsGetgoodsbypaging(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenGoodsGetgoodsbypagingAPIRequest, session string) (*lstpos.AlibabaLstPosOpenGoodsGetgoodsbypagingAPIResponse, error) {
    var resp lstpos.AlibabaLstPosOpenGoodsGetgoodsbypagingAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
