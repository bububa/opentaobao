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
type AlibabaLstVasTradeflowSaveAPIRequest struct {
    model.Params
    // 交易流水信息
    _tradeFlowModelList   *TradeFlowModel
}

// 初始化AlibabaLstVasTradeflowSaveAPIRequest对象
func NewAlibabaLstVasTradeflowSaveRequest() *AlibabaLstVasTradeflowSaveAPIRequest{
    return &AlibabaLstVasTradeflowSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.vas.tradeflow.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeFlowModelList Setter
// 交易流水信息
func (r *AlibabaLstVasTradeflowSaveAPIRequest) SetTradeFlowModelList(_tradeFlowModelList *TradeFlowModel) error {
    r._tradeFlowModelList = _tradeFlowModelList
    r.Set("trade_flow_model_list", _tradeFlowModelList)
    return nil
}

// TradeFlowModelList Getter
func (r AlibabaLstVasTradeflowSaveAPIRequest) GetTradeFlowModelList() *TradeFlowModel {
    return r._tradeFlowModelList
}
