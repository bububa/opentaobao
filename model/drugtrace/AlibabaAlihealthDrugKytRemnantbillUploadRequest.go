package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零头出入库单据上传 API请求
alibaba.alihealth.drug.kyt.remnantbill.upload

零头出入库单据上传
*/
type AlibabaAlihealthDrugKytRemnantbillUploadRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 零头入库：106；零头出库：210
    billType   string
    // 单据编号
    billCode   string
    // 单据时间:yyyy-MM-dd HH:mm:ss
    billTime   string
    // 发货企业ID
    fromRefUserId   string
    // 收货企业ID
    toRefUserId   string
    // 委托企业ID
    assRefEntId   string
    // 配送企业ID
    disRefEntId   string
    // 药品ID
    drugEntBaseInfoId   string
    // 生产日期：yyyy-MM-dd
    produceDate   string
    // 有效期:yyyyMMdd
    expireDate   string
    // 生产批次
    produceBatchNo   string
    // 药品数量
    inputAmount   string
}

// 初始化AlibabaAlihealthDrugKytRemnantbillUploadRequest对象
func NewAlibabaAlihealthDrugKytRemnantbillUploadRequest() *AlibabaAlihealthDrugKytRemnantbillUploadRequest{
    return &AlibabaAlihealthDrugKytRemnantbillUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.remnantbill.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetRefEntId() string {
    return r.refEntId
}
// BillType Setter
// 零头入库：106；零头出库：210
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetBillType() string {
    return r.billType
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据时间:yyyy-MM-dd HH:mm:ss
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetBillTime() string {
    return r.billTime
}
// FromRefUserId Setter
// 发货企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetFromRefUserId(fromRefUserId string) error {
    r.fromRefUserId = fromRefUserId
    r.Set("from_ref_user_id", fromRefUserId)
    return nil
}

// FromRefUserId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetFromRefUserId() string {
    return r.fromRefUserId
}
// ToRefUserId Setter
// 收货企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetToRefUserId(toRefUserId string) error {
    r.toRefUserId = toRefUserId
    r.Set("to_ref_user_id", toRefUserId)
    return nil
}

// ToRefUserId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetToRefUserId() string {
    return r.toRefUserId
}
// AssRefEntId Setter
// 委托企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetAssRefEntId(assRefEntId string) error {
    r.assRefEntId = assRefEntId
    r.Set("ass_ref_ent_id", assRefEntId)
    return nil
}

// AssRefEntId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetAssRefEntId() string {
    return r.assRefEntId
}
// DisRefEntId Setter
// 配送企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetDisRefEntId(disRefEntId string) error {
    r.disRefEntId = disRefEntId
    r.Set("dis_ref_ent_id", disRefEntId)
    return nil
}

// DisRefEntId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetDisRefEntId() string {
    return r.disRefEntId
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetDrugEntBaseInfoId(drugEntBaseInfoId string) error {
    r.drugEntBaseInfoId = drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetDrugEntBaseInfoId() string {
    return r.drugEntBaseInfoId
}
// ProduceDate Setter
// 生产日期：yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetProduceDate(produceDate string) error {
    r.produceDate = produceDate
    r.Set("produce_date", produceDate)
    return nil
}

// ProduceDate Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetProduceDate() string {
    return r.produceDate
}
// ExpireDate Setter
// 有效期:yyyyMMdd
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetExpireDate(expireDate string) error {
    r.expireDate = expireDate
    r.Set("expire_date", expireDate)
    return nil
}

// ExpireDate Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetExpireDate() string {
    return r.expireDate
}
// ProduceBatchNo Setter
// 生产批次
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetProduceBatchNo(produceBatchNo string) error {
    r.produceBatchNo = produceBatchNo
    r.Set("produce_batch_no", produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetProduceBatchNo() string {
    return r.produceBatchNo
}
// InputAmount Setter
// 药品数量
func (r *AlibabaAlihealthDrugKytRemnantbillUploadRequest) SetInputAmount(inputAmount string) error {
    r.inputAmount = inputAmount
    r.Set("input_amount", inputAmount)
    return nil
}

// InputAmount Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadRequest) GetInputAmount() string {
    return r.inputAmount
}
