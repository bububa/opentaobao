package tmallservice

// ServiceProvider 
type ServiceProvider struct {
    // 网点编码
    ServiceStoreCode   string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
    // 网点名称
    ServiceStoreName   string `json:"service_store_name,omitempty" xml:"service_store_name,omitempty"`
    // 工人电话
    WorkerMobile   string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
    // 网点id
    ServiceStoreId   int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
    // 工人姓名
    WorkerName   string `json:"worker_name,omitempty" xml:"worker_name,omitempty"`
    // 服务商昵称
    TpNick   string `json:"tp_nick,omitempty" xml:"tp_nick,omitempty"`
    // 服务商id
    TpId   int64 `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
    // 工人id
    WorkerId   int64 `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
    // isv服务商
    IsvTpId   int64 `json:"isv_tp_id,omitempty" xml:"isv_tp_id,omitempty"`
    // isv服务商
    IsvTpName   string `json:"isv_tp_name,omitempty" xml:"isv_tp_name,omitempty"`
}
