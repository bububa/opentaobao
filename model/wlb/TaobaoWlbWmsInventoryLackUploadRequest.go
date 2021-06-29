package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
缺货通知 API请求
taobao.wlb.wms.inventory.lack.upload

缺货通知
*/
type TaobaoWlbWmsInventoryLackUploadRequest struct {
    model.Params
    // 缺货通知信息
    _content   *WlbWmsInventoryLackUpload
}

// 初始化TaobaoWlbWmsInventoryLackUploadRequest对象
func NewTaobaoWlbWmsInventoryLackUploadRequest() *TaobaoWlbWmsInventoryLackUploadRequest{
    return &TaobaoWlbWmsInventoryLackUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryLackUploadRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.lack.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryLackUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 缺货通知信息
func (r *TaobaoWlbWmsInventoryLackUploadRequest) SetContent(_content *WlbWmsInventoryLackUpload) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoWlbWmsInventoryLackUploadRequest) GetContent() *WlbWmsInventoryLackUpload {
    return r._content
}
