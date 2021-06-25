package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
缺货通知 APIRequest
taobao.wlb.wms.inventory.lack.upload

缺货通知
*/
type TaobaoWlbWmsInventoryLackUploadRequest struct {
    model.Params

    // 缺货通知信息
    content   *WlbWmsInventoryLackUpload 

}

func NewTaobaoWlbWmsInventoryLackUploadRequest() *TaobaoWlbWmsInventoryLackUploadRequest{
    return &TaobaoWlbWmsInventoryLackUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWmsInventoryLackUploadRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.lack.upload"
}

func (r TaobaoWlbWmsInventoryLackUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWmsInventoryLackUploadRequest) SetContent(content *WlbWmsInventoryLackUpload) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoWlbWmsInventoryLackUploadRequest) GetContent() *WlbWmsInventoryLackUpload {
    return r.content
}

