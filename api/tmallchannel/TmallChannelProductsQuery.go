package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
渠道中心-查询产品列表 
tmall.channel.products.query

渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
*/
func TmallChannelProductsQuery(clt *core.SDKClient, req *tmallchannel.TmallChannelProductsQueryRequest, session string) (*tmallchannel.TmallChannelProductsQueryAPIResponse, error) {
    var resp tmallchannel.TmallChannelProductsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
