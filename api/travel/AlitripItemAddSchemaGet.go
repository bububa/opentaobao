package travel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/travel"
)

/* 
获取商品发布模板 
alitrip.item.add.schema.get

发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
func AlitripItemAddSchemaGet(clt *core.SDKClient, req *travel.AlitripItemAddSchemaGetAPIRequest, session string) (*travel.AlitripItemAddSchemaGetAPIResponse, error) {
    var resp travel.AlitripItemAddSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
