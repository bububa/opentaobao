package alilabs

// Payload 
type Payload struct {
    // 用户token，云云接入设备必填
    Token   string `json:"token,omitempty" xml:"token,omitempty"`
    // 设备信息列表
    DeviceInfoList   []DeviceInfo `json:"device_info_list,omitempty" xml:"device_info_list>device_info,omitempty"`
    // 云云对接的技能id
    SkillId   string `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
    // 设备接入类型，1: 表示零配  2：表示 云云
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
