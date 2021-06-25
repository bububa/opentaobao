package elife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询对账文件地址接口 APIRequest
taobao.elife.lifecard.recon

查询对账文件地址接口
*/
type TaobaoElifeLifecardReconRequest struct {
    model.Params

    // 对账日期(YYYYMMDD)
    opDate   string 

}

func NewTaobaoElifeLifecardReconRequest() *TaobaoElifeLifecardReconRequest{
    return &TaobaoElifeLifecardReconRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoElifeLifecardReconRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.recon"
}

func (r TaobaoElifeLifecardReconRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoElifeLifecardReconRequest) SetOpDate(opDate string) error {
    r.opDate = opDate
    r.Set("op_date", opDate)
    return nil
}

func (r TaobaoElifeLifecardReconRequest) GetOpDate() string {
    return r.opDate
}

