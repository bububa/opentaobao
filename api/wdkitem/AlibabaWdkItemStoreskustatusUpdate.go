package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
修改门店商品状态 
alibaba.wdk.item.storeskustatus.update

五道口商品 修改门店商品状态
*/
func AlibabaWdkItemStoreskustatusUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemStoreskustatusUpdateRequest, session string) (*wdkitem.AlibabaWdkItemStoreskustatusUpdateAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemStoreskustatusUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
