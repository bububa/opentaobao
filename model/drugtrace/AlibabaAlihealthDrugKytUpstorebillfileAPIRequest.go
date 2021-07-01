package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpstorebillfileAPIRequest
上传零售出入库单(上传文件) API请求
alibaba.alihealth.drug.kyt.upstorebillfile

上传零售出入库单(上传文件) */
type AlibabaAlihealthDrugKytUpstorebillfileAPIRequest struct {
	model.Params
	// 单据编号
	_billCode string
	// 单据日期
	_billTime string
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[3,普药]
	_physicType int64
	// 上传企业的单位编码
	_refUserId string
	// 发货企业(参与人标识，为null时会通过refEntId自动得到)
	_fromUserId string
	// 单据提交者(key编号)
	_operIcCode string
	// 据提交者姓名
	_operIcName string
	// 文件内容
	_fileContent string
	// 文件名
	_uploadFileName string
	// 客户端类型[暂定都写2]
	_clientType string
}

// NewAlibabaAlihealthDrugKytUpstorebillfileRequest 初始化AlibabaAlihealthDrugKytUpstorebillfileAPIRequest对象
func NewAlibabaAlihealthDrugKytUpstorebillfileRequest() *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest {
	return &AlibabaAlihealthDrugKytUpstorebillfileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.upstorebillfile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is BillTime Setter
// 单据日期
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// Get BillTime Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetBillTime() string {
	return r._billTime
}

// Set is BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// Get BillType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetBillType() int64 {
	return r._billType
}

// Set is PhysicType Setter
// 药品类型[3,普药]
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// Get PhysicType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// Set is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// Get RefUserId Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// Set is FromUserId Setter
// 发货企业(参与人标识，为null时会通过refEntId自动得到)
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// Get FromUserId Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// Set is OperIcCode Setter
// 单据提交者(key编号)
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// Get OperIcCode Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// Set is OperIcName Setter
// 据提交者姓名
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// Get OperIcName Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// Set is FileContent Setter
// 文件内容
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// Get FileContent Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetFileContent() string {
	return r._fileContent
}

// Set is UploadFileName Setter
// 文件名
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetUploadFileName(_uploadFileName string) error {
	r._uploadFileName = _uploadFileName
	r.Set("upload_file_name", _uploadFileName)
	return nil
}

// Get UploadFileName Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetUploadFileName() string {
	return r._uploadFileName
}

// Set is ClientType Setter
// 客户端类型[暂定都写2]
func (r *AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// Get ClientType Getter
func (r AlibabaAlihealthDrugKytUpstorebillfileAPIRequest) GetClientType() string {
	return r._clientType
}
