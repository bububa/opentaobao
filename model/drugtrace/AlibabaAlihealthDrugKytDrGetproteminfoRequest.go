package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗，获取生产企业的存储和运输温度 API请求
alibaba.alihealth.drug.kyt.dr.getproteminfo

疫苗，获取生产企业的存储和运输温度
*/
type AlibabaAlihealthDrugKytDrGetproteminfoRequest struct {
    model.Params
    // 调用企业ID
    refEntId   string
    // 药品ID
    drugEntBaseInfoId   string
    // 批次编号
    batchNo   string
    // 出库单号
    billOutCode   string
}

// 初始化AlibabaAlihealthDrugKytDrGetproteminfoRequest对象
func NewAlibabaAlihealthDrugKytDrGetproteminfoRequest() *AlibabaAlihealthDrugKytDrGetproteminfoRequest{
    return &AlibabaAlihealthDrugKytDrGetproteminfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetproteminfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getproteminfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetproteminfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 调用企业ID
func (r *AlibabaAlihealthDrugKytDrGetproteminfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetproteminfoRequest) GetRefEntId() string {
    return r.refEntId
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytDrGetproteminfoRequest) SetDrugEntBaseInfoId(drugEntBaseInfoId string) error {
    r.drugEntBaseInfoId = drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytDrGetproteminfoRequest) GetDrugEntBaseInfoId() string {
    return r.drugEntBaseInfoId
}
// BatchNo Setter
// 批次编号
func (r *AlibabaAlihealthDrugKytDrGetproteminfoRequest) SetBatchNo(batchNo string) error {
    r.batchNo = batchNo
    r.Set("batch_no", batchNo)
    return nil
}

// BatchNo Getter
func (r AlibabaAlihealthDrugKytDrGetproteminfoRequest) GetBatchNo() string {
    return r.batchNo
}
// BillOutCode Setter
// 出库单号
func (r *AlibabaAlihealthDrugKytDrGetproteminfoRequest) SetBillOutCode(billOutCode string) error {
    r.billOutCode = billOutCode
    r.Set("bill_out_code", billOutCode)
    return nil
}

// BillOutCode Getter
func (r AlibabaAlihealthDrugKytDrGetproteminfoRequest) GetBillOutCode() string {
    return r.billOutCode
}
