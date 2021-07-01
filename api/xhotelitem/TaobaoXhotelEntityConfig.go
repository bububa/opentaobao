package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
飞猪商品各实体通用配置 
taobao.xhotel.entity.config

飞猪商品各实体通用配置服务
*/
func TaobaoXhotelEntityConfig(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelEntityConfigAPIRequest, session string) (*xhotelitem.TaobaoXhotelEntityConfigAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelEntityConfigAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
