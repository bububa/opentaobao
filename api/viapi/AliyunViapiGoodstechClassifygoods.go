package viapi

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/viapi"
)

/* 
商品分类 
aliyun.viapi.goodstech.classifygoods

可以识别图像中的商品分类，返回商品类目、置信度等信息。目前已经支持服饰鞋包、3C数码、家居用品等超过1万种类目分类。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
*/
func AliyunViapiGoodstechClassifygoods(clt *core.SDKClient, req *viapi.AliyunViapiGoodstechClassifygoodsAPIRequest, session string) (*viapi.AliyunViapiGoodstechClassifygoodsAPIResponse, error) {
    var resp viapi.AliyunViapiGoodstechClassifygoodsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
