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
type AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest struct {
    model.Params
    // 库存更新对象
    _warStqtyList   []TopIsvStqtyLstDTO
}

// 初始化AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest对象
func NewAlibabaHealthNrCepWarstqtyBatchupdateRequest() *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest{
    return &AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.warstqty.batchupdate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WarStqtyList Setter
// 库存更新对象
func (r *AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) SetWarStqtyList(_warStqtyList []TopIsvStqtyLstDTO) error {
    r._warStqtyList = _warStqtyList
    r.Set("war_stqty_list", _warStqtyList)
    return nil
}

// WarStqtyList Getter
func (r AlibabaHealthNrCepWarstqtyBatchupdateAPIRequest) GetWarStqtyList() []TopIsvStqtyLstDTO {
    return r._warStqtyList
}
