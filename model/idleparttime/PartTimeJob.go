package idleparttime

// PartTimeJob 
type PartTimeJob struct {
    // 工作地点
    Location   string `json:"location,omitempty" xml:"location,omitempty"`
    // 津贴
    Royalties   string `json:"royalties,omitempty" xml:"royalties,omitempty"`
    // 支付方式, 支付宝/微信/现金等
    PayWay   string `json:"pay_way,omitempty" xml:"pay_way,omitempty"`
    // 工作描述
    JobDescription   string `json:"job_description,omitempty" xml:"job_description,omitempty"`
    // 招聘人数
    RecruitCount   int64 `json:"recruit_count,omitempty" xml:"recruit_count,omitempty"`
    // 岗位名称: 该兼职岗位的名称
    JobTitle   string `json:"job_title,omitempty" xml:"job_title,omitempty"`
    // 招聘公司
    Company   string `json:"company,omitempty" xml:"company,omitempty"`
    // 工作时长
    WorkDuration   string `json:"work_duration,omitempty" xml:"work_duration,omitempty"`
    // 岗位种类: 该兼职信息的性质
    JobType   string `json:"job_type,omitempty" xml:"job_type,omitempty"`
    // 工作时间
    WorkTime   string `json:"work_time,omitempty" xml:"work_time,omitempty"`
    // 商品标题: 职位在 Feed 流中展示的标题,
    JobItemTitle   string `json:"job_item_title,omitempty" xml:"job_item_title,omitempty"`
    // 工资
    Salary   string `json:"salary,omitempty" xml:"salary,omitempty"`
    // 岗位id
    JobId   int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
    // 招聘公司描述
    CompanyDescription   string `json:"company_description,omitempty" xml:"company_description,omitempty"`
    // 企业logo, 是一个图片的URL
    CompanyLogo   string `json:"company_logo,omitempty" xml:"company_logo,omitempty"`
    // 岗位具体的要求和要求的类型
    JobRequirements   []PartTimeRequireSchema `json:"job_requirements,omitempty" xml:"job_requirements>part_time_require_schema,omitempty"`
    // 工作点点的经纬度
    Gps   string `json:"gps,omitempty" xml:"gps,omitempty"`
    // 岗位类目
    Category   string `json:"category,omitempty" xml:"category,omitempty"`
    // 联系人手机号码
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
    // 岗位是否新增, 0: 是 1: 否
    IsAdd   int64 `json:"is_add,omitempty" xml:"is_add,omitempty"`
    // 是否显示"取消报名" 0: 显示 1: 不显示
    ShowCancel   int64 `json:"show_cancel,omitempty" xml:"show_cancel,omitempty"`
    // 是否可以主动联系商家, 0: 可以, 1: 不可以
    ContactMerchant   int64 `json:"contact_merchant,omitempty" xml:"contact_merchant,omitempty"`
    // 发布岗位的经纬度
    PublishGps   string `json:"publish_gps,omitempty" xml:"publish_gps,omitempty"`
}
