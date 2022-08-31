package wdk

// EmployeeBasic 结构体
type EmployeeBasic struct {
	// 私人邮箱
	EMail string `json:"e_mail,omitempty" xml:"e_mail,omitempty"`
	// 开户名
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 民事纠纷金额
	AmountOfCivilDisputes string `json:"amount_of_civil_disputes,omitempty" xml:"amount_of_civil_disputes,omitempty"`
	// 附件* {&#34;身份证-正面&#34;:&#34;url&#34;}
	Attachment string `json:"attachment,omitempty" xml:"attachment,omitempty"`
	// 银行账号
	BankAccount string `json:"bank_account,omitempty" xml:"bank_account,omitempty"`
	// 开户银行
	BankBranch string `json:"bank_branch,omitempty" xml:"bank_branch,omitempty"`
	// 银行国家
	BankCountry string `json:"bank_country,omitempty" xml:"bank_country,omitempty"`
	// 开户支行
	BankSubbranch string `json:"bank_subbranch,omitempty" xml:"bank_subbranch,omitempty"`
	// 基本工资
	BasicSalary string `json:"basic_salary,omitempty" xml:"basic_salary,omitempty"`
	// 生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// companiesEmployed
	CompaniesEmployed string `json:"companies_employed,omitempty" xml:"companies_employed,omitempty"`
	// 电脑操作
	ComputerOperate string `json:"computer_operate,omitempty" xml:"computer_operate,omitempty"`
	// 行政处罚或犯罪记录具体描述
	CrimeDesc string `json:"crime_desc,omitempty" xml:"crime_desc,omitempty"`
	// 最高学位
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 是否有重大疾病/手术记录
	DiseasesName string `json:"diseases_name,omitempty" xml:"diseases_name,omitempty"`
	// 最高学历
	Education string `json:"education,omitempty" xml:"education,omitempty"`
	// 员工子类
	EmpSubType string `json:"emp_sub_type,omitempty" xml:"emp_sub_type,omitempty"`
	// 员工类型
	EmpType string `json:"emp_type,omitempty" xml:"emp_type,omitempty"`
	// 曾受雇的结束时间
	EndTimeEmployed string `json:"end_time_employed,omitempty" xml:"end_time_employed,omitempty"`
	// 入职小管家
	EntryAssistant string `json:"entry_assistant,omitempty" xml:"entry_assistant,omitempty"`
	// 入职时间
	GmtEntry string `json:"gmt_entry,omitempty" xml:"gmt_entry,omitempty"`
	// 健康证到期日期
	GmtHealthCertEnd string `json:"gmt_health_cert_end,omitempty" xml:"gmt_health_cert_end,omitempty"`
	// 转正日期
	GmtOrigRegular string `json:"gmt_orig_regular,omitempty" xml:"gmt_orig_regular,omitempty"`
	// 住房公积金截止日期
	HouseFundPayTime string `json:"house_fund_pay_time,omitempty" xml:"house_fund_pay_time,omitempty"`
	// 户口类型
	HukouType string `json:"hukou_type,omitempty" xml:"hukou_type,omitempty"`
	// 是否有无犯罪记录
	IsCrime string `json:"is_crime,omitempty" xml:"is_crime,omitempty"`
	// 是否有残疾证
	IsDisability string `json:"is_disability,omitempty" xml:"is_disability,omitempty"`
	// 是否有重大疾病/手术记录
	IsDiseases string `json:"is_diseases,omitempty" xml:"is_diseases,omitempty"`
	// 是否曾受雇与本公司/联营商家/促销厂商
	IsEmployed string `json:"is_employed,omitempty" xml:"is_employed,omitempty"`
	// 是否内部推荐
	IsInternalReferral string `json:"is_internal_referral,omitempty" xml:"is_internal_referral,omitempty"`
	// 是否海外
	IsOversea string `json:"is_oversea,omitempty" xml:"is_oversea,omitempty"`
	// 是否缴纳住房公积金
	IsPayHouseFund string `json:"is_pay_house_fund,omitempty" xml:"is_pay_house_fund,omitempty"`
	// 是否缴纳社保
	IsPaySocialSecurity string `json:"is_pay_social_security,omitempty" xml:"is_pay_social_security,omitempty"`
	// 与其他单位仍有民事纠纷
	IsStillCivilDisputes string `json:"is_still_civil_disputes,omitempty" xml:"is_still_civil_disputes,omitempty"`
	// 曾担任的职务
	JobEmployed string `json:"job_employed,omitempty" xml:"job_employed,omitempty"`
	// 层级
	JobLevel string `json:"job_level,omitempty" xml:"job_level,omitempty"`
	// 曾受雇的工作地点
	LocationEmployed string `json:"location_employed,omitempty" xml:"location_employed,omitempty"`
	// 工作地点
	LocationNo string `json:"location_no,omitempty" xml:"location_no,omitempty"`
	// 婚姻状况
	Marriage string `json:"marriage,omitempty" xml:"marriage,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 民族
	Nation string `json:"nation,omitempty" xml:"nation,omitempty"`
	// 国籍
	NationCountry string `json:"nation_country,omitempty" xml:"nation_country,omitempty"`
	// 籍贯
	NativePlace string `json:"native_place,omitempty" xml:"native_place,omitempty"`
	// 部门组织
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 其他技能
	OtherSkill string `json:"other_skill,omitempty" xml:"other_skill,omitempty"`
	// 职位
	PostNo string `json:"post_no,omitempty" xml:"post_no,omitempty"`
	// 推荐人姓名
	ReferralName string `json:"referral_name,omitempty" xml:"referral_name,omitempty"`
	// 推荐人关系
	ReferralRelation string `json:"referral_relation,omitempty" xml:"referral_relation,omitempty"`
	// 计薪方式
	SalaryType string `json:"salary_type,omitempty" xml:"salary_type,omitempty"`
	// 服务公司
	ServiceCompany string `json:"service_company,omitempty" xml:"service_company,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 法人公司
	SignCompany string `json:"sign_company,omitempty" xml:"sign_company,omitempty"`
	// 招聘渠道
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 曾受雇的开始时间
	StartTimeEmployed string `json:"start_time_employed,omitempty" xml:"start_time_employed,omitempty"`
	// 实线主管
	SuperName string `json:"super_name,omitempty" xml:"super_name,omitempty"`
	// 失业保险缴纳截止日期
	UnemploymentPayTime string `json:"unemployment_pay_time,omitempty" xml:"unemployment_pay_time,omitempty"`
	// 虚线主管
	VirtName string `json:"virt_name,omitempty" xml:"virt_name,omitempty"`
	// 层级工资
	LevelSalary string `json:"level_salary,omitempty" xml:"level_salary,omitempty"`
	// 绩效奖金
	PerformancePay string `json:"performance_pay,omitempty" xml:"performance_pay,omitempty"`
	// 餐补
	MealAllowance string `json:"meal_allowance,omitempty" xml:"meal_allowance,omitempty"`
}
