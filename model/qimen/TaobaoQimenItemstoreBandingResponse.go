package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品关联绑定接口 APIResponse
taobao.qimen.itemstore.banding

商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
*/
type TaobaoQimenItemstoreBandingAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenItemstoreBandingResponse `json:"qimen_itemstore_banding_response,omitempty"` 
    TaobaoQimenItemstoreBandingResponse
}

/* model for simplify = false
type TaobaoQimenItemstoreBandingResponse struct {

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 响应描述
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应编码
    
    QimenCode   string `json:"qimen_code,omitempty"`
    

}
*/

type TaobaoQimenItemstoreBandingResponse struct {

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 响应描述
    Flag   string `json:"flag,omitempty"`

    // 响应编码
    QimenCode   string `json:"qimen_code,omitempty"`

}
