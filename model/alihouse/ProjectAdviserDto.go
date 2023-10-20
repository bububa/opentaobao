package alihouse

import (
	"sync"
)

// ProjectAdviserDto 结构体
type ProjectAdviserDto struct {
	// 中文标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 1-区域认证小B
	CertificationServices []string `json:"certification_services,omitempty" xml:"certification_services>string,omitempty"`
	// 外部公司ID
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 姓名
	AdviserUserName string `json:"adviser_user_name,omitempty" xml:"adviser_user_name,omitempty"`
	// 手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 归属门店
	OuterShopId string `json:"outer_shop_id,omitempty" xml:"outer_shop_id,omitempty"`
	// 主机号
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 400分机号
	SubPhone string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
	// 个人简介
	Introduction string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// 服务评价
	ServiceEvaluation string `json:"service_evaluation,omitempty" xml:"service_evaluation,omitempty"`
	// 身份证号
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// 头像
	HeadUrl string `json:"head_url,omitempty" xml:"head_url,omitempty"`
	// 身份证正面
	CardFrontUrl string `json:"card_front_url,omitempty" xml:"card_front_url,omitempty"`
	// 身份证反面
	CardBackUrl string `json:"card_back_url,omitempty" xml:"card_back_url,omitempty"`
	// 名片
	BusinessCard string `json:"business_card,omitempty" xml:"business_card,omitempty"`
	// 执业资格证图片
	AgentLicenseUrl string `json:"agent_license_url,omitempty" xml:"agent_license_url,omitempty"`
	// 执业资格证
	AgentLicenseNumber string `json:"agent_license_number,omitempty" xml:"agent_license_number,omitempty"`
	// 从业信息卡图片
	AgentEmployUrl string `json:"agent_employ_url,omitempty" xml:"agent_employ_url,omitempty"`
	// 从业信息卡编号
	AgentEmployNumber string `json:"agent_employ_number,omitempty" xml:"agent_employ_number,omitempty"`
	// 核验返回信息
	VerificationInfo string `json:"verification_info,omitempty" xml:"verification_info,omitempty"`
	// 经纪人子账号名称
	SubUserNick string `json:"sub_user_nick,omitempty" xml:"sub_user_nick,omitempty"`
	// 负责的业务域
	PartakeBusiness string `json:"partake_business,omitempty" xml:"partake_business,omitempty"`
	// 外部门店ID
	OutStoreId string `json:"out_store_id,omitempty" xml:"out_store_id,omitempty"`
	// 名字
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 经纪人职业
	Profession string `json:"profession,omitempty" xml:"profession,omitempty"`
	// 从业资格备案信息地址
	AgentEmployNumberUrl string `json:"agent_employ_number_url,omitempty" xml:"agent_employ_number_url,omitempty"`
	// 身份标识
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// eCode
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 公司简称
	CompanyShortName string `json:"company_short_name,omitempty" xml:"company_short_name,omitempty"`
	// 公司名称
	Company string `json:"company,omitempty" xml:"company,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 审核通过或驳回原因
	AuditReason string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 经纪人ID
	BrokerId int64 `json:"broker_id,omitempty" xml:"broker_id,omitempty"`
	// 性别
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 工作年限
	WorkYear int64 `json:"work_year,omitempty" xml:"work_year,omitempty"`
	// 在职状态
	JobStatus int64 `json:"job_status,omitempty" xml:"job_status,omitempty"`
	// 核验状态
	VerificationStatus int64 `json:"verification_status,omitempty" xml:"verification_status,omitempty"`
	// 实名认证状态 0-未审 1-通过 2-拒绝
	RealNameStatus int64 `json:"real_name_status,omitempty" xml:"real_name_status,omitempty"`
	// 经纪人附属属性
	BrokerAttach *ProjectAdviserAttachDto `json:"broker_attach,omitempty" xml:"broker_attach,omitempty"`
	// 经纪人教育水平
	Education int64 `json:"education,omitempty" xml:"education,omitempty"`
	// 经纪人类型，3-社区志愿者 4-客服，5-委托管家（无忧卖房） 6-委托客服（无忧卖房） 7-公寓管家 8-交易员；不填默认为经纪人
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 诚信状态1-优秀 2-诚信
	SincerityStatus int64 `json:"sincerity_status,omitempty" xml:"sincerity_status,omitempty"`
	// 经纪人是否有从业资格证
	AgentEmployStatus int64 `json:"agent_employ_status,omitempty" xml:"agent_employ_status,omitempty"`
	// 是否测试 0-否 1-是
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 是否可接待客户，1 - 是，0 - 否，接待模式3.0时，此字段必填。
	IsReceiveCustomer int64 `json:"is_receive_customer,omitempty" xml:"is_receive_customer,omitempty"`
	// 区域接待状态0-不可接待 1-可接待
	RegionReceptiveStatus int64 `json:"region_receptive_status,omitempty" xml:"region_receptive_status,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 状态：0无效，1有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 默认0，1内场顾问，2外场分销
	Role int64 `json:"role,omitempty" xml:"role,omitempty"`
	// 接待模式，2 - 接待模式2.0，3 - 接待模式3.0
	ReceiveModel int64 `json:"receive_model,omitempty" xml:"receive_model,omitempty"`
	// 审批状态
	AuditStatus int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
	// 请求时间戳，精确到毫秒
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 金牌顾问编辑的审核状态
	EditAuditStatus int64 `json:"edit_audit_status,omitempty" xml:"edit_audit_status,omitempty"`
	// 是否拥有淘宝账号
	HasTbAccount bool `json:"has_tb_account,omitempty" xml:"has_tb_account,omitempty"`
}

var poolProjectAdviserDto = sync.Pool{
	New: func() any {
		return new(ProjectAdviserDto)
	},
}

// GetProjectAdviserDto() 从对象池中获取ProjectAdviserDto
func GetProjectAdviserDto() *ProjectAdviserDto {
	return poolProjectAdviserDto.Get().(*ProjectAdviserDto)
}

// ReleaseProjectAdviserDto 释放ProjectAdviserDto
func ReleaseProjectAdviserDto(v *ProjectAdviserDto) {
	v.Tags = v.Tags[:0]
	v.CertificationServices = v.CertificationServices[:0]
	v.OuterCompanyId = ""
	v.AdviserUserName = ""
	v.MobilePhone = ""
	v.OuterShopId = ""
	v.MainPhone = ""
	v.SubPhone = ""
	v.Introduction = ""
	v.ServiceEvaluation = ""
	v.IdNumber = ""
	v.HeadUrl = ""
	v.CardFrontUrl = ""
	v.CardBackUrl = ""
	v.BusinessCard = ""
	v.AgentLicenseUrl = ""
	v.AgentLicenseNumber = ""
	v.AgentEmployUrl = ""
	v.AgentEmployNumber = ""
	v.VerificationInfo = ""
	v.SubUserNick = ""
	v.PartakeBusiness = ""
	v.OutStoreId = ""
	v.UserName = ""
	v.OuterConsultantId = ""
	v.Profession = ""
	v.AgentEmployNumberUrl = ""
	v.Identity = ""
	v.ECode = ""
	v.OuterStoreId = ""
	v.ShopName = ""
	v.CompanyShortName = ""
	v.Company = ""
	v.OuterId = ""
	v.AuditReason = ""
	v.StoreId = 0
	v.BrokerId = 0
	v.Gender = 0
	v.WorkYear = 0
	v.JobStatus = 0
	v.VerificationStatus = 0
	v.RealNameStatus = 0
	v.BrokerAttach = nil
	v.Education = 0
	v.Type = 0
	v.SincerityStatus = 0
	v.AgentEmployStatus = 0
	v.IsTest = 0
	v.EtcVersion = 0
	v.IsReceiveCustomer = 0
	v.RegionReceptiveStatus = 0
	v.Sort = 0
	v.Status = 0
	v.Role = 0
	v.ReceiveModel = 0
	v.AuditStatus = 0
	v.Version = 0
	v.EditAuditStatus = 0
	v.HasTbAccount = false
	poolProjectAdviserDto.Put(v)
}
