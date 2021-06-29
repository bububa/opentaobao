package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新ISV库存 API请求
alibaba.health.nr.cep.warstqty.batchupdate

青岛医保服务-ISV批量更新孔雀翎中库存数据
*/
type AlibabaHealthNrCepWarstqtyBatchupdateRequest struct {
    model.Params
    // 库存更新对象
    _warStqtyList   []TopIsvStqtyLstDto
}

// 初始化AlibabaHealthNrCepWarstqtyBatchupdateRequest对象
func NewAlibabaHealthNrCepWarstqtyBatchupdateRequest() *AlibabaHealthNrCepWarstqtyBatchupdateRequest{
    return &AlibabaHealthNrCepWarstqtyBatchupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepWarstqtyBatchupdateRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.warstqty.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepWarstqtyBatchupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarStqtyList Setter
// 库存更新对象
func (r *AlibabaHealthNrCepWarstqtyBatchupdateRequest) SetWarStqtyList(_warStqtyList []TopIsvStqtyLstDto) error {
    r._warStqtyList = _warStqtyList
    r.Set("war_stqty_list", _warStqtyList)
    return nil
}

// WarStqtyList Getter
func (r AlibabaHealthNrCepWarstqtyBatchupdateRequest) GetWarStqtyList() []TopIsvStqtyLstDto {
    return r._warStqtyList
}
