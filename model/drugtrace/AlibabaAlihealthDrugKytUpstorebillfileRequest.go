package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传零售出入库单(上传文件) APIRequest
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

func NewAlibabaAlihealthDrugKytUpstorebillfileRequest() *AlibabaAlihealthDrugKytUpstorebillfileRequest{
    return &AlibabaAlihealthDrugKytUpstorebillfileRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.upstorebillfile"
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetFileContent() string {
    return r.fileContent
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetUploadFileName(uploadFileName string) error {
    r.uploadFileName = uploadFileName
    r.Set("upload_file_name", uploadFileName)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetUploadFileName() string {
    return r.uploadFileName
}

func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetClientType() string {
    return r.clientType
}

