package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商出票中 API请求
taobao.train.agent.handleticket.confirm

代理商出票中
*/
type TaobaoTrainAgentHandleticketConfirmRequest struct {
    model.Params
    // 扩展参数
    extendParams   string
    // 主站id
    mainOrderId   int64
    // 代理商id
    sellerId   int64
}

// 初始化TaobaoTrainAgentHandleticketConfirmRequest对象
func NewTaobaoTrainAgentHandleticketConfirmRequest() *TaobaoTrainAgentHandleticketConfirmRequest{
    return &TaobaoTrainAgentHandleticketConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHandleticketConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.handleticket.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHandleticketConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtendParams Setter
// 扩展参数
func (r *TaobaoTrainAgentHandleticketConfirmRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

// ExtendParams Getter
func (r TaobaoTrainAgentHandleticketConfirmRequest) GetExtendParams() string {
    return r.extendParams
}
// MainOrderId Setter
// 主站id
func (r *TaobaoTrainAgentHandleticketConfirmRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentHandleticketConfirmRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}
// SellerId Setter
// 代理商id
func (r *TaobaoTrainAgentHandleticketConfirmRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r TaobaoTrainAgentHandleticketConfirmRequest) GetSellerId() int64 {
    return r.sellerId
}
