package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易信息回流 API请求
alibaba.lst.vas.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
*/
type AlibabaLstVasTradeflowSaveRequest struct {
    model.Params
    // 交易流水信息
    tradeFlowModelList   *TradeFlowModel
}

// 初始化AlibabaLstVasTradeflowSaveRequest对象
func NewAlibabaLstVasTradeflowSaveRequest() *AlibabaLstVasTradeflowSaveRequest{
    return &AlibabaLstVasTradeflowSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVasTradeflowSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.vas.tradeflow.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVasTradeflowSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeFlowModelList Setter
// 交易流水信息
func (r *AlibabaLstVasTradeflowSaveRequest) SetTradeFlowModelList(tradeFlowModelList *TradeFlowModel) error {
    r.tradeFlowModelList = tradeFlowModelList
    r.Set("trade_flow_model_list", tradeFlowModelList)
    return nil
}

// TradeFlowModelList Getter
func (r AlibabaLstVasTradeflowSaveRequest) GetTradeFlowModelList() *TradeFlowModel {
    return r.tradeFlowModelList
}
