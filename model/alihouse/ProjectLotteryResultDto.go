package alihouse

// ProjectLotteryResultDto 结构体
type ProjectLotteryResultDto struct {
	// 摇号编号
	LotteryNum string `json:"lottery_num,omitempty" xml:"lottery_num,omitempty"`
	// 摇号姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 选房序号 摇号结果表才有数据
	SortNum string `json:"sort_num,omitempty" xml:"sort_num,omitempty"`
	// 身份证
	IdentityCard string `json:"identity_card,omitempty" xml:"identity_card,omitempty"`
}
