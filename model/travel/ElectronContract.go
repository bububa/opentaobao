package travel

// ElectronContract 
type ElectronContract struct {
    // 拼团约定。0-不会采取拼团，1-可能采取拼团
    WillJoinGroup   int64 `json:"will_join_group,omitempty" xml:"will_join_group,omitempty"`
    // 是否支持电子合同。0-不支持，1-支持。新发布商品时不填默认为0
    SupportElectron   int64 `json:"support_electron,omitempty" xml:"support_electron,omitempty"`
    // 不成团约定，多选，英文逗号分隔。1-委托其他旅行社履行合同，2-延期出团，3-改签其他线路出行，4-解除合同
    GroupFailArrange   string `json:"group_fail_arrange,omitempty" xml:"group_fail_arrange,omitempty"`
}
