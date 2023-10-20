package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytupstorebillfileAPIRequest 上传零售出入库单(上传文件) API请求
// alibaba.alihealth.drug.kyt.upstorebillfile
//
// 上传零售出入库单(上传文件)
type AlibabaalihealthdrugkytupstorebillfileAPIRequest struct {
	model.Params
	// 单据编号
	_billCode string
	// 单据日期
	_billTime string
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
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[3,普药]
	_physicType int64
}

// NewAlibabaalihealthdrugkytupstorebillfileRequest 初始化AlibabaalihealthdrugkytupstorebillfileAPIRequest对象
func NewAlibabaalihealthdrugkytupstorebillfileRequest() *AlibabaalihealthdrugkytupstorebillfileAPIRequest {
	return &AlibabaalihealthdrugkytupstorebillfileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.upstorebillfile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据日期
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 上传企业的单位编码
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业(参与人标识，为null时会通过refEntId自动得到)
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者(key编号)
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 据提交者姓名
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetFileContent is FileContent Setter
// 文件内容
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetUploadFileName is UploadFileName Setter
// 文件名
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetUploadFileName(_uploadFileName string) error {
	r._uploadFileName = _uploadFileName
	r.Set("upload_file_name", _uploadFileName)
	return nil
}

// GetUploadFileName UploadFileName Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetUploadFileName() string {
	return r._uploadFileName
}

// SetClientType is ClientType Setter
// 客户端类型[暂定都写2]
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetClientType() string {
	return r._clientType
}

// SetBillType is BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型[3,普药]
func (r *AlibabaalihealthdrugkytupstorebillfileAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaalihealthdrugkytupstorebillfileAPIRequest) GetPhysicType() int64 {
	return r._physicType
}
