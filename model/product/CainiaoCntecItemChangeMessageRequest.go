package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品变更消息 API请求
cainiao.cntec.item.change.message

供货商商品信息变更消息
*/
type CainiaoCntecItemChangeMessageRequest struct {
    model.Params
    // 供应商商品变更信息
    _itemChangeMessage   *SupplyItemChangeMessage
}

// 初始化CainiaoCntecItemChangeMessageRequest对象
func NewCainiaoCntecItemChangeMessageRequest() *CainiaoCntecItemChangeMessageRequest{
    return &CainiaoCntecItemChangeMessageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCntecItemChangeMessageRequest) GetApiMethodName() string {
    return "cainiao.cntec.item.change.message"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCntecItemChangeMessageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemChangeMessage Setter
// 供应商商品变更信息
func (r *CainiaoCntecItemChangeMessageRequest) SetItemChangeMessage(_itemChangeMessage *SupplyItemChangeMessage) error {
    r._itemChangeMessage = _itemChangeMessage
    r.Set("item_change_message", _itemChangeMessage)
    return nil
}

// ItemChangeMessage Getter
func (r CainiaoCntecItemChangeMessageRequest) GetItemChangeMessage() *SupplyItemChangeMessage {
    return r._itemChangeMessage
}
