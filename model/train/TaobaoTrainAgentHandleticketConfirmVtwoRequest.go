package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商出票中v2--增加鉴权校验 API请求
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中
*/
type TaobaoTrainAgentHandleticketConfirmVtwoRequest struct {
    model.Params
    // 扩展参数
    extendParams   string
    // 主站id
    mainOrderId   int64
    // 代理商id
    sellerId   int64
}

// 初始化TaobaoTrainAgentHandleticketConfirmVtwoRequest对象
func NewTaobaoTrainAgentHandleticketConfirmVtwoRequest() *TaobaoTrainAgentHandleticketConfirmVtwoRequest{
    return &TaobaoTrainAgentHandleticketConfirmVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.handleticket.confirm.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendParams Setter
// 扩展参数
func (r *TaobaoTrainAgentHandleticketConfirmVtwoRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

// ExtendParams Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetExtendParams() string {
    return r.extendParams
}
// MainOrderId Setter
// 主站id
func (r *TaobaoTrainAgentHandleticketConfirmVtwoRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// SellerId Setter
// 代理商id
func (r *TaobaoTrainAgentHandleticketConfirmVtwoRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetSellerId() int64 {
    return r.sellerId
}
