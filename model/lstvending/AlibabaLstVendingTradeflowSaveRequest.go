package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机交易信息回流 API请求
alibaba.lst.vending.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
*/
type AlibabaLstVendingTradeflowSaveRequest struct {
    model.Params
    // 交易流水信息
    tradeFlowDTOList   []VendingTradeFlowDto
}

// 初始化AlibabaLstVendingTradeflowSaveRequest对象
func NewAlibabaLstVendingTradeflowSaveRequest() *AlibabaLstVendingTradeflowSaveRequest{
    return &AlibabaLstVendingTradeflowSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingTradeflowSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.tradeflow.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingTradeflowSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeFlowDTOList Setter
// 交易流水信息
func (r *AlibabaLstVendingTradeflowSaveRequest) SetTradeFlowDTOList(tradeFlowDTOList []VendingTradeFlowDto) error {
    r.tradeFlowDTOList = tradeFlowDTOList
    r.Set("trade_flow_d_t_o_list", tradeFlowDTOList)
    return nil
}

// TradeFlowDTOList Getter
func (r AlibabaLstVendingTradeflowSaveRequest) GetTradeFlowDTOList() []VendingTradeFlowDto {
    return r.tradeFlowDTOList
}
