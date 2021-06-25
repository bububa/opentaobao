package promotion

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/promotion"
)

/* 
查询工具 
taobao.ump.tool.get

根据工具id获取一个工具对象
*/
func TaobaoUmpToolGet(clt *core.SDKClient, req *promotion.TaobaoUmpToolGetRequest, session string) (*promotion.TaobaoUmpToolGetResponse, error) {
    var resp promotion.TaobaoUmpToolGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
