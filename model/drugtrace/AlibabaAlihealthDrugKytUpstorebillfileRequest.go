package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传零售出入库单(上传文件) API请求
alibaba.alihealth.drug.kyt.upstorebillfile

上传零售出入库单(上传文件)
*/
type AlibabaAlihealthDrugKytUpstorebillfileRequest struct {
    model.Params
    // 单据编号
    billCode   string
    // 单据日期
    billTime   string
    // 单据类型[321,零售出库][322,疫苗接种]
    billType   int64
    // 药品类型[3,普药]
    physicType   int64
    // 上传企业的单位编码
    refUserId   string
    // 发货企业(参与人标识，为null时会通过refEntId自动得到)
    fromUserId   string
    // 单据提交者(key编号)
    operIcCode   string
    // 据提交者姓名
    operIcName   string
    // 文件内容
    fileContent   string
    // 文件名
    uploadFileName   string
    // 客户端类型[暂定都写2]
    clientType   string
}

// 初始化AlibabaAlihealthDrugKytUpstorebillfileRequest对象
func NewAlibabaAlihealthDrugKytUpstorebillfileRequest() *AlibabaAlihealthDrugKytUpstorebillfileRequest{
    return &AlibabaAlihealthDrugKytUpstorebillfileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.upstorebillfile"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据日期
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillTime() string {
    return r.billTime
}
// BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillType() int64 {
    return r.billType
}
// PhysicType Setter
// 药品类型[3,普药]
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetPhysicType() int64 {
    return r.physicType
}
// RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetRefUserId() string {
    return r.refUserId
}
// FromUserId Setter
// 发货企业(参与人标识，为null时会通过refEntId自动得到)
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetFromUserId() string {
    return r.fromUserId
}
// OperIcCode Setter
// 单据提交者(key编号)
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetOperIcCode() string {
    return r.operIcCode
}
// OperIcName Setter
// 据提交者姓名
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetOperIcName() string {
    return r.operIcName
}
// FileContent Setter
// 文件内容
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetFileContent() string {
    return r.fileContent
}
// UploadFileName Setter
// 文件名
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetUploadFileName(uploadFileName string) error {
    r.uploadFileName = uploadFileName
    r.Set("upload_file_name", uploadFileName)
    return nil
}

// UploadFileName Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetUploadFileName() string {
    return r.uploadFileName
}
// ClientType Setter
// 客户端类型[暂定都写2]
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetClientType() string {
    return r.clientType
}
