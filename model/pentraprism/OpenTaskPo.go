package pentraprism

// OpenTaskPo 结构体
type OpenTaskPo struct {
	// 应用名称
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 应用版本
	AppVersion string `json:"app_version,omitempty" xml:"app_version,omitempty"`
	// 领奖励位置，默认为null
	AwardIndex int64 `json:"award_index,omitempty" xml:"award_index,omitempty"`
	// 任务系统后台配置投放ID
	DeliveryId int64 `json:"delivery_id,omitempty" xml:"delivery_id,omitempty"`
	// 来源详细标识
	FromAppName string `json:"from_app_name,omitempty" xml:"from_app_name,omitempty"`
	// 是否忽略任务规则
	IgnoreRules bool `json:"ignore_rules,omitempty" xml:"ignore_rules,omitempty"`
	// 幂等ID，业务控制
	ImplId string `json:"impl_id,omitempty" xml:"impl_id,omitempty"`
	// 做任务时间
	Now string `json:"now,omitempty" xml:"now,omitempty"`
	// 操作系统名称
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 任务系统后台配置场景ID
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 用户类别，微博用户写死userType=2
	UserType int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 分享任务的token，回流的时候传入
	FromToken string `json:"from_token,omitempty" xml:"from_token,omitempty"`
	// 请求对应任务的token，从五棱镜后台生成
	OpenToken string `json:"open_token,omitempty" xml:"open_token,omitempty"`
}
