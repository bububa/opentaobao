package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
收银快照同步接口(最多10条订单信息) API请求
alibaba.lst.pos.open.cashier.synccashierdata

收银快照同步接口(最多10条订单信息)
*/
type AlibabaLstPosOpenCashierSynccashierdataRequest struct {
    model.Params
    // 订单对象列表
    _cashierFlowDTOList   []CashierFlowDto
    // 门店对应的主账号id
    _userId   int64
}

// 初始化AlibabaLstPosOpenCashierSynccashierdataRequest对象
func NewAlibabaLstPosOpenCashierSynccashierdataRequest() *AlibabaLstPosOpenCashierSynccashierdataRequest{
    return &AlibabaLstPosOpenCashierSynccashierdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.cashier.synccashierdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CashierFlowDTOList Setter
// 订单对象列表
func (r *AlibabaLstPosOpenCashierSynccashierdataRequest) SetCashierFlowDTOList(_cashierFlowDTOList []CashierFlowDto) error {
    r._cashierFlowDTOList = _cashierFlowDTOList
    r.Set("cashier_flow_d_t_o_list", _cashierFlowDTOList)
    return nil
}

// CashierFlowDTOList Getter
func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetCashierFlowDTOList() []CashierFlowDto {
    return r._cashierFlowDTOList
}
// UserId Setter
// 门店对应的主账号id
func (r *AlibabaLstPosOpenCashierSynccashierdataRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenCashierSynccashierdataRequest) GetUserId() int64 {
    return r._userId
}
