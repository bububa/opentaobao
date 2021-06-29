package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
根据标签查询实体 
taobao.xhotel.get.entity.by.tag

根据标签查询实体
*/
func TaobaoXhotelGetEntityByTag(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelGetEntityByTagRequest, session string) (*xhotelitem.TaobaoXhotelGetEntityByTagAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelGetEntityByTagAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
