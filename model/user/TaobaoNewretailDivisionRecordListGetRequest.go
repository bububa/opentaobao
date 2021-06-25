package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导购分佣明细列表 APIRequest
taobao.newretail.division.record.list.get

提供分页查询导购分佣明细的能力
*/
type TaobaoNewretailDivisionRecordListGetRequest struct {
    model.Params

    // 入参
    param   *TopDivisionRecordReqDto 

}

func NewTaobaoNewretailDivisionRecordListGetRequest() *TaobaoNewretailDivisionRecordListGetRequest{
    return &TaobaoNewretailDivisionRecordListGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoNewretailDivisionRecordListGetRequest) GetApiMethodName() string {
    return "taobao.newretail.division.record.list.get"
}

func (r TaobaoNewretailDivisionRecordListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoNewretailDivisionRecordListGetRequest) SetParam(param *TopDivisionRecordReqDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoNewretailDivisionRecordListGetRequest) GetParam() *TopDivisionRecordReqDto {
    return r.param
}

