package simba

// Wordscorelist 
type Wordscorelist struct {

    // 移动创意质量(创意效果)
    WirelessCreativescore   string `json:"wireless_creativescore,omitempty"`

    // 移动店铺质量(账户表现)
    WirelessCustscore   string `json:"wireless_custscore,omitempty"`

    // 相关性，同kwscore
    Relescore   string `json:"relescore,omitempty"`

    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 词是否有首屏展示机会。值含义(4: 有展现机会,且能上首屏 2: 有展示机会,上不了首屏、1: 无展现机会,更上不了首屏)
    WirelessMatchflag   int64 `json:"wireless_matchflag,omitempty"`

    // 推广组id
    AdgroupId   int64 `json:"adgroup_id,omitempty"`

    // 词是否能推左
    Plflag   int64 `json:"plflag,omitempty"`

    // 移动相关性，同kwscore
    WirelessRelescore   string `json:"wireless_relescore,omitempty"`

    // 创意质量(创意效果)
    Creativescore   string `json:"creativescore,omitempty"`

    // 移动质量得分（另外值为-1:所属的campaign没有该投放设备，且所属的adgroup有该投放设备的投放中的创意、-2:所属的campaign有该投放设备，且所属的adgroup没有该投放设备的投放中的创意、-3:所属的campaign没有该投放设备，且所属的adgroup没有该投放设备的投放中的创意）
    WirelessQscore   string `json:"wireless_qscore,omitempty"`

    // PC质量得分（另外值为-1:所属的campaign没有该投放设备，且所属的adgroup有该投放设备的投放中的创意、-2:所属的campaign有该投放设备，且所属的adgroup没有该投放设备的投放中的创意、-3:所属的campaign没有该投放设备，且所属的adgroup没有该投放设备的投放中的创意）
    Qscore   string `json:"qscore,omitempty"`

    // 类目质量得分
    Catscore   string `json:"catscore,omitempty"`

    // 客户id
    CustomerId   int64 `json:"customer_id,omitempty"`

    // 店铺质量(账户表现)
    Custscore   string `json:"custscore,omitempty"`

    // 昵称
    Nick   string `json:"nick,omitempty"`

    // 词相关性分数
    Kwscore   string `json:"kwscore,omitempty"`

    // 点击转化率(买家体验)
    Cvrscore   string `json:"cvrscore,omitempty"`

    // 属性得分
    Pscore   string `json:"pscore,omitempty"`

    // 关键词
    Word   string `json:"word,omitempty"`

    // 移动点击转化率(买家体验)
    WirelessCvrscore   string `json:"wireless_cvrscore,omitempty"`

    // 广告类型  单品： &quot;tbuad&quot;; 店铺： &quot;addp&quot;;
    AdType   string `json:"ad_type,omitempty"`

    // 词id
    KeywordId   int64 `json:"keyword_id,omitempty"`

    // 最低展现出价
    MinPrice   int64 `json:"min_price,omitempty"`

    // 词在pc端是否能首页推左(0:不能推左、1:可以推左)，此标记仅代表首页推左标
    PcLeftFlag   int64 `json:"pc_left_flag,omitempty"`

}
