package vaccin

// AlipayVaccineUserBindDto 
type AlipayVaccineUserBindDto struct {
    // 预约日期：2019-02-08 严格按照
    ReserveDate   string `json:"reserve_date,omitempty" xml:"reserve_date,omitempty"`
    // 年龄类型(1-宝宝2-成人)
    AgeType   string `json:"age_type,omitempty" xml:"age_type,omitempty"`
    // 接种点编码
    PovNo   string `json:"pov_no,omitempty" xml:"pov_no,omitempty"`
    // 绑定人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // ISV 侧接种人 ID
    IsvUserId   string `json:"isv_user_id,omitempty" xml:"isv_user_id,omitempty"`
    // 接种点名称
    PovName   string `json:"pov_name,omitempty" xml:"pov_name,omitempty"`
    // 接种人出生日期
    Birthday   string `json:"birthday,omitempty" xml:"birthday,omitempty"`
}
