package idle

// OneSpuSpRegisterUpdateParam 结构体
type OneSpuSpRegisterUpdateParam struct {
	// 要挂载的闲鱼业务 PD 值，如：传 &#34;YHB&#34; 表示验货宝
	PdCode string `json:"pd_code,omitempty" xml:"pd_code,omitempty"`
	// 要挂载的业务场景，如：传 &#34;YHB_PHONE&#34; 表示验货宝手机类目
	SceneType string `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	// 要挂载的 SPU id
	XpuId string `json:"xpu_id,omitempty" xml:"xpu_id,omitempty"`
	// 挂载动作。-1（下线）；1（挂载）
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 当前服务提供方的标识符，如 APPKEY
	SpCode string `json:"sp_code,omitempty" xml:"sp_code,omitempty"`
}
