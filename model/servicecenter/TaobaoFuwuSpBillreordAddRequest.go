package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
内购服务确认单明细上传接口 APIRequest
taobao.fuwu.sp.billreord.add

isv能通过该接口上传确认单明细数据
*/
type TaobaoFuwuSpBillreordAddRequest struct {
    model.Params

    // 确认单的账单明细
    paramBillRecordDTO   *BillRecordDto 

}

func NewTaobaoFuwuSpBillreordAddRequest() *TaobaoFuwuSpBillreordAddRequest{
    return &TaobaoFuwuSpBillreordAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFuwuSpBillreordAddRequest) GetApiMethodName() string {
    return "taobao.fuwu.sp.billreord.add"
}

func (r TaobaoFuwuSpBillreordAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFuwuSpBillreordAddRequest) SetParamBillRecordDTO(paramBillRecordDTO *BillRecordDto) error {
    r.paramBillRecordDTO = paramBillRecordDTO
    r.Set("param_bill_record_d_t_o", paramBillRecordDTO)
    return nil
}

func (r TaobaoFuwuSpBillreordAddRequest) GetParamBillRecordDTO() *BillRecordDto {
    return r.paramBillRecordDTO
}

