package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈服务的执行情况 APIRequest
tmall.servicecenter.workcard.status.update

1 如果服务商受理了此服务，修改合同状态为：已受理=3<br/>2 如果服务商没有受理此服务，修改合同状态为：已拒绝=10<br/>3 如果服务商执行了此服务，修改合同状态为：已执行=4<br/>4 如果服务商执行服务成功，修改合同状态为：已完成=5<br/>5 如果此工单为合同类型的工单，当服务商受理了此服务后，会进行分账
*/
type TmallServicecenterWorkcardStatusUpdateRequest struct {
    model.Params

    // 工单id
    workcardId   int64 

    // 工单类型： 2（合同） 或者 1(任务）
    type   *model.File 

    // 目前仅支持5种状态的反馈：3=授理， 10=拒绝 ，4=执行 ，5=成功，11=失败。（所有状态列表： -1： 初始化 0： 生成 1： 生效 2： 申请 3： 受理 4： 执行 5： 成功 9： 结算 10： 拒绝 11： 失败 12 ： 撤销 13： 暂停 19： 终止）
    status   *model.File 

    // 买家id
    buyerId   int64 

    // api调用者
    updater   string 

    // 更新时间
    updateDate   int64 

    // 服务生效时间 ：工单类型为合同工单时，必选！
    effectDate   int64 

    // 服务失效时间 ：工单类型为合同工单时，必选！
    expireDate   int64 

    // 备注,256个字符以内
    comments   string 

    // 任务类工单，预约或者上门地址
    address   string 

    // 任务执行，预约联系人
    contactName   string 

    // 任务执行，预约联系人电话
    contactPhone   string 

    // 服务预约时间
    serviceDate   int64 

    // 服务完成时间
    completeDate   int64 

    // 服务凭证上传的图片URL链接，多个以;隔开
    serviceVoucherPics   string 

    // 属性定义。例如无忧退货服务，K-V对定义，每对KV用“;”分割，“:”号左边是key右边是value，value如果有多个则以“,”分割。 reasons   :  原因，可能有多个 succeedCount     :    取件成功个数 failedCount    :    取件失败个数 cancelCount      :     取件取消个数 totalCount       :      总取件个数，totalCount= succeedCount + failedCount + cancelCount
    attribute   string 

    // 服务商网点内部编码
    serviceCenterCode   string 

    // 服务商网点名字
    serviceCenterName   string 

    // 单元是分
    serviceFee   int64 

    // 是否上门
    isVisit   bool 

    // 说明
    beforeServiceMemo   string 

    // 说明
    afterServiceMemo   string 

    // 手机号码
    phoneImei   string 

    // 服务子状态：30 表示“服务已申请（上门）” 31表示“服务改约（上门）” 400表示“服务结果（待件上门）” 410表示“服务结果（拖机维修）” 411表示“服务结果（换机）” 420表示“服务结果（上门不可维修）”
    subStatus   int64 

    // 网点负责人联系电话
    serviceCenterManagerPhone   string 

    // 网点负责人
    serviceCenterManagerName   string 

    // 网点地址
    serviceCenterAddress   string 

    // 一个工单可能包含多件商品，比如空调可能有两台，录入每天机器的安装情况
    workCardInstallDetailList   []WorkCardInstallDetail 

    // json string。费用单位为分
    serviceFeeDetail   string 

    // 物流单号
    expressCode   string 

    // 物流公司名字
    expressCompany   string 

}

func NewTmallServicecenterWorkcardStatusUpdateRequest() *TmallServicecenterWorkcardStatusUpdateRequest{
    return &TmallServicecenterWorkcardStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.status.update"
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetWorkcardId() int64 {
    return r.workcardId
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetType(type *model.File) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetType() *model.File {
    return r.type
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetStatus(status *model.File) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetStatus() *model.File {
    return r.status
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetBuyerId() int64 {
    return r.buyerId
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetUpdater(updater string) error {
    r.updater = updater
    r.Set("updater", updater)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetUpdater() string {
    return r.updater
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetUpdateDate(updateDate int64) error {
    r.updateDate = updateDate
    r.Set("update_date", updateDate)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetUpdateDate() int64 {
    return r.updateDate
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetEffectDate(effectDate int64) error {
    r.effectDate = effectDate
    r.Set("effect_date", effectDate)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetEffectDate() int64 {
    return r.effectDate
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetExpireDate(expireDate int64) error {
    r.expireDate = expireDate
    r.Set("expire_date", expireDate)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetExpireDate() int64 {
    return r.expireDate
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetComments(comments string) error {
    r.comments = comments
    r.Set("comments", comments)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetComments() string {
    return r.comments
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetAddress() string {
    return r.address
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetContactName(contactName string) error {
    r.contactName = contactName
    r.Set("contact_name", contactName)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetContactName() string {
    return r.contactName
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetContactPhone(contactPhone string) error {
    r.contactPhone = contactPhone
    r.Set("contact_phone", contactPhone)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetContactPhone() string {
    return r.contactPhone
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceDate(serviceDate int64) error {
    r.serviceDate = serviceDate
    r.Set("service_date", serviceDate)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceDate() int64 {
    return r.serviceDate
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetCompleteDate(completeDate int64) error {
    r.completeDate = completeDate
    r.Set("complete_date", completeDate)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetCompleteDate() int64 {
    return r.completeDate
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceVoucherPics(serviceVoucherPics string) error {
    r.serviceVoucherPics = serviceVoucherPics
    r.Set("service_voucher_pics", serviceVoucherPics)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceVoucherPics() string {
    return r.serviceVoucherPics
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetAttribute(attribute string) error {
    r.attribute = attribute
    r.Set("attribute", attribute)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetAttribute() string {
    return r.attribute
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceCenterCode(serviceCenterCode string) error {
    r.serviceCenterCode = serviceCenterCode
    r.Set("service_center_code", serviceCenterCode)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceCenterCode() string {
    return r.serviceCenterCode
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceCenterName(serviceCenterName string) error {
    r.serviceCenterName = serviceCenterName
    r.Set("service_center_name", serviceCenterName)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceCenterName() string {
    return r.serviceCenterName
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceFee(serviceFee int64) error {
    r.serviceFee = serviceFee
    r.Set("service_fee", serviceFee)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceFee() int64 {
    return r.serviceFee
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetIsVisit(isVisit bool) error {
    r.isVisit = isVisit
    r.Set("is_visit", isVisit)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetIsVisit() bool {
    return r.isVisit
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetBeforeServiceMemo(beforeServiceMemo string) error {
    r.beforeServiceMemo = beforeServiceMemo
    r.Set("before_service_memo", beforeServiceMemo)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetBeforeServiceMemo() string {
    return r.beforeServiceMemo
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetAfterServiceMemo(afterServiceMemo string) error {
    r.afterServiceMemo = afterServiceMemo
    r.Set("after_service_memo", afterServiceMemo)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetAfterServiceMemo() string {
    return r.afterServiceMemo
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetPhoneImei(phoneImei string) error {
    r.phoneImei = phoneImei
    r.Set("phone_imei", phoneImei)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetPhoneImei() string {
    return r.phoneImei
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetSubStatus(subStatus int64) error {
    r.subStatus = subStatus
    r.Set("sub_status", subStatus)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetSubStatus() int64 {
    return r.subStatus
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceCenterManagerPhone(serviceCenterManagerPhone string) error {
    r.serviceCenterManagerPhone = serviceCenterManagerPhone
    r.Set("service_center_manager_phone", serviceCenterManagerPhone)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceCenterManagerPhone() string {
    return r.serviceCenterManagerPhone
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceCenterManagerName(serviceCenterManagerName string) error {
    r.serviceCenterManagerName = serviceCenterManagerName
    r.Set("service_center_manager_name", serviceCenterManagerName)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceCenterManagerName() string {
    return r.serviceCenterManagerName
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceCenterAddress(serviceCenterAddress string) error {
    r.serviceCenterAddress = serviceCenterAddress
    r.Set("service_center_address", serviceCenterAddress)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceCenterAddress() string {
    return r.serviceCenterAddress
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetWorkCardInstallDetailList(workCardInstallDetailList []WorkCardInstallDetail) error {
    r.workCardInstallDetailList = workCardInstallDetailList
    r.Set("work_card_install_detail_list", workCardInstallDetailList)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetWorkCardInstallDetailList() []WorkCardInstallDetail {
    return r.workCardInstallDetailList
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetServiceFeeDetail(serviceFeeDetail string) error {
    r.serviceFeeDetail = serviceFeeDetail
    r.Set("service_fee_detail", serviceFeeDetail)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetServiceFeeDetail() string {
    return r.serviceFeeDetail
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetExpressCode(expressCode string) error {
    r.expressCode = expressCode
    r.Set("express_code", expressCode)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetExpressCode() string {
    return r.expressCode
}

func (r *TmallServicecenterWorkcardStatusUpdateRequest) SetExpressCompany(expressCompany string) error {
    r.expressCompany = expressCompany
    r.Set("express_company", expressCompany)
    return nil
}

func (r TmallServicecenterWorkcardStatusUpdateRequest) GetExpressCompany() string {
    return r.expressCompany
}

