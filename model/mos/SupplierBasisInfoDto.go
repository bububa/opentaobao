package mos

// SupplierBasisInfoDTO 
type SupplierBasisInfoDTO struct {
    // 类型名
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`
    // 状态
    ApprovalStatus   int64 `json:"approval_status,omitempty" xml:"approval_status,omitempty"`
    // 经营范围
    BusinessScope   string `json:"business_scope,omitempty" xml:"business_scope,omitempty"`
    // ocr认证id
    CertRecordId   string `json:"cert_record_id,omitempty" xml:"cert_record_id,omitempty"`
    // 变更类型：0-资质变更，1-主体变更
    ChangeType   int64 `json:"change_type,omitempty" xml:"change_type,omitempty"`
    // 公司名称
    CompanyName   string `json:"company_name,omitempty" xml:"company_name,omitempty"`
    // 公司性质，0：企业，1：个人
    CompanyNature   int64 `json:"company_nature,omitempty" xml:"company_nature,omitempty"`
    // 公司类型
    CompanyType   string `json:"company_type,omitempty" xml:"company_type,omitempty"`
    // 结束时间
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 公司成立日期
    EstablishDate   string `json:"establish_date,omitempty" xml:"establish_date,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 对象id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 分页-offset
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 操作人id
    OperateId   string `json:"operate_id,omitempty" xml:"operate_id,omitempty"`
    // 操作人name
    OperateName   string `json:"operate_name,omitempty" xml:"operate_name,omitempty"`
    // 经营期限结束日期
    OperatingEnd   string `json:"operating_end,omitempty" xml:"operating_end,omitempty"`
    // 经营期限开始日期
    OperatingStart   string `json:"operating_start,omitempty" xml:"operating_start,omitempty"`
    // 操作人id
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    // 操作人name
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    // 当前页
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 页数
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 审批流instanceId， 冪等用
    ProcessInstanceId   string `json:"process_instance_id,omitempty" xml:"process_instance_id,omitempty"`
    // 注册资本
    RegisteredCapital   string `json:"registered_capital,omitempty" xml:"registered_capital,omitempty"`
    // 幂等id
    RelationId   string `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
    // 供应商来源， 0：招商平台，1：供应商平台，2：初始化
    Source   int64 `json:"source,omitempty" xml:"source,omitempty"`
    // 开始时间
    StartTime   int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 供应商id
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 供应商状态，0：未入库，1：已入库，2：已转正
    SupplierStatus   int64 `json:"supplier_status,omitempty" xml:"supplier_status,omitempty"`
}
