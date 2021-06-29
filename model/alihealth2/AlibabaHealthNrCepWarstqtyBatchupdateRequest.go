package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新ISV库存 APIRequest
alibaba.health.nr.cep.warstqty.batchupdate

青岛医保服务-ISV批量更新孔雀翎中库存数据
*/
type AlibabaHealthNrCepWarstqtyBatchupdateRequest struct {
    model.Params

    // 库存更新对象
    warStqtyList   []TopIsvStqtyLstDto 

}

func NewAlibabaHealthNrCepWarstqtyBatchupdateRequest() *AlibabaHealthNrCepWarstqtyBatchupdateRequest{
    return &AlibabaHealthNrCepWarstqtyBatchupdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHealthNrCepWarstqtyBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.warstqty.batchupdate"
}

func (r AlibabaHealthNrCepWarstqtyBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHealthNrCepWarstqtyBatchupdateRequest) SetWarStqtyList(warStqtyList []TopIsvStqtyLstDto) error {
    r.warStqtyList = warStqtyList
    r.Set("war_stqty_list", warStqtyList)
    return nil
}

func (r AlibabaHealthNrCepWarstqtyBatchupdateRequest) GetWarStqtyList() []TopIsvStqtyLstDto {
    return r.warStqtyList
}

