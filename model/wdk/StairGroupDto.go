package wdk

// StairGroupDto 结构体
type StairGroupDto struct {
	// 逻辑分组1的number为1，逻辑分组2的number为2，示例值[1&amp;2]：代表同时满足逻辑分组1和逻辑分组2时才可享受优惠
	ConditionExpress string `json:"condition_express,omitempty" xml:"condition_express,omitempty"`
	// 逻辑分组1的number为1，逻辑分组2的number为2，示例值[1|2]：代表逻辑分组1或者逻辑分组2可以享受优惠
	ActionExpress string `json:"action_express,omitempty" xml:"action_express,omitempty"`
	// 分组序号
	Number int64 `json:"number,omitempty" xml:"number,omitempty"`
	// 优惠门槛
	Condition *Condition `json:"condition,omitempty" xml:"condition,omitempty"`
	// 优惠效果
	Action *Action `json:"action,omitempty" xml:"action,omitempty"`
}
