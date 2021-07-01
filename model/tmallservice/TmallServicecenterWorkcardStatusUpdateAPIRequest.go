package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈服务的执行情况 API请求
tmall.servicecenter.workcard.status.update

1 如果服务商受理了此服务，修改合同状态为：已受理=3<br/>2 如果服务商没有受理此服务，修改合同状态为：已拒绝=10<br/>3 如果服务商执行了此服务，修改合同状态为：已执行=4<br/>4 如果服务商执行服务成功，修改合同状态为：已完成=5<br/>5 如果此工单为合同类型的工单，当服务商受理了此服务后，会进行分账
*/
type TmallServicecenterWorkcardStatusUpdateAPIRequest struct {
    model.Params
    // 工单id
    _workcardId   int64
    // 工单类型： 2（合同） 或者 1(任务）
    _type   *model.File
    // 目前仅支持5种状态的反馈：3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败。（所有状态列表： -1： 初始化 0： 生成 1： 生效 2： 申请 3： 受理 4： 执行 5： 成功 9： 结算 10： 拒绝 11： 失败 12 ： 撤销 13： 暂停 19： 终止）
    _status   *model.File
    // 买家id
    _buyerId   int64
    // api调用者
    _updater   string
    // 更新时间
    _updateDate   int64
    // 服务生效时间 ：工单类型为合同工单时，必选！
    _effectDate   int64
    // 服务失效时间 ：工单类型为合同工单时，必选！
    _expireDate   int64
    // 备注,256个字符以内
    _comments   string
    // 任务类工单，预约或者上门地址
    _address   string
    // 任务执行，预约联系人
    _contactName   string
    // 任务执行，预约联系人电话
    _contactPhone   string
    // 服务预约时间
    _serviceDate   int64
    // 服务完成时间
    _completeDate   int64
    // 服务凭证上传的图片URL链接，多个以;隔开
    _serviceVoucherPics   string
    // 属性定义。例如无忧退货服务，K-V对定义，每对KV用“;”分割，“:”号左边是key右边是value，value如果有多个则以“,”分割。 reasons   :  原因，可能有多个 succeedCount     :    取件成功个数 failedCount    :    取件失败个数 cancelCount      :     取件取消个数 totalCount       :      总取件个数，totalCount= succeedCount + failedCount + cancelCount
    _attribute   string
    // 服务商网点内部编码
    _serviceCenterCode   string
    // 服务商网点名字
    _serviceCenterName   string
    // 单元是分
    _serviceFee   int64
    // 是否上门
    _isVisit   bool
    // 说明
    _beforeServiceMemo   string
    // 说明
    _afterServiceMemo   string
    // 手机号码
    _phoneImei   string
    // 服务子状态：30 表示“服务已申请（上门）” 31表示“服务改约（上门）” 400表示“服务结果（待件上门）” 410表示“服务结果（拖机维修）” 411表示“服务结果（换机）” 420表示“服务结果（上门不可维修）”
    _subStatus   int64
    // 网点负责人联系电话
    _serviceCenterManagerPhone   string
    // 网点负责人
    _serviceCenterManagerName   string
    // 网点地址
    _serviceCenterAddress   string
    // 一个工单可能包含多件商品，比如空调可能有两台，录入每天机器的安装情况
    _workCardInstallDetailList   []WorkCardInstallDetail
    // json string。费用单位为分
    _serviceFeeDetail   string
    // 物流单号
    _expressCode   string
    // 物流公司名字
    _expressCompany   string
}

// 初始化TmallServicecenterWorkcardStatusUpdateAPIRequest对象
func NewTmallServicecenterWorkcardStatusUpdateRequest() *TmallServicecenterWorkcardStatusUpdateAPIRequest{
    return &TmallServicecenterWorkcardStatusUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.status.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// Type Setter
// 工单类型： 2（合同） 或者 1(任务）
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetType(_type *model.File) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetType() *model.File {
    return r._type
}
// Status Setter
// 目前仅支持5种状态的反馈：3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败。（所有状态列表： -1： 初始化 0： 生成 1： 生效 2： 申请 3： 受理 4： 执行 5： 成功 9： 结算 10： 拒绝 11： 失败 12 ： 撤销 13： 暂停 19： 终止）
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetStatus(_status *model.File) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetStatus() *model.File {
    return r._status
}
// BuyerId Setter
// 买家id
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetBuyerId(_buyerId int64) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetBuyerId() int64 {
    return r._buyerId
}
// Updater Setter
// api调用者
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetUpdater(_updater string) error {
    r._updater = _updater
    r.Set("updater", _updater)
    return nil
}

// Updater Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetUpdater() string {
    return r._updater
}
// UpdateDate Setter
// 更新时间
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetUpdateDate(_updateDate int64) error {
    r._updateDate = _updateDate
    r.Set("update_date", _updateDate)
    return nil
}

// UpdateDate Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetUpdateDate() int64 {
    return r._updateDate
}
// EffectDate Setter
// 服务生效时间 ：工单类型为合同工单时，必选！
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetEffectDate(_effectDate int64) error {
    r._effectDate = _effectDate
    r.Set("effect_date", _effectDate)
    return nil
}

// EffectDate Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetEffectDate() int64 {
    return r._effectDate
}
// ExpireDate Setter
// 服务失效时间 ：工单类型为合同工单时，必选！
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetExpireDate(_expireDate int64) error {
    r._expireDate = _expireDate
    r.Set("expire_date", _expireDate)
    return nil
}

// ExpireDate Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetExpireDate() int64 {
    return r._expireDate
}
// Comments Setter
// 备注,256个字符以内
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetComments(_comments string) error {
    r._comments = _comments
    r.Set("comments", _comments)
    return nil
}

// Comments Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetComments() string {
    return r._comments
}
// Address Setter
// 任务类工单，预约或者上门地址
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetAddress() string {
    return r._address
}
// ContactName Setter
// 任务执行，预约联系人
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetContactName(_contactName string) error {
    r._contactName = _contactName
    r.Set("contact_name", _contactName)
    return nil
}

// ContactName Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetContactName() string {
    return r._contactName
}
// ContactPhone Setter
// 任务执行，预约联系人电话
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetContactPhone(_contactPhone string) error {
    r._contactPhone = _contactPhone
    r.Set("contact_phone", _contactPhone)
    return nil
}

// ContactPhone Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetContactPhone() string {
    return r._contactPhone
}
// ServiceDate Setter
// 服务预约时间
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceDate(_serviceDate int64) error {
    r._serviceDate = _serviceDate
    r.Set("service_date", _serviceDate)
    return nil
}

// ServiceDate Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceDate() int64 {
    return r._serviceDate
}
// CompleteDate Setter
// 服务完成时间
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetCompleteDate(_completeDate int64) error {
    r._completeDate = _completeDate
    r.Set("complete_date", _completeDate)
    return nil
}

// CompleteDate Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetCompleteDate() int64 {
    return r._completeDate
}
// ServiceVoucherPics Setter
// 服务凭证上传的图片URL链接，多个以;隔开
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceVoucherPics(_serviceVoucherPics string) error {
    r._serviceVoucherPics = _serviceVoucherPics
    r.Set("service_voucher_pics", _serviceVoucherPics)
    return nil
}

// ServiceVoucherPics Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceVoucherPics() string {
    return r._serviceVoucherPics
}
// Attribute Setter
// 属性定义。例如无忧退货服务，K-V对定义，每对KV用“;”分割，“:”号左边是key右边是value，value如果有多个则以“,”分割。 reasons   :  原因，可能有多个 succeedCount     :    取件成功个数 failedCount    :    取件失败个数 cancelCount      :     取件取消个数 totalCount       :      总取件个数，totalCount= succeedCount + failedCount + cancelCount
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetAttribute(_attribute string) error {
    r._attribute = _attribute
    r.Set("attribute", _attribute)
    return nil
}

// Attribute Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetAttribute() string {
    return r._attribute
}
// ServiceCenterCode Setter
// 服务商网点内部编码
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceCenterCode(_serviceCenterCode string) error {
    r._serviceCenterCode = _serviceCenterCode
    r.Set("service_center_code", _serviceCenterCode)
    return nil
}

// ServiceCenterCode Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceCenterCode() string {
    return r._serviceCenterCode
}
// ServiceCenterName Setter
// 服务商网点名字
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceCenterName(_serviceCenterName string) error {
    r._serviceCenterName = _serviceCenterName
    r.Set("service_center_name", _serviceCenterName)
    return nil
}

// ServiceCenterName Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceCenterName() string {
    return r._serviceCenterName
}
// ServiceFee Setter
// 单元是分
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceFee(_serviceFee int64) error {
    r._serviceFee = _serviceFee
    r.Set("service_fee", _serviceFee)
    return nil
}

// ServiceFee Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceFee() int64 {
    return r._serviceFee
}
// IsVisit Setter
// 是否上门
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetIsVisit(_isVisit bool) error {
    r._isVisit = _isVisit
    r.Set("is_visit", _isVisit)
    return nil
}

// IsVisit Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetIsVisit() bool {
    return r._isVisit
}
// BeforeServiceMemo Setter
// 说明
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetBeforeServiceMemo(_beforeServiceMemo string) error {
    r._beforeServiceMemo = _beforeServiceMemo
    r.Set("before_service_memo", _beforeServiceMemo)
    return nil
}

// BeforeServiceMemo Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetBeforeServiceMemo() string {
    return r._beforeServiceMemo
}
// AfterServiceMemo Setter
// 说明
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetAfterServiceMemo(_afterServiceMemo string) error {
    r._afterServiceMemo = _afterServiceMemo
    r.Set("after_service_memo", _afterServiceMemo)
    return nil
}

// AfterServiceMemo Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetAfterServiceMemo() string {
    return r._afterServiceMemo
}
// PhoneImei Setter
// 手机号码
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetPhoneImei(_phoneImei string) error {
    r._phoneImei = _phoneImei
    r.Set("phone_imei", _phoneImei)
    return nil
}

// PhoneImei Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetPhoneImei() string {
    return r._phoneImei
}
// SubStatus Setter
// 服务子状态：30 表示“服务已申请（上门）” 31表示“服务改约（上门）” 400表示“服务结果（待件上门）” 410表示“服务结果（拖机维修）” 411表示“服务结果（换机）” 420表示“服务结果（上门不可维修）”
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetSubStatus(_subStatus int64) error {
    r._subStatus = _subStatus
    r.Set("sub_status", _subStatus)
    return nil
}

// SubStatus Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetSubStatus() int64 {
    return r._subStatus
}
// ServiceCenterManagerPhone Setter
// 网点负责人联系电话
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceCenterManagerPhone(_serviceCenterManagerPhone string) error {
    r._serviceCenterManagerPhone = _serviceCenterManagerPhone
    r.Set("service_center_manager_phone", _serviceCenterManagerPhone)
    return nil
}

// ServiceCenterManagerPhone Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceCenterManagerPhone() string {
    return r._serviceCenterManagerPhone
}
// ServiceCenterManagerName Setter
// 网点负责人
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceCenterManagerName(_serviceCenterManagerName string) error {
    r._serviceCenterManagerName = _serviceCenterManagerName
    r.Set("service_center_manager_name", _serviceCenterManagerName)
    return nil
}

// ServiceCenterManagerName Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceCenterManagerName() string {
    return r._serviceCenterManagerName
}
// ServiceCenterAddress Setter
// 网点地址
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceCenterAddress(_serviceCenterAddress string) error {
    r._serviceCenterAddress = _serviceCenterAddress
    r.Set("service_center_address", _serviceCenterAddress)
    return nil
}

// ServiceCenterAddress Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceCenterAddress() string {
    return r._serviceCenterAddress
}
// WorkCardInstallDetailList Setter
// 一个工单可能包含多件商品，比如空调可能有两台，录入每天机器的安装情况
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetWorkCardInstallDetailList(_workCardInstallDetailList []WorkCardInstallDetail) error {
    r._workCardInstallDetailList = _workCardInstallDetailList
    r.Set("work_card_install_detail_list", _workCardInstallDetailList)
    return nil
}

// WorkCardInstallDetailList Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetWorkCardInstallDetailList() []WorkCardInstallDetail {
    return r._workCardInstallDetailList
}
// ServiceFeeDetail Setter
// json string。费用单位为分
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetServiceFeeDetail(_serviceFeeDetail string) error {
    r._serviceFeeDetail = _serviceFeeDetail
    r.Set("service_fee_detail", _serviceFeeDetail)
    return nil
}

// ServiceFeeDetail Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetServiceFeeDetail() string {
    return r._serviceFeeDetail
}
// ExpressCode Setter
// 物流单号
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetExpressCode(_expressCode string) error {
    r._expressCode = _expressCode
    r.Set("express_code", _expressCode)
    return nil
}

// ExpressCode Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetExpressCode() string {
    return r._expressCode
}
// ExpressCompany Setter
// 物流公司名字
func (r *TmallServicecenterWorkcardStatusUpdateAPIRequest) SetExpressCompany(_expressCompany string) error {
    r._expressCompany = _expressCompany
    r.Set("express_company", _expressCompany)
    return nil
}

// ExpressCompany Getter
func (r TmallServicecenterWorkcardStatusUpdateAPIRequest) GetExpressCompany() string {
    return r._expressCompany
}
