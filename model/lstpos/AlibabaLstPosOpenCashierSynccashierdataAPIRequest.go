package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenCashierSynccashierdataAPIRequest 收银快照同步接口(最多10条订单信息) API请求
// alibaba.lst.pos.open.cashier.synccashierdata
//
// 收银快照同步接口(最多10条订单信息)
type AlibabaLstPosOpenCashierSynccashierdataAPIRequest struct {
	model.Params
	// 订单对象列表
	_cashierFlowDTOList []CashierFlowDto
	// 门店对应的主账号id
	_userId int64
}

// NewAlibabaLstPosOpenCashierSynccashierdataRequest 初始化AlibabaLstPosOpenCashierSynccashierdataAPIRequest对象
func NewAlibabaLstPosOpenCashierSynccashierdataRequest() *AlibabaLstPosOpenCashierSynccashierdataAPIRequest {
	return &AlibabaLstPosOpenCashierSynccashierdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenCashierSynccashierdataAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.cashier.synccashierdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenCashierSynccashierdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCashierFlowDTOList is CashierFlowDTOList Setter
// 订单对象列表
func (r *AlibabaLstPosOpenCashierSynccashierdataAPIRequest) SetCashierFlowDTOList(_cashierFlowDTOList []CashierFlowDto) error {
	r._cashierFlowDTOList = _cashierFlowDTOList
	r.Set("cashier_flow_d_t_o_list", _cashierFlowDTOList)
	return nil
}

// GetCashierFlowDTOList CashierFlowDTOList Getter
func (r AlibabaLstPosOpenCashierSynccashierdataAPIRequest) GetCashierFlowDTOList() []CashierFlowDto {
	return r._cashierFlowDTOList
}

// SetUserId is UserId Setter
// 门店对应的主账号id
func (r *AlibabaLstPosOpenCashierSynccashierdataAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLstPosOpenCashierSynccashierdataAPIRequest) GetUserId() int64 {
	return r._userId
}
