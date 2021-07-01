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
type AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 零头入库：106；零头出库：210
    _billType   string
    // 单据编号
    _billCode   string
    // 单据时间:yyyy-MM-dd HH:mm:ss
    _billTime   string
    // 发货企业ID
    _fromRefUserId   string
    // 收货企业ID
    _toRefUserId   string
    // 委托企业ID
    _assRefEntId   string
    // 配送企业ID
    _disRefEntId   string
    // 药品ID
    _drugEntBaseInfoId   string
    // 生产日期：yyyy-MM-dd
    _produceDate   string
    // 有效期:yyyyMMdd
    _expireDate   string
    // 生产批次
    _produceBatchNo   string
    // 药品数量
    _inputAmount   string
}

// 初始化AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest对象
func NewAlibabaAlihealthDrugKytRemnantbillUploadRequest() *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest{
    return &AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.remnantbill.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// BillType Setter
// 零头入库：106；零头出库：210
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetBillType() string {
    return r._billType
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据时间:yyyy-MM-dd HH:mm:ss
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetBillTime() string {
    return r._billTime
}
// FromRefUserId Setter
// 发货企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
    r._fromRefUserId = _fromRefUserId
    r.Set("from_ref_user_id", _fromRefUserId)
    return nil
}

// FromRefUserId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetFromRefUserId() string {
    return r._fromRefUserId
}
// ToRefUserId Setter
// 收货企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetToRefUserId(_toRefUserId string) error {
    r._toRefUserId = _toRefUserId
    r.Set("to_ref_user_id", _toRefUserId)
    return nil
}

// ToRefUserId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetToRefUserId() string {
    return r._toRefUserId
}
// AssRefEntId Setter
// 委托企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetAssRefEntId(_assRefEntId string) error {
    r._assRefEntId = _assRefEntId
    r.Set("ass_ref_ent_id", _assRefEntId)
    return nil
}

// AssRefEntId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetAssRefEntId() string {
    return r._assRefEntId
}
// DisRefEntId Setter
// 配送企业ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetDisRefEntId(_disRefEntId string) error {
    r._disRefEntId = _disRefEntId
    r.Set("dis_ref_ent_id", _disRefEntId)
    return nil
}

// DisRefEntId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetDisRefEntId() string {
    return r._disRefEntId
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
    r._drugEntBaseInfoId = _drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetDrugEntBaseInfoId() string {
    return r._drugEntBaseInfoId
}
// ProduceDate Setter
// 生产日期：yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetProduceDate(_produceDate string) error {
    r._produceDate = _produceDate
    r.Set("produce_date", _produceDate)
    return nil
}

// ProduceDate Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetProduceDate() string {
    return r._produceDate
}
// ExpireDate Setter
// 有效期:yyyyMMdd
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetExpireDate(_expireDate string) error {
    r._expireDate = _expireDate
    r.Set("expire_date", _expireDate)
    return nil
}

// ExpireDate Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetExpireDate() string {
    return r._expireDate
}
// ProduceBatchNo Setter
// 生产批次
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
    r._produceBatchNo = _produceBatchNo
    r.Set("produce_batch_no", _produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetProduceBatchNo() string {
    return r._produceBatchNo
}
// InputAmount Setter
// 药品数量
func (r *AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) SetInputAmount(_inputAmount string) error {
    r._inputAmount = _inputAmount
    r.Set("input_amount", _inputAmount)
    return nil
}

// InputAmount Getter
func (r AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest) GetInputAmount() string {
    return r._inputAmount
}
