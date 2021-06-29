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
    _billCode   string
    // 单据日期
    _billTime   string
    // 单据类型[321,零售出库][322,疫苗接种]
    _billType   int64
    // 药品类型[3,普药]
    _physicType   int64
    // 上传企业的单位编码
    _refUserId   string
    // 发货企业(参与人标识，为null时会通过refEntId自动得到)
    _fromUserId   string
    // 单据提交者(key编号)
    _operIcCode   string
    // 据提交者姓名
    _operIcName   string
    // 文件内容
    _fileContent   string
    // 文件名
    _uploadFileName   string
    // 客户端类型[暂定都写2]
    _clientType   string
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
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据日期
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型[3,普药]
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetRefUserId() string {
    return r._refUserId
}
// FromUserId Setter
// 发货企业(参与人标识，为null时会通过refEntId自动得到)
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetFromUserId() string {
    return r._fromUserId
}
// OperIcCode Setter
// 单据提交者(key编号)
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 据提交者姓名
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetOperIcName() string {
    return r._operIcName
}
// FileContent Setter
// 文件内容
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetFileContent(_fileContent string) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetFileContent() string {
    return r._fileContent
}
// UploadFileName Setter
// 文件名
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetUploadFileName(_uploadFileName string) error {
    r._uploadFileName = _uploadFileName
    r.Set("upload_file_name", _uploadFileName)
    return nil
}

// UploadFileName Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetUploadFileName() string {
    return r._uploadFileName
}
// ClientType Setter
// 客户端类型[暂定都写2]
func (r *AlibabaAlihealthDrugKytUpstorebillfileRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileRequest) GetClientType() string {
    return r._clientType
}