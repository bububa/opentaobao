package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
收银快照同步接口(最多10条订单信息) APIRequest
alibaba.lst.pos.open.cashier.synccashierdata

收银快照同步接口(最多10条订单信息)
*/
type AlibabaLstPosOpenCashierSynccashierdataRequest struct {
    model.Params

    // 订单对象列表
    cashierFlowDTOList   []CashierFlowDto 

    // 门店对应的主账号id
    userId   int64 

}

func NewAlibabaLstPosOpenCashierSynccashierdataRequest() *AlibabaLstPosOpenCashierSynccashierdataRequest{
    return &AlibabaLstPosOpenCashierSynccashierdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.cashier.synccashierdata"
}

func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenCashierSynccashierdataRequest) SetCashierFlowDTOList(cashierFlowDTOList []CashierFlowDto) error {
    r.cashierFlowDTOList = cashierFlowDTOList
    r.Set("cashier_flow_d_t_o_list", cashierFlowDTOList)
    return nil
}

func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetCashierFlowDTOList() []CashierFlowDto {
    return r.cashierFlowDTOList
}

func (r *AlibabaLstPosOpenCashierSynccashierdataRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetUserId() int64 {
    return r.userId
}

