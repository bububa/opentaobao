package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易信息回流 APIRequest
alibaba.lst.vas.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
*/
type AlibabaLstVasTradeflowSaveRequest struct {
    model.Params

    // 交易流水信息
    tradeFlowModelList   *TradeFlowModel 

}

func NewAlibabaLstVasTradeflowSaveRequest() *AlibabaLstVasTradeflowSaveRequest{
    return &AlibabaLstVasTradeflowSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstVasTradeflowSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.vas.tradeflow.save"
}

func (r AlibabaLstVasTradeflowSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstVasTradeflowSaveRequest) SetTradeFlowModelList(tradeFlowModelList *TradeFlowModel) error {
    r.tradeFlowModelList = tradeFlowModelList
    r.Set("trade_flow_model_list", tradeFlowModelList)
    return nil
}

func (r AlibabaLstVasTradeflowSaveRequest) GetTradeFlowModelList() *TradeFlowModel {
    return r.tradeFlowModelList
}

