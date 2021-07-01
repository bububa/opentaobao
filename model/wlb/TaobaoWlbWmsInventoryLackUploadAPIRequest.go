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
type TaobaoWlbWmsInventoryLackUploadAPIRequest struct {
    model.Params
    // 缺货通知信息
    _content   *WlbWmsInventoryLackUpload
}

// 初始化TaobaoWlbWmsInventoryLackUploadAPIRequest对象
func NewTaobaoWlbWmsInventoryLackUploadRequest() *TaobaoWlbWmsInventoryLackUploadAPIRequest{
    return &TaobaoWlbWmsInventoryLackUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.wms.inventory.lack.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 缺货通知信息
func (r *TaobaoWlbWmsInventoryLackUploadAPIRequest) SetContent(_content *WlbWmsInventoryLackUpload) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoWlbWmsInventoryLackUploadAPIRequest) GetContent() *WlbWmsInventoryLackUpload {
    return r._content
}
